package service

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
	"github.com/justindfuller/financial/internal/db"
)

const messageMissingEmail = "Missing Email"
const messageAlreadyExists = "Already Exists"

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

	if req.Data == nil {
		return kit.NewProtoStatusResponse(&financial.Error{Message: messageMissingEmail}, http.StatusBadRequest), nil
	}

	userId, err := s.db.CreateUserByEmail(req.Data.Email)
	if err == db.ErrAlreadyExists {
		return kit.NewProtoStatusResponse(&financial.Error{Message: messageAlreadyExists}, http.StatusBadRequest), nil
	}
	if err == db.ErrMissingEmail {
		return kit.NewProtoStatusResponse(&financial.Error{Message: messageMissingEmail}, http.StatusBadRequest), nil
	}

	return kit.NewProtoStatusResponse(&financial.UserResponse{
		Id:    userId,
		Email: req.Data.Email,
	}, http.StatusCreated), nil
}
