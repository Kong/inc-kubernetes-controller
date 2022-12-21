/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"

	corev1 "k8s.io/api/core/v1"
	// networkingv1 "k8s.io/api/networking/v1"
	"github.com/kong/inc-kubernetes-controller/internal/control"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server/admin"
	serverUtil "github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
	rzap "go.uber.org/zap"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	configurationkonghqcomv1 "github.com/kong/inc-kubernetes-controller/apis/configuration.konghq.com/v1"
	networkingk8siocontrollers "github.com/kong/inc-kubernetes-controller/controllers/networking.k8s.io"
	"github.com/kong/inc-kubernetes-controller/internal/datastore"
	relay "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/relay/service/v1"
	relayImpl "github.com/kong/inc-kubernetes-controller/internal/koko/server/relay"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(configurationkonghqcomv1.AddToScheme(scheme))
	// TODO unclear if necessary for core types. kic does not do this for networking.
	// utilruntime.Must(networkingv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	// TODO get appropriate ctx for runnables
	ctx := context.Background()
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	// TODO controller-runtime wants a logr.Logger using zap, koko wants a zap.Logger. deja vu, 5000 logging libs
	zlogger := zap.NewRaw(zap.UseFlagOptions(&opts))
	logger := zap.New(zap.UseFlagOptions(&opts))
	ctrl.SetLogger(logger)

	// TODO temp local config
	//mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
	kcfg, err := getKubeconfig()
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}
	mgr, err := ctrl.NewManager(kcfg, ctrl.Options{
		Scheme:                 scheme,
		Namespace:              corev1.NamespaceAll,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "613b5b82.konghq.com",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	var store datastore.StoreRunner

	if err = mgr.Add(&store); err != nil {
		setupLog.Error(err, "could not add data store")
		os.Exit(1)
	}

	// relay
	storeLoader := serverUtil.DefaultStoreLoader{Store: store.Store}
	adminOpts := admin.HandlerOpts{
		Logger:      zlogger.With(rzap.String("component", "admin-server")),
		StoreLoader: storeLoader,
	}
	rawGRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			serverUtil.LoggerInterceptor(zlogger),
			serverUtil.PanicInterceptor(zlogger)),
		grpc.ChainStreamInterceptor(serverUtil.PanicStreamInterceptor(adminOpts.Logger)))
	admin.RegisterAdminService(rawGRPCServer, adminOpts)

	grpcServer, err := server.NewGRPC(server.GRPCOpts{
		Address:    ":3001",
		GRPCServer: rawGRPCServer,
		Logger:     zlogger.With(rzap.String("component", "relay-server")),
	})
	if err != nil {
		setupLog.Error(err, "unable to create relay")
		os.Exit(1)
	}

	eventService := relayImpl.NewEventService(ctx,
		relayImpl.EventServiceOpts{
			Store:  store.Store,
			Logger: zlogger.With(rzap.String("component", "relay-server")),
		})
	relay.RegisterEventServiceServer(rawGRPCServer, eventService)
	statusService := relayImpl.NewStatusService(relayImpl.StatusServiceOpts{
		StoreLoader: storeLoader,
		Logger:      zlogger.With(rzap.String("component", "relay-server")),
	})
	relay.RegisterStatusServiceServer(rawGRPCServer, statusService)
	services := datastore.BuildServices(datastore.HandlerOpts{
		Logger:      zlogger,
		StoreLoader: &store,
		// TODO validator
	})
	if err = mgr.Add(grpcServer); err != nil {
		setupLog.Error(err, "unable to start relay")
		os.Exit(1)
	}

	wrpcServer := &control.WRPCServer{}
	if err = mgr.Add(wrpcServer); err != nil {
		setupLog.Error(err, "unable to start control server")
		os.Exit(1)
	}

	if err = (&networkingk8siocontrollers.IngressReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Store:  &services,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Ingress")
		os.Exit(1)
	}

	// TODO
	//if err = (&configurationkonghqcomcontrollers.KongPluginReconciler{
	//	Client: mgr.GetClient(),
	//	Scheme: mgr.GetScheme(),
	//	Store:  &services,
	//}).SetupWithManager(mgr); err != nil {
	//	setupLog.Error(err, "unable to create controller", "controller", "KongPlugin")
	//	os.Exit(1)
	//}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

// TODO testing only
func getKubeconfig() (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return nil, err
	}

	return config, err
}
