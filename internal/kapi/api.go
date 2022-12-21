package kapi

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	v1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/service/v1"
	"github.com/kong/inc-kubernetes-controller/internal/koko/plugin"
	svcs "github.com/kong/inc-kubernetes-controller/internal/koko/server/admin"
	"github.com/kong/inc-kubernetes-controller/internal/koko/server/util"
)

// KOKO ripped from internal/server/admin/handler.go

// ContextKey type must be used to manipulate the context of a request.
type ContextKey struct{}

type HandlerOpts struct {
	Logger *zap.Logger

	StoreLoader util.StoreLoader

	Validator plugin.Validator
}

type CommonOpts struct {
	StoreLoader  util.StoreLoader
	LoggerFields []zapcore.Field
}

type Services struct {
	Service       v1.ServiceServiceServer
	Route         v1.RouteServiceServer
	Plugin        v1.PluginServiceServer
	PluginSchema  v1.PluginSchemaServiceServer
	Upstream      v1.UpstreamServiceServer
	Target        v1.TargetServiceServer
	Schemas       v1.SchemasServiceServer
	Certificate   v1.CertificateServiceServer
	Consumer      v1.ConsumerServiceServer
	CACertificate v1.CACertificateServiceServer
	SNI           v1.SNIServiceServer
	Vault         v1.VaultServiceServer

	status v1.StatusServiceServer
	node   v1.NodeServiceServer
}

func BuildServices(opts HandlerOpts) Services {
	return Services{
		Service: &svcs.ServiceService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Service"),
				},
			},
		},
		Route: &svcs.RouteService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Route"),
				},
			},
		},
		Plugin: &svcs.PluginService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Plugin"),
				},
			},
			Validator: opts.Validator,
		},
		PluginSchema: &svcs.PluginSchemaService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Plugin-schema"),
				},
			},
			Validator: opts.Validator,
		},
		Upstream: &svcs.UpstreamService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Upstream"),
				},
			},
		},
		Target: &svcs.TargetService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Target"),
				},
			},
		},
		Schemas: &svcs.SchemasService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Schemas"),
				},
			},
			Validator: opts.Validator,
		},
		node: &svcs.NodeService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "node"),
				},
			},
		},
		status: &svcs.StatusService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "status"),
				},
			},
		},
		Certificate: &svcs.CertificateService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Certificate"),
				},
			},
		},
		CACertificate: &svcs.CACertificateService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "ca-Certificate"),
				},
			},
		},
		Consumer: &svcs.ConsumerService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Consumer"),
				},
			},
		},
		SNI: &svcs.SNIService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "SNI"),
				},
			},
		},
		Vault: &svcs.VaultService{
			CommonOpts: svcs.CommonOpts{
				StoreLoader: opts.StoreLoader,
				LoggerFields: []zapcore.Field{
					zap.String("admin-Service", "Vault"),
				},
			},
		},
	}
}
