package service

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
)

var id int64

func decodeUser(ctx context.Context, req *http.Request) (interface{}, error) {
	var user PostUserData

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&Error{Message: err.Error()}, http.StatusBadRequest)
	}

	err = proto.Unmarshal(body, &user)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&Error{Message: err.Error()}, http.StatusBadRequest)
	}

	return &PostUserRequest{Data: &user}, nil
}

func (s *service) postUser(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(*PostUserRequest)
	id += 1
	return kit.NewProtoStatusResponse(&PostUserResponse{
		Id:    id,
		Email: req.Data.Email,
	}, http.StatusCreated), nil
}
