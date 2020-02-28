package service

import (
	context "context"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/justindfuller/financial"
)

func decodePostAccount(ctx context.Context, req *http.Request) (interface{}, error) {
	/* var request financial.PostAccountRequest

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&financial.Error{Message: err.Error()}, http.StatusBadRequest)
	}

	err = proto.Unmarshal(body, &request)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&financial.Error{Message: err.Error()}, http.StatusBadRequest)
	}

	return &request, nil */
	return nil, nil
}

var accountId int64

func (s *service) postAccount(ctx context.Context, req interface{}) (interface{}, error) {
	accountId += 1
	return kit.NewProtoStatusResponse(&financial.PostAccountResponse{
		Id: accountId,
	}, http.StatusCreated), nil
}
