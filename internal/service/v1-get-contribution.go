package service

import (
	context "context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

func decodeGetContribution(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.GetContributionRequest

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&financial.Error{Message: err.Error()}, http.StatusBadRequest)
	}

	err = proto.Unmarshal(body, &request)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&financial.Error{Message: err.Error()}, http.StatusBadRequest)
	}

	return &request, nil

}

func (s *service) getContribution(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*financial.GetContributionRequest)

	if req.Data == nil || req.Data.AccountId == 0 {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageInvalidEntity,
		}, http.StatusBadRequest), nil

	}

	contribution, err := s.db.GetContributionByAccountId(req.Data.AccountId)
	if err != nil {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageNotFound,
		}, http.StatusNotFound), nil
	}

	return kit.NewProtoStatusResponse(contribution, http.StatusOK), nil
}
