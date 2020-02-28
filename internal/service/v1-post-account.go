package service

import (
	context "context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
)

func decodePostAccount(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.PostAccountRequest

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

var accountId int64
var accountsByUserId = map[int64][]*financial.Account{}

func (s *service) postAccount(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*financial.PostAccountRequest)
	accountId += 1
	r.Data.Id = accountId

	if accounts, ok := accountsByUserId[r.Data.UserId]; !ok {
		accountsByUserId[r.Data.UserId] = []*financial.Account{r.Data}
	} else {
		accountsByUserId[r.Data.UserId] = append(accounts, r.Data)
	}

	return kit.NewProtoStatusResponse(&financial.PostAccountResponse{
		Id: accountId,
	}, http.StatusCreated), nil
}
