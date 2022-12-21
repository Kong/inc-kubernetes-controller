package control

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/client-go/util/retry"

	servicev1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/service/v1"
	relay "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/relay/service/v1"
	grpcKongUtil "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/util/v1"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server/kong/ws"
	kongConfigWS "github.com/kong/inc-kubernetes-controller/internal/koko/server/kong/ws/config"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server/kong/ws/config/compat"
	serverUtil "github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
)

var (
	demoCert = []byte(`-----BEGIN CERTIFICATE-----
MIICejCCAiCgAwIBAgIUM/uXwOFH6RmU8DZMj85TerHjrQgwCgYIKoZIzj0EAwIw
gYQxCzAJBgNVBAYTAlVTMRQwEgYDVQQHEwtNaW5uZWFwb2xpczEUMBIGA1UEChML
WWFrIFNoYXZpbmcxDDAKBgNVBAsTA1dvbzE7MDkGA1UEAxMyWWFrIFNoYXZlcyBG
YWtlIEludGVybWVkaWF0ZSBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkwHhcNMTgwMzE1
MjIwNDAwWhcNNDgwMzA3MjIwNDAwWjCBjDELMAkGA1UEBhMCVVMxFDASBgNVBAcT
C01pbm5lYXBvbGlzMRQwEgYDVQQKEwtZYWsgU2hhdmluZzEMMAoGA1UECxMDV29v
MUMwQQYDVQQDEzpZYWsgU2hhdmVzIEZha2UgSW50ZXJtZWRpYXRlIENlcnRpZmlj
YXRlIEF1dGhvcml0eSBMZXZlbCAyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
tBwGcExavS1hi32ovoU/VUtbQw9Ah15ypMOsUCKkvbz7waPCpRMtkiO7TqFH9gQD
ZYUz+RKvzh2XGTvvwiEBGqNmMGQwDgYDVR0PAQH/BAQDAgGGMBIGA1UdEwEB/wQI
MAYBAf8CAQQwHQYDVR0OBBYEFIo/bUwbPmcFeJKLx/yYcgFaFP/VMB8GA1UdIwQY
MBaAFNjbCjYBxylShp5O8FCPhx4MxrqUMAoGCCqGSM49BAMCA0gAMEUCIQCwsXPy
k+JSGQNPVlZyZyI74WDb0y8fnuGaZl0kTQNw2wIgANjimow/L1M5uwl22uAS0tcA
VRXF4m/aMVPkrJTGP6s=
-----END CERTIFICATE-----`)
	demoKey = []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIMLjAzy9GIRvZpJ29CHaW04bWfVtpqrMFTwYPtDpmQBQoAoGCCqGSM49
AwEHoUQDQgAEtBwGcExavS1hi32ovoU/VUtbQw9Ah15ypMOsUCKkvbz7waPCpRMt
kiO7TqFH9gQDZYUz+RKvzh2XGTvvwiEBGg==
-----END EC PRIVATE KEY-----`)
)

const (
	readTimeout = 5 * time.Second
)

// TODO setting up the wrpc handler requires the relay (3001) and event (no port?) services, which aren't being set up
// yet

type WRPCServer struct {
	server *http.Server
}

func (w *WRPCServer) Start(ctx context.Context) error {
	errCh := make(chan error)
	// TODO chicken-egg problem here where setup needs to execute after the controller-runtime manager starts runnables
	// to connect to the relay server, but Run doesn't accept the logger. setup stages need to be separated somehow.
	// for now, cheat and just create an unrelated logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	if w.server, err = SetUpWRPCHandler(ctx, logger); err != nil {
		return err
	}
	go func() {
		listener, err := net.Listen("tcp", w.server.Addr)
		if err != nil {
			errCh <- err
			return
		}
		if w.server.TLSConfig != nil {
			listener = tls.NewListener(listener, w.server.TLSConfig)
		}
		err = w.server.Serve(listener)
		if err != nil {
			if err != http.ErrServerClosed {
				errCh <- err
			}
		}
	}()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		ctx, cleanup := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanup()
		// ctx not inherited since the parent ctx will already be Done()
		// at this point
		return w.server.Shutdown(ctx)
	}
}

func SetUpWRPCHandler(ctx context.Context, logger *zap.Logger) (*http.Server, error) {
	vcLogger := logger.With(zap.String("component", "version-compatibility"))
	vc, err := kongConfigWS.NewVersionCompatibilityProcessor(kongConfigWS.VersionCompatibilityOpts{
		Logger:         vcLogger,
		KongCPVersion:  kongConfigWS.KongGatewayCompatibilityVersion,
		ExtraProcessor: compat.VersionCompatibilityExtraProcessing,
	})

	// setup relay client
	const grpcMaxSendMsgSize = 1024 * 1024 * 8
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize),
			grpc.MaxCallRecvMsgSize(grpcMaxSendMsgSize),
		),
	}
	// TODO the relay server may not be running at start, since we can't guarantee the order of runnables. this retry
	// is intended to deal with that. not sure if there's a better way (maybe we can define runnable dependencies or
	// retry behavior? not seeing any way to)
	var cc *grpc.ClientConn
	retry.OnError(
		// TODO this apparently never hits a limit, there's no Cap defined in the default
		retry.DefaultBackoff,
		func(_ error) bool {
			return true
		},
		func() error {
			cc, err = grpc.Dial("localhost:3001", dialOpts...)
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	grpcClientList := setupGRPCClients(cc)

	loader := &kongConfigWS.KongConfigurationLoader{}
	if err = registerClients(loader, grpcClientList); err != nil {
		return nil, err
	}

	m, err := ws.NewManager(ws.ManagerOpts{
		Ctx:                    ctx,
		Logger:                 logger,
		DPConfigLoader:         loader,
		DPVersionCompatibility: vc,
		Client: ws.ConfigClient{
			Node:   grpcClientList.Node,
			Status: grpcClientList.Status,
			Event:  grpcClientList.Event,
		},
		Cluster: ws.DefaultCluster{},
		// TODO(hbagdi): make this configurable
		Config: ws.ManagerConfig{
			DataPlaneRequisites: []*grpcKongUtil.DataPlanePrerequisite{
				{
					Config: &grpcKongUtil.DataPlanePrerequisite_RequiredPlugins{
						RequiredPlugins: &grpcKongUtil.RequiredPluginsFilter{
							RequiredPlugins: []string{"rate-limiting"},
						},
					},
				},
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create manager: %w", err)
	}
	// TODO this is normally determined by cmd/run and would select between shared/PKI modes, and allow cert config
	cert, err := tls.X509KeyPair(demoCert, demoKey)
	if err != nil {
		return nil, err
	}
	authFn, err := ws.AuthFnSharedTLS(cert)
	if err != nil {
		return nil, err
	}
	authenticator := &ws.DefaultAuthenticator{
		Manager: m,
		Context: ctx,
		AuthFn:  authFn,
	}

	controlLogger := logger.With(zap.String("component", "control-server"))

	negotiator, err := ws.NewNegotiationRegisterer(controlLogger.With(
		zap.String("protocol", "wRPC"),
		zap.String("wrpc-service", "negotiation")))
	if err != nil {
		return nil, err
	}

	err = negotiator.AddService("config", "v1", "wRPC configuration", &ws.ConfigRegisterer{})
	if err != nil {
		return nil, err
	}

	handler, err := ws.NewHandler(ws.HandlerOpts{
		Logger:        controlLogger,
		Authenticator: authenticator,
		BaseServices:  negotiator,
	})
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    ":3100",
		Handler: serverUtil.HandlerWithRecovery(serverUtil.HandlerWithLogger(handler, controlLogger), controlLogger),
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert}, // TODO use non fake cert
			ClientAuth:   tls.RequestClientCert,
		},
		ReadHeaderTimeout: readTimeout,
		ReadTimeout:       readTimeout,
	}
	return server, nil
}

func registerClients(loader *kongConfigWS.KongConfigurationLoader, grpcClientList grpcClients) error {
	err := loader.Register(&kongConfigWS.KongServiceLoader{Client: grpcClientList.Service})
	if err != nil {
		return fmt.Errorf("failed to register service configuration loader: %w", err)
	}
	err = loader.Register(&kongConfigWS.KongRouteLoader{Client: grpcClientList.Route})
	if err != nil {
		return fmt.Errorf("failed to register route configuration loader: %w", err)
	}
	err = loader.Register(&kongConfigWS.KongPluginLoader{Client: grpcClientList.Plugin})
	if err != nil {
		return fmt.Errorf("failed to register plugin configuration loader: %w", err)
	}

	err = loader.Register(&kongConfigWS.KongUpstreamLoader{Client: grpcClientList.Upstream})
	if err != nil {
		return fmt.Errorf("failed to register upstream configuration loader"+
			": %w", err)
	}

	err = loader.Register(&kongConfigWS.KongTargetLoader{Client: grpcClientList.Target})
	if err != nil {
		return fmt.Errorf("failed to register target configuration loader: %w"+
			"", err)
	}

	err = loader.Register(&kongConfigWS.KongConsumerLoader{Client: grpcClientList.Consumer})
	if err != nil {
		return fmt.Errorf("failed to register consumer configuration loader"+
			": %w", err)
	}

	err = loader.Register(&kongConfigWS.KongCertificateLoader{Client: grpcClientList.Certificate})
	if err != nil {
		return fmt.Errorf("failed to register certificate configuration"+
			" loader: %w", err)
	}

	err = loader.Register(&kongConfigWS.KongCACertificateLoader{Client: grpcClientList.CACertificate})
	if err != nil {
		return fmt.Errorf("failed to register ca-certificate configuration"+
			" loader: %w", err)
	}

	err = loader.Register(&kongConfigWS.KongSNILoader{Client: grpcClientList.SNI})
	if err != nil {
		return fmt.Errorf("failed to register sni configuration loader: %w",
			err)
	}

	err = loader.Register(&kongConfigWS.KongVaultLoader{Client: grpcClientList.Vault})
	if err != nil {
		return fmt.Errorf("failed to register vault configuration loader: %w",
			err)
	}

	err = loader.Register(&kongConfigWS.VersionLoader{})
	if err != nil {
		return fmt.Errorf("failed to register version configuration loader"+
			": %w", err)
	}

	// TODO this argument was originally generated from cmd.registerInstallation. unclear if it'd matter for an
	// embedded instance
	parametersLoader, err := kongConfigWS.NewParametersLoader("local")
	if err != nil {
		return fmt.Errorf("failed to create parameters configuration loader"+
			": %w", err)
	}
	err = loader.Register(parametersLoader)
	if err != nil {
		return fmt.Errorf("failed to register parameters configuration loader"+
			": %w", err)
	}
	return nil
}

// yanked from cmd/run as-is minus the v1->servicev1 rename, didn't want to import cmd here

type grpcClients struct {
	Service       servicev1.ServiceServiceClient
	Route         servicev1.RouteServiceClient
	Plugin        servicev1.PluginServiceClient
	PluginSchema  servicev1.PluginSchemaServiceClient
	Upstream      servicev1.UpstreamServiceClient
	Target        servicev1.TargetServiceClient
	Consumer      servicev1.ConsumerServiceClient
	Certificate   servicev1.CertificateServiceClient
	CACertificate servicev1.CACertificateServiceClient
	SNI           servicev1.SNIServiceClient
	Vault         servicev1.VaultServiceClient

	Status relay.StatusServiceClient
	Node   servicev1.NodeServiceClient
	Event  relay.EventServiceClient
}

func setupGRPCClients(cc *grpc.ClientConn) grpcClients {
	return grpcClients{
		Service:       servicev1.NewServiceServiceClient(cc),
		Route:         servicev1.NewRouteServiceClient(cc),
		Plugin:        servicev1.NewPluginServiceClient(cc),
		PluginSchema:  servicev1.NewPluginSchemaServiceClient(cc),
		Upstream:      servicev1.NewUpstreamServiceClient(cc),
		Target:        servicev1.NewTargetServiceClient(cc),
		Consumer:      servicev1.NewConsumerServiceClient(cc),
		Certificate:   servicev1.NewCertificateServiceClient(cc),
		CACertificate: servicev1.NewCACertificateServiceClient(cc),
		SNI:           servicev1.NewSNIServiceClient(cc),
		Vault:         servicev1.NewVaultServiceClient(cc),

		Node:   servicev1.NewNodeServiceClient(cc),
		Event:  relay.NewEventServiceClient(cc),
		Status: relay.NewStatusServiceClient(cc),
	}
}
