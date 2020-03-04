package service

import (
	context "context"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/justindfuller/financial"
)

func (s *service) getHealth(ctx context.Context, req interface{}) (interface{}, error) {
	return kit.NewProtoStatusResponse(&financial.GetHealthResponse{
		Ok: true,
	}, http.StatusOK), nil
}
