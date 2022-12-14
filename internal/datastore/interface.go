package datastore

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kong/inc-kubernetes-controller/internal/koko/json"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	v1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/service/v1"
	svcs "github.com/kong/inc-kubernetes-controller/internal/koko/server/admin"
)

// KOKO ripped from internal/server/admin/handler.go

type services struct {
	service       v1.ServiceServiceServer
	route         v1.RouteServiceServer
	plugin        v1.PluginServiceServer
	pluginSchema  v1.PluginSchemaServiceServer
	upstream      v1.UpstreamServiceServer
	target        v1.TargetServiceServer
	schemas       v1.SchemasServiceServer
	certificate   v1.CertificateServiceServer
	consumer      v1.ConsumerServiceServer
	caCertificate v1.CACertificateServiceServer
	sni           v1.SNIServiceServer
	vault         v1.VaultServiceServer

	status v1.StatusServiceServer
	node   v1.NodeServiceServer
}

func buildServices(store StoreRunner) services {
	return services{
		service: &svcs.ServiceService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "service"),
				},
			},
		},
		route: &svcs.RouteService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "route"),
				},
			},
		},
		plugin: &svcs.PluginService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "plugin"),
				},
			},
			validator: opts.Validator,
		},
		pluginSchema: &svcs.PluginSchemaService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "plugin-schema"),
				},
			},
			validator: opts.Validator,
		},
		upstream: &svcs.UpstreamService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "upstream"),
				},
			},
		},
		target: &svcs.TargetService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "target"),
				},
			},
		},
		schemas: &svcs.SchemasService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "schemas"),
				},
			},
			validator: opts.Validator,
		},
		node: &svcs.NodeService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "node"),
				},
			},
		},
		status: &svcs.StatusService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "status"),
				},
			},
		},
		certificate: &svcs.CertificateService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "certificate"),
				},
			},
		},
		caCertificate: &svcs.CACertificateService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "ca-certificate"),
				},
			},
		},
		consumer: &svcs.ConsumerService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "consumer"),
				},
			},
		},
		sni: &svcs.SNIService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "sni"),
				},
			},
		},
		vault: &svcs.VaultService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: &store,
				LoggerFields: []zapcore.Field{
					zap.String("admin-service", "vault"),
				},
			},
		},
	}
}

func NewHandler(opts HandlerOpts) (http.Handler, error) {
	err := validateOpts(opts)
	if err != nil {
		return nil, err
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, json.Marshaller),
		runtime.WithErrorHandler(util.ErrorHandler),
		runtime.WithForwardResponseOption(util.SetHTTPStatus),
		runtime.WithForwardResponseOption(util.FinishTrace),
	)

	err = v1.RegisterMetaServiceHandlerServer(context.Background(),
		mux, &MetaService{})
	if err != nil {
		return nil, err
	}

	services := buildServices(opts)
	err = v1.RegisterServiceServiceHandlerServer(context.Background(),
		mux, services.service)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterRouteServiceHandlerServer(context.Background(),
		mux, services.route)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterPluginServiceHandlerServer(context.Background(),
		mux, services.plugin)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterPluginSchemaServiceHandlerServer(context.Background(),
		mux, services.pluginSchema)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterUpstreamServiceHandlerServer(context.Background(),
		mux, services.upstream)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterTargetServiceHandlerServer(context.Background(),
		mux, services.target)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterSchemasServiceHandlerServer(context.Background(),
		mux, services.schemas)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterNodeServiceHandlerServer(context.Background(),
		mux, services.node)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterStatusServiceHandlerServer(context.Background(),
		mux, services.status)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterConsumerServiceHandlerServer(context.Background(),
		mux, services.consumer)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterCertificateServiceHandlerServer(context.Background(),
		mux, services.certificate)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterCACertificateServiceHandlerServer(context.Background(),
		mux, services.caCertificate)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterSNIServiceHandlerServer(context.Background(),
		mux, services.sni)
	if err != nil {
		return nil, err
	}

	err = v1.RegisterVaultServiceHandlerServer(context.Background(),
		mux, services.vault)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
