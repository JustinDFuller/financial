package service

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

const messageNotFound = "Not found"

func decodeGetUser(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.GetUserRequest

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

func (s *service) getUser(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(*financial.GetUserRequest)

	user, err := s.store.GetUserByEmail(req.Data.Email)
	if err != nil {
		return kit.NewProtoStatusResponse(&financial.Error{Message: messageNotFound}, http.StatusNotFound), nil
	}

	return kit.NewProtoStatusResponse(&financial.UserResponse{Id: user.Id, Email: user.Email}, http.StatusOK), nil
}
