package service

import (
	context "context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

func decodeGetGoals(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.GetGoalsRequest

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

func (s *service) getGoals(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*financial.GetGoalsRequest)

	if r.Data == nil || r.Data.UserId == 0 {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageInvalidEntity,
		}, http.StatusBadRequest), nil

	}

	goals, err := s.db.GetGoalsByUserId(r.Data.UserId)
	if err != nil {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageNotFound,
		}, http.StatusNotFound), nil
	}

	return kit.NewProtoStatusResponse(&financial.GetGoalsResponse{
		Goals: goals,
	}, http.StatusOK), nil
}
