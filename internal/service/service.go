package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"
)

const (
	endpointUser      = "/svc/v1/user"
	endpointCalculate = "/svc/v1/user/calculate"
	endpointAccount   = "/svc/v1/account"
	endpointAccounts  = "/svc/v1/accounts"
)

func New() kit.Service {
	return &service{}
}

type service struct{}

func (s service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	return map[string]map[string]kit.HTTPEndpoint{
		endpointUser: {
			http.MethodPost: {
				Decoder:  decodePostUser,
				Endpoint: s.postUser,
				Encoder:  kit.EncodeProtoResponse,
			},
			http.MethodGet: {
				Decoder:  decodeGetUser,
				Endpoint: s.getUser,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
		endpointCalculate: {
			http.MethodGet: {
				Decoder:  decodeUserCalculate,
				Endpoint: s.getUserCalculate,
			},
		},
		endpointAccount: {
			http.MethodPost: {
				Decoder:  decodePostAccount,
				Endpoint: s.postAccount,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
	}
}

// HTTPMiddleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s service) HTTPMiddleware(h http.Handler) http.Handler {
	return h
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
