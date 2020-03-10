package service

import (
	context "context"
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/justindfuller/financial"
	"github.com/justindfuller/financial/internal/db"
	"github.com/justindfuller/financial/internal/db/memory"
	"google.golang.org/grpc"
)

const (
	endpointHealth       = "/svc/v1/health"
	endpointUser         = "/svc/v1/user"
	endpointCalculate    = "/svc/v1/calculate"
	endpointAccount      = "/svc/v1/account"
	endpointAccounts     = "/svc/v1/accounts"
	endpointContribution = "/svc/v1/contribution"
	endpointGoal         = "/svc/v1/goal"
	endpointGoals        = "/svc/v1/goals"
)

func New() kit.Service {
	return &service{
		db: memory.New(),
	}
}

type service struct {
	db db.Store
}

func (s service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	return map[string]map[string]kit.HTTPEndpoint{
		endpointHealth: {
			http.MethodGet: {
				Endpoint: s.getHealth,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
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
				Encoder:  kit.EncodeProtoResponse,
			},
		},
		endpointAccount: {
			http.MethodPost: {
				Decoder:  decodePostAccount,
				Endpoint: s.postAccount,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
		endpointAccounts: {
			http.MethodGet: {
				Decoder:  decodeGetAccounts,
				Endpoint: s.getAccounts,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
		endpointContribution: {
			http.MethodPost: {
				Decoder:  decodeContribution,
				Endpoint: s.postContribution,
				Encoder:  kit.EncodeProtoResponse,
			},
			http.MethodGet: {
				Decoder:  decodeGetContribution,
				Endpoint: s.getContribution,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
		endpointGoal: {
			http.MethodPost: {
				Decoder:  decodePostGoal,
				Endpoint: s.postGoal,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
		endpointGoals: {
			http.MethodGet: {
				Decoder:  decodeGetGoals,
				Endpoint: s.getGoals,
				Encoder:  kit.EncodeProtoResponse,
			},
		},
	}
}

// HTTPMiddleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s service) HTTPMiddleware(h http.Handler) http.Handler {
	return server.CORSHandler(h, "localhost:3000")
}

func (s service) HTTPRouterOptions() []kit.RouterOption {
	notFound := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		kit.EncodeProtoResponse(context.Background(), w, kit.NewProtoStatusResponse(&financial.Error{Message: "not found"}, http.StatusNotFound))
	})

	return []kit.RouterOption{
		kit.RouterSelect(""),
		kit.RouterNotFound(notFound),
	}
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
