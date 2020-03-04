package service

import (
	context "context"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/justindfuller/financial"
)

func decodeContribution(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func (s *service) postContribution(ctx context.Context, r interface{}) (interface{}, error) {
	id, _ := s.db.CreateContributionByAccountId(0)
	return kit.NewProtoStatusResponse(&financial.PostAccountResponse{
		Id: id,
	}, http.StatusCreated), nil
}
