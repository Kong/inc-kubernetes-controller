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

	"github.com/kong/inc-kubernetes-controller/internal/datastore"
	pbmodel "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/model/v1"
	kokov1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/service/v1"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
)

// IngressReconciler reconciles a Ingress object
type IngressReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Store  *datastore.Services
}

var (
	testRoute = pbmodel.Route{
		Id:        "207b3449-f2d1-463b-9177-4cbe6dd9612d",
		Name:      "default.nanana.httpbin.kong.example.80",
		Hosts:     []string{"kong.example"},
		Paths:     []string{"/", "/example"},
		Protocols: []string{"http", "https"},
		StripPath: &wrapperspb.BoolValue{Value: true},
		Tags:      []string{"k8s-uid:3191a4ce-0102-4eef-b65c-567886d95971"},
		Service: &pbmodel.Service{
			Id: "0e2dc99a-d395-4828-b3b3-31393b9f1583",
		},
	}

	testService = pbmodel.Service{
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

	// TODO garbage to provide koko API with a logger
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
	}

	// TODO class filtering goes either in here or in a filter in SetupWithManager. skipping it for now

	// TODO load and update services in backends. these need to be created before the route
	// TODO fake service for now. do we need anything other than the Item? who knows!
	r.Store.Service.UpsertService(ktx, &kokov1.UpsertServiceRequest{Item: &testService})

	// TODO create/update route properly
	r.Store.Route.UpsertRoute(ktx, &kokov1.UpsertRouteRequest{Item: &testRoute})

	gotService, err := r.Store.Service.GetService(ctx, &kokov1.GetServiceRequest{Id: "0e2dc99a-d395-4828-b3b3-31393b9f1583"})
	if err != nil {
		logger.Error(err, "could not get service")
	}
	gotRoute, err := r.Store.Route.GetRoute(ctx, &kokov1.GetRouteRequest{Id: "207b3449-f2d1-463b-9177-4cbe6dd9612d"})
	if err != nil {
		logger.Error(err, "could not get route")
	}
	logger.Info(fmt.Sprintf("added service %s and route %s to the store", gotService.Item.Id, gotRoute.Item.Id))

	// TODO load and update certificates/SNIs. these can be handled whenever
	// TODO load and update plugins. these must be created after the route

	// TODO reporting? probably not for POC

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&netv1.Ingress{}).
		Complete(r)
}
