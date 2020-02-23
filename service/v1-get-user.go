package service

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
)

const messageNotFound = "Not found"

func decodeGetUser(ctx context.Context, req *http.Request) (interface{}, error) {
	var request GetUserRequest

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&Error{Message: err.Error()}, http.StatusBadRequest)
	}

	err = proto.Unmarshal(body, &request)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&Error{Message: err.Error()}, http.StatusBadRequest)
	}

	return &request, nil
}

func (s *service) getUser(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(*GetUserRequest)

	if users == nil {
		users = map[string]*UserResponse{}
	}

	if user, ok := users[req.Data.Email]; ok {
		return kit.NewProtoStatusResponse(&UserResponse{Id: user.Id, Email: user.Email}, http.StatusOK), nil
	}

	return kit.NewProtoStatusResponse(&Error{Message: messageNotFound}, http.StatusNotFound), nil
}
