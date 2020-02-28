package service

import (
	context "context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

func decodeGetAccounts(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.GetAccountsRequest

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

func (s *service) getAccounts(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*financial.GetAccountsRequest)

	if accounts, ok := accountsByUserId[r.Data.UserId]; ok {
		return kit.NewProtoStatusResponse(&financial.GetAccountsResponse{
			Accounts: accounts,
		}, http.StatusOK), nil
	}

	// TODO.
	// return kit.NewProtoStatusResponse(&financial.GetAccountsResponse{}, http.StatusNotFound), nil
	return nil, nil
}
