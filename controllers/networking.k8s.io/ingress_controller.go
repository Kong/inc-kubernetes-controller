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

package networkingk8sio

import (
	"context"
	"fmt"
	"time"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/kong/inc-kubernetes-controller/internal/kapi"
	modelv1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/model/v1"
	servicev1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/service/v1"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
)

// IngressReconciler reconciles a Ingress object
type IngressReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Store  *kapi.Services
}

var (
	testRoute = modelv1.Route{
		Id:        "207b3449-f2d1-463b-9177-4cbe6dd9612d",
		Name:      "default.nanana.httpbin.kong.example.80",
		Hosts:     []string{"kong.example"},
		Paths:     []string{"/", "/example"},
		Protocols: []string{"http", "https"},
		StripPath: &wrapperspb.BoolValue{Value: true},
		Tags:      []string{"k8s-uid:3191a4ce-0102-4eef-b65c-567886d95971"},
		Service: &modelv1.Service{
			Id: "0e2dc99a-d395-4828-b3b3-31393b9f1583",
		},
	}

	testService = modelv1.Service{
		Id:       "0e2dc99a-d395-4828-b3b3-31393b9f1583",
		Name:     "default.nanana.httpbin.80",
		Host:     "httpbin.org",
		Port:     80,
		Protocol: "http",
		Tags:     []string{"k8s-uid:3ac16689-f41c-4d08-b429-75250741fa71"},
	}
)

//+kubebuilder:rbac:groups=networking.k8s.io.konghq.com,resources=ingresses,verbs=get;list;watch
//+kubebuilder:rbac:groups=networking.k8s.io.konghq.com,resources=ingresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networking.k8s.io.konghq.com,resources=ingresses/finalizers,verbs=update

func (r *IngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// TODO confirm how this works, and that it just uses the zap log from the manager
	// Existing KIC controllers have their logger defined as part of the Reconciler struct
	logger := log.FromContext(ctx)
	var err error

	// TODO garbage to provide koko API with a logger
	// TODO koko wants a logger via the context, somehow, and this _does not_ actually work to set it. it needs to
	opts := zap.Options{
		Development: true,
	}
	ktx := context.WithValue(ctx, util.LoggerKey, zap.NewRaw(zap.UseFlagOptions(&opts)))

	var ingress netv1.Ingress
	if err := r.Get(ctx, req.NamespacedName, &ingress); err != nil {
		if errors.IsNotFound(err) {
			// TODO Delete from store

			// These should only happen without finalizers, but we probably still run without finalizers for DB-less:
			// Either we're running and see the request, or we're not running and the resource simply won't make it
			// into the store when the next start bootstrap populates it
		}
	}

	if !ingress.DeletionTimestamp.IsZero() && time.Now().After(ingress.DeletionTimestamp.Time) {
		// TODO delete from store
		// Because we can no longer simply delete the exact object (the <Namespace>/<Name> Ingress) from the store,
		// we need to generate route names or memoize. the former requires iterating through all the rules. the latter
		// requires either annotations containing a CSV of IDs or maintaining in-memory state. in-memory state may
		// be fine, since sqlite gets repopulated from scratch by the new leader, but not sure
	}

	// TODO class filtering goes either in here or in a filter in SetupWithManager. skipping it for now

	// TODO load and update services in backends. these need to be created before the route
	translator := newIngressServiceTranslator()
	translator.populateProtoRoutesFromIngress(&ingress)
	translator.Services, err = r.getProtoRouteServices(ctx, translator.ProtoRoutes)
	if err != nil {
		return ctrl.Result{}, err
	}
	for _, service := range translator.Services {
		translator.PortNames[service.Name], translator.PortNumbers[service.Name] = servicePortIndices(service)
	}
	err = translator.fillKongServices()
	if err != nil {
		return ctrl.Result{}, err
	}
	err = translator.fillKongRoutes()
	if err != nil {
		return ctrl.Result{}, err
	}
	for _, i := range translator.KongServices {
		// TODO do we have any use for the response this currently discards?
		_, err = r.Store.Service.UpsertService(ktx, &servicev1.UpsertServiceRequest{Item: i})
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	for _, i := range translator.KongRoutes {
		_, err = r.Store.Route.UpsertRoute(ktx, &servicev1.UpsertRouteRequest{Item: i})
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	for _, i := range translator.KongServices {
		gotService, err := r.Store.Service.GetService(ctx, &servicev1.GetServiceRequest{Id: i.Id})
		if err != nil {
			logger.Error(err, "could not get service")
		}
		logger.Info(fmt.Sprintf("added service %s to the store", gotService.Item.Id))
	}

	for _, i := range translator.KongRoutes {
		gotRoute, err := r.Store.Route.GetRoute(ctx, &servicev1.GetRouteRequest{Id: i.Id})
		if err != nil {
			logger.Error(err, "could not get route")
		}
		logger.Info(fmt.Sprintf("added route %s to the store", gotRoute.Item.Id))
	}

	// TODO load and update certificates/SNIs. these can be handled whenever
	// TODO load and update plugins. these must be created after the route

	// TODO reporting? probably not for POC

	// TODO requeues currently on success, shouldn't. not sure why, we're not even updating status
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&netv1.Ingress{}).
		Complete(r)
}
