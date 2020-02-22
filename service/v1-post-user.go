package service

import (
	"context"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
)

func decodeUser(context.Context, *http.Request) (request interface{}, err error) {
	return nil, nil
}

func (s *service) postUser(ctx context.Context, request interface{}) (response interface{}, err error) {
	return kit.NewProtoStatusResponse(nil, http.StatusCreated), nil
}
