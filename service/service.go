package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/NYTimes/gziphandler"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"
)

func New() kit.Service {
	return &service{}
}

type service struct{}

func (s service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	return map[string]map[string]kit.HTTPEndpoint{
		"/svc/v1/user/calculate": {
			"GET": {
				Decoder:  decodeUserCalculate,
				Endpoint: s.getUserCalculate,
			},
		},
	}
}

// HTTPMiddleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s service) HTTPMiddleware(h http.Handler) http.Handler {
	return gziphandler.GzipHandler(server.CORSHandler(h, ""))
}

func (s service) HTTPRouterOptions() []kit.RouterOption {
	return nil
}

func (s service) HTTPOptions() []httptransport.ServerOption {
	return nil
}

// Middleware provides an kit/endpoint.Middleware hook wrapped around all requests.
func (s service) Middleware(e endpoint.Endpoint) endpoint.Endpoint {
	return e
}

func (s service) RPCMiddleware() grpc.UnaryServerInterceptor {
	return nil
}

func (s service) RPCOptions() []grpc.ServerOption {
	return nil
}

func (s service) RPCServiceDesc() *grpc.ServiceDesc {
	return nil
}
