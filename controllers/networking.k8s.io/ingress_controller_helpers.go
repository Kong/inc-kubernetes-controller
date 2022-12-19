package networkingk8sio

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/types"

	modelv1 "github.com/kong/inc-kubernetes-controller/internal/koko/gen/grpc/kong/admin/model/v1"
)

var (
	defaultHTTPIngressPathType = netv1.PathTypeImplementationSpecific
	serviceIDNamespace         = uuid.MustParse("dd2b2f41-e654-418d-9cd3-7aa8674cfcfe")
	routeIDNamespace           = uuid.MustParse("3cff93d8-276b-42ab-8374-97d7933b28b6")
)

// Ultimately the Koko services want these types. Worth noting that Get doesn't support name, that's apparently
// handled elsewhere in the API implementation and needs to be re-implemented in our datastore
//
// Koko types invert KIC's Service > []Route relationship. A Koko Route can contain a single Koko Service, but this
// appears to be a bit of a lie, at least in the Konnect instance: this service only accepts the ID, and trying to
// provide the entity yields:
//
// "messages": [
//   "missing properties: 'id'",
//   "additionalProperties 'read_timeout', 'retries', 'url', 'write_timeout', 'enabled', 'name', 'connect_timeout' not allowed"
// ]

// internal/koko/gen/grpc/kong/admin/model/v1/service.pb.go

//type Service struct {
//	state         protoimpl.MessageState
//	sizeCache     protoimpl.SizeCache
//	unknownFields protoimpl.UnknownFields
//
//	Id                string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
//	Name              string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
//	ConnectTimeout    int32                 `protobuf:"varint,3,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
//	CreatedAt         int32                 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
//	Host              string                `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
//	Path              string                `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
//	Port              int32                 `protobuf:"varint,7,opt,name=port,proto3" json:"port,omitempty"`
//	Protocol          string                `protobuf:"bytes,8,opt,name=protocol,proto3" json:"protocol,omitempty"`
//	ReadTimeout       int32                 `protobuf:"varint,9,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
//	Retries           int32                 `protobuf:"varint,10,opt,name=retries,proto3" json:"retries,omitempty"`
//	UpdatedAt         int32                 `protobuf:"varint,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
//	Url               string                `protobuf:"bytes,12,opt,name=url,proto3" json:"url,omitempty"`
//	WriteTimeout      int32                 `protobuf:"varint,13,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
//	Tags              []string              `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
//	TlsVerify         *bool                 `protobuf:"varint,15,opt,name=tls_verify,json=tlsVerify,proto3,oneof" json:"tls_verify,omitempty"`
//	TlsVerifyDepth    int32                 `protobuf:"varint,16,opt,name=tls_verify_depth,json=tlsVerifyDepth,proto3" json:"tls_verify_depth,omitempty"`
//	ClientCertificate *Certificate          `protobuf:"bytes,17,opt,name=client_certificate,json=clientCertificate,proto3" json:"client_certificate,omitempty"`
//	CaCertificates    []string              `protobuf:"bytes,18,rep,name=ca_certificates,json=caCertificates,proto3" json:"ca_certificates,omitempty"`
//	Enabled           *wrapperspb.BoolValue `protobuf:"bytes,19,opt,name=enabled,proto3" json:"enabled,omitempty"`
//}

// internal/koko/gen/grpc/kong/admin/model/v1/route.pb.go

//type Route struct {
//	state         protoimpl.MessageState
//	sizeCache     protoimpl.SizeCache
//	unknownFields protoimpl.UnknownFields
//
//	Id                      string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
//	Name                    string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
//	Headers                 map[string]*HeaderValues `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
//	Hosts                   []string                 `protobuf:"bytes,4,rep,name=hosts,proto3" json:"hosts,omitempty"`
//	CreatedAt               int32                    `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
//	Methods                 []string                 `protobuf:"bytes,6,rep,name=methods,proto3" json:"methods,omitempty"`
//	Paths                   []string                 `protobuf:"bytes,7,rep,name=paths,proto3" json:"paths,omitempty"`
//	PathHandling            string                   `protobuf:"bytes,8,opt,name=path_handling,json=pathHandling,proto3" json:"path_handling,omitempty"`
//	PreserveHost            *wrapperspb.BoolValue    `protobuf:"bytes,9,opt,name=preserve_host,json=preserveHost,proto3" json:"preserve_host,omitempty"`
//	Protocols               []string                 `protobuf:"bytes,10,rep,name=protocols,proto3" json:"protocols,omitempty"`
//	RegexPriority           *wrapperspb.Int32Value   `protobuf:"bytes,11,opt,name=regex_priority,json=regexPriority,proto3" json:"regex_priority,omitempty"`
//	StripPath               *wrapperspb.BoolValue    `protobuf:"bytes,12,opt,name=strip_path,json=stripPath,proto3" json:"strip_path,omitempty"`
//	UpdatedAt               int32                    `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
//	Snis                    []string                 `protobuf:"bytes,14,rep,name=snis,proto3" json:"snis,omitempty"`
//	Sources                 []*CIDRPort              `protobuf:"bytes,15,rep,name=sources,proto3" json:"sources,omitempty"`
//	Destinations            []*CIDRPort              `protobuf:"bytes,16,rep,name=destinations,proto3" json:"destinations,omitempty"`
//	Tags                    []string                 `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags,omitempty"`
//	HttpsRedirectStatusCode int32                    `protobuf:"varint,18,opt,name=https_redirect_status_code,json=httpsRedirectStatusCode,proto3" json:"https_redirect_status_code,omitempty"`
//	RequestBuffering        *wrapperspb.BoolValue    `protobuf:"bytes,19,opt,name=request_buffering,json=requestBuffering,proto3" json:"request_buffering,omitempty"`
//	ResponseBuffering       *wrapperspb.BoolValue    `protobuf:"bytes,20,opt,name=response_buffering,json=responseBuffering,proto3" json:"response_buffering,omitempty"`
//	Service                 *Service                 `protobuf:"bytes,21,opt,name=service,proto3" json:"service,omitempty"`
//}

// TODO stole this port struct from kongstate and simplified it to just panic if both are set (the API server forbids
// setting both and rejects the zero value, so it should never happen). it probably needs to be in a shared lib
// eventually, for non-Ingress
type PortDef struct {
	// Name is a port name.
	Name string

	// Number is a port number.
	Number int32
}

type protoRoute struct {
	ingressAnnotations map[string]string
	ingressNamespace   string
	ingressName        string
	ingressHost        string
	serviceName        string
	servicePortNumber  int32
	servicePortName    string
	paths              []netv1.HTTPIngressPath
	addRegexPrefix     bool
}

type ingressServiceTranslator struct {
	PortNames    map[string]map[string]int32
	PortNumbers  map[string]map[int32]interface{}
	ProtoRoutes  map[string]*protoRoute
	Services     map[string]*corev1.Service
	KongServices map[string]*modelv1.Service
	KongRoutes   map[string]*modelv1.Route
}

func newIngressServiceTranslator() ingressServiceTranslator {
	var i ingressServiceTranslator
	i.PortNames = map[string]map[string]int32{}
	i.PortNumbers = map[string]map[int32]interface{}{}
	i.Services = map[string]*corev1.Service{}
	i.KongServices = map[string]*modelv1.Service{}
	i.KongRoutes = map[string]*modelv1.Route{}
	return i
}

func (i *ingressServiceTranslator) fillKongServices() error {
	for key, route := range i.ProtoRoutes {
		service := i.Services[key]

		port, err := i.resolvePort(route)
		if err != nil {
			return err
		}
		// TODO we currently (in KIC) include the ingress name, but I forget if/why this is required for collision
		// avoidance. are we unable to de-dup services across ingresses otherwise?
		name := fmt.Sprintf("%s.%s.%d", service.Namespace, service.Name, route.servicePortNumber)
		kongService := modelv1.Service{
			Name: name,
			Id:   uuid.NewSHA1(serviceIDNamespace, []byte(name)).String(),
			// there isn't really any fitting upstream reserved TLD, but .local is arguably the best of meh options
			Host: fmt.Sprintf("%s.%s.%d.local", service.Namespace, service.Name, port),
			Port: port,
			// the rest we'll need to eventually populate from annotations or similar
		}
		i.KongServices[key] = &kongService
	}
	return nil
}

func (i *ingressServiceTranslator) fillKongRoutes() error {
	for key, route := range i.ProtoRoutes {
		service := i.KongServices[key]
		var paths []string
		for _, p := range route.paths {
			// TODO proper type handling
			paths = append(paths, p.Path)
		}
		name := fmt.Sprintf("%s.%s.%s.%s.%d", route.ingressNamespace, route.ingressName, route.serviceName, route.ingressHost, service.Port)
		kongRoute := modelv1.Route{
			Name:  name,
			Id:    uuid.NewSHA1(routeIDNamespace, []byte(name)).String(),
			Hosts: []string{route.ingressHost},
			Paths: paths,
			Service: &modelv1.Service{
				Id: service.Id,
			},
		}
		i.KongRoutes[key] = &kongRoute
	}
	return nil
}

func (i *ingressServiceTranslator) resolvePort(route *protoRoute) (int32, error) {
	var port int32
	if route.servicePortNumber != 0 {
		if _, ok := i.PortNumbers[route.serviceName][route.servicePortNumber]; !ok {
			return port, fmt.Errorf("Ingress %s/%s references non-existent Service %s port %d",
				route.ingressNamespace, route.ingressName, route.serviceName, route.servicePortNumber)
		}
		port = route.servicePortNumber
	} else if route.servicePortName != "" {
		if _, ok := i.PortNames[route.serviceName][route.servicePortName]; !ok {
			return port, fmt.Errorf("Ingress %s/%s references non-existent Service %s port %s",
				route.ingressNamespace, route.ingressName, route.serviceName, route.servicePortName)
		}
		port = i.PortNames[route.serviceName][route.servicePortName]
	} else {
		// implicit port
		if _, ok := i.PortNames[route.serviceName][""]; !ok {
			return port, fmt.Errorf("Ingress %s/%s cannot use implicit port for %s",
				route.ingressNamespace, route.ingressName, route.serviceName)
		}
		port = i.PortNames[route.serviceName][""]
	}
	return port, nil
}

func (i *ingressServiceTranslator) populateProtoRoutesFromIngress(ingress *netv1.Ingress) {
	routes := map[string]*protoRoute{}

	for _, rule := range ingress.Spec.Rules {
		if rule.HTTP == nil || len(rule.HTTP.Paths) < 1 {
			continue
		}

		for _, httpIngressPath := range rule.HTTP.Paths {
			// TODO necessary? Kong or the API server already apply normalization rules for this AFAIK
			//httpIngressPath.Path = flattenMultipleSlashes(httpIngressPath.Path)

			if httpIngressPath.Path == "" {
				httpIngressPath.Path = "/"
			}

			if httpIngressPath.PathType == nil {
				httpIngressPath.PathType = &defaultHTTPIngressPathType
			}

			serviceName := httpIngressPath.Backend.Service.Name
			servicePortNumber := httpIngressPath.Backend.Service.Port.Number

			cacheKey := fmt.Sprintf("%s.%s.%s.%s.%d", ingress.Namespace, ingress.Name, rule.Host, serviceName, servicePortNumber)
			meta, ok := routes[cacheKey]
			if !ok {
				meta = &protoRoute{
					ingressNamespace:  ingress.Namespace,
					ingressName:       ingress.Name,
					ingressHost:       rule.Host,
					serviceName:       serviceName,
					servicePortNumber: servicePortNumber,
				}
			}

			meta.paths = append(meta.paths, httpIngressPath)
			meta.ingressAnnotations = ingress.Annotations
			routes[cacheKey] = meta
		}
	}
	i.ProtoRoutes = routes
}

func (r *IngressReconciler) getProtoRouteServices(
	ctx context.Context,
	routes map[string]*protoRoute,
) (map[string]*corev1.Service, error) {
	services := make(map[string]*corev1.Service, len(routes))
	for key, route := range routes {
		var service corev1.Service
		nn := types.NamespacedName{
			Name:      route.serviceName,
			Namespace: route.ingressNamespace,
		}
		if err := r.Get(ctx, nn, &service); err != nil {
			return services, fmt.Errorf("could not retrieve Service %s/%s for Ingress %s/%s: %w",
				route.ingressNamespace, route.serviceName, route.ingressNamespace, route.ingressName, err)
		}
		services[key] = &service
	}
	return services, nil
}

func servicePortIndices(service *corev1.Service) (map[string]int32, map[int32]interface{}) {
	portNames := make(map[string]int32, len(service.Spec.Ports))
	ports := make(map[int32]interface{}, len(service.Spec.Ports))

	if len(service.Spec.Ports) == 1 {
		// implicit available
		portNames[""] = service.Spec.Ports[0].Port
	}

	for _, port := range service.Spec.Ports {
		if port.Name != "" {
			portNames[port.Name] = port.Port
		}
		ports[port.Port] = nil
	}
	return portNames, ports
}
