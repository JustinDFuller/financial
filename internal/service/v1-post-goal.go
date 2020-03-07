package service

import (
	context "context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

func decodePostGoal(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.PostGoalRequest

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

func (s *service) postGoal(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*financial.PostGoalRequest)

	if r.Data == nil || r.Data.UserId == 0 {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageInvalidEntity,
		}, http.StatusBadRequest), nil
	}

	goalId, err := s.db.CreateGoalByUserId(r.Data.UserId, r.Data)
	if err != nil {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageAlreadyExists,
		}, http.StatusBadRequest), nil
	}

	return kit.NewProtoStatusResponse(&financial.PostGoalResponse{
		Id: goalId,
	}, http.StatusCreated), nil
}
