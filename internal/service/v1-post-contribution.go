package service

import (
	context "context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

func decodeContribution(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.PostContributionRequest

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

func (s *service) postContribution(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*financial.PostContributionRequest)

	if req.Data == nil || req.Data.AccountId == 0 {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageInvalidEntity,
		}, http.StatusBadRequest), nil
	}

	id, err := s.db.CreateContributionByAccountId(req.Data.AccountId, req.Data)
	if err != nil {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageAlreadyExists,
		}, http.StatusBadRequest), nil
	}

	return kit.NewProtoStatusResponse(&financial.PostAccountResponse{
		Id: id,
	}, http.StatusCreated), nil
}
