package service

import (
	context "context"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
)

func decodeContribution(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func (s *service) postContribution(ctx context.Context, r interface{}) (interface{}, error) {
	return kit.NewProtoStatusResponse(nil, http.StatusCreated), nil
}
