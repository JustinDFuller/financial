package service

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

const messageMissingEmail = "Missing Email"

var id int64
var users map[string]*financial.UserResponse

func decodePostUser(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.PostUserRequest

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

func (s *service) postUser(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(*financial.PostUserRequest)

	if users == nil {
		users = map[string]*financial.UserResponse{}
	}

	if _, ok := users[req.Data.Email]; ok {
		return kit.NewProtoStatusResponse(&financial.Error{Message: messageMissingEmail}, http.StatusBadRequest), nil
	}

	id += 1
	users[req.Data.Email] = &financial.UserResponse{
		Id:    id,
		Email: req.Data.Email,
	}
	return kit.NewProtoStatusResponse(users[req.Data.Email], http.StatusCreated), nil
}
