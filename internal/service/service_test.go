package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/golang/protobuf/proto"
	"github.com/justindfuller/financial"
)

func httpRequest(server *kit.Server, endpoint, httpMethod string, data []byte, headers map[string]string) (*http.Response, []byte, error) {
	r, err := http.NewRequest(httpMethod, endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, nil, err
	}

	if headers != nil {
		for key, val := range headers {
			r.Header.Add(key, val)
		}
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, r)
	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		return nil, nil, err
	}

	return w.Result(), body, nil
}

func protoRequest(server *kit.Server, endpoint, httpMethod string, request, response proto.Message) (*http.Response, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	result, body, err := httpRequest(server, endpoint, httpMethod, data, nil)
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type test struct {
	name       string
	endpoint   string
	httpMethod string
	statusCode int
	request    proto.Message
	response   proto.Message
	expected   proto.Message
}

func TestService(t *testing.T) {
	server := kit.NewServer(New())

	tests := []*test{
		{
			name:       "/POST user",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostUserRequest{
				Data: &financial.PostUserData{
					Email: "service_test@example.com",
				},
			},
			response: &financial.UserResponse{},
			expected: &financial.UserResponse{
				Id:    1,
				Email: "service_test@example.com",
			},
		},
		{
			name:       "POST /user again",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostUserRequest{
				Data: &financial.PostUserData{
					Email: "service_test2@example.com",
				},
			},
			response: &financial.UserResponse{},
			expected: &financial.UserResponse{
				Id:    2,
				Email: "service_test2@example.com",
			},
		},
		{
			name:       "POST /user already exists",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostUserRequest{
				Data: &financial.PostUserData{
					Email: "service_test2@example.com",
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageAlreadyExists,
			},
		},
		{
			name:       "POST /user missing data",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request:    &financial.PostUserRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageMissingEmail,
			},
		},
		{
			name:       "POST /user missing email",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostUserRequest{
				Data: &financial.PostUserData{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageMissingEmail,
			},
		},
		{
			name:       "GET /user",
			endpoint:   endpointUser,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request: &financial.GetUserRequest{
				Data: &financial.GetUserData{
					Email: "service_test@example.com",
				},
			},
			response: &financial.UserResponse{},
			expected: &financial.UserResponse{
				Id:    1,
				Email: "service_test@example.com",
			},
		},
		{
			name:       "GET /user not found",
			endpoint:   endpointUser,
			httpMethod: http.MethodGet,
			statusCode: http.StatusNotFound,
			request: &financial.GetUserRequest{
				Data: &financial.GetUserData{
					Email: "not even a real email",
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageNotFound,
			},
		},
		{
			name:       "POST /account",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostAccountRequest{
				Data: &financial.Account{
					Name:    "Savings",
					UserId:  2,
					Balance: 27585.45,
					Mode:    financial.Mode_INVESTMENTS,
				},
			},
			response: &financial.PostAccountResponse{},
			expected: &financial.PostAccountResponse{
				Id: 1,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := protoRequest(server, tc.endpoint, tc.httpMethod, tc.request, tc.response)
			if err != nil {
				t.Fatal("It should not return an error.", err)
			}
			if res.StatusCode != tc.statusCode {
				t.Fatalf("It should return Status %d. Got %d.", tc.statusCode, res.StatusCode)
			}
			if !proto.Equal(tc.response, tc.expected) {
				t.Fatal("It should return the expected response. \n Got:", tc.response, "\n Expected:", tc.expected)
			}
		})
	}

	user3 := &financial.UserResponse{
		Id: 2,
	}

	var accountResponse2 financial.PostAccountResponse
	postAccountRequest2 := &financial.PostAccountRequest{
		Data: &financial.Account{
			Name:    "Credit Card",
			UserId:  user3.Id,
			Balance: 3496.45,
			Mode:    financial.Mode_DEBT,
		},
	}
	res, err := protoRequest(server, endpointAccount, http.MethodPost, postAccountRequest2, &accountResponse2)
	if err != nil {
		t.Fatal("It should not return an error for POST /account", err)
	}
	expected := &financial.PostAccountResponse{
		Id: 2,
	}
	if !proto.Equal(expected, &accountResponse2) {
		t.Fatal("It should return the expected PostAccountResponse.", expected, &accountResponse2)
	}

	var getAccountsResponse financial.GetAccountsResponse
	getAccountsRequest := &financial.GetAccountsRequest{
		Data: &financial.GetAccountsData{
			UserId: user3.Id,
		},
	}
	expectedGetAccountsResponse := &financial.GetAccountsResponse{
		Accounts: []*financial.Account{
			{
				Id:      1,
				Name:    "Savings",
				UserId:  user3.Id,
				Balance: 27585.45,
				Mode:    financial.Mode_INVESTMENTS,
			},
			{
				Id:      2,
				Name:    "Credit Card",
				UserId:  user3.Id,
				Balance: 3496.45,
				Mode:    financial.Mode_DEBT,
			},
		},
	}
	res, err = protoRequest(server, endpointAccounts, http.MethodGet, getAccountsRequest, &getAccountsResponse)
	if err != nil {
		t.Fatal("It should not return an error for GET /accounts", err)
	}
	if !proto.Equal(&getAccountsResponse, expectedGetAccountsResponse) {
		t.Fatal("It should return the expected getAccountsResponse", &getAccountsResponse, expectedGetAccountsResponse)
	}

	var postContributionResponse financial.PostContributionResponse
	postContributionRequest := &financial.PostContributionRequest{
		Data: &financial.PostContributionData{
			AccountId: getAccountsResponse.Accounts[0].Id,
			Amount:    500,
		},
	}
	res, err = protoRequest(server, endpointContribution, http.MethodPost, postContributionRequest, &postContributionResponse)
	if err != nil {
		t.Fatal("It should not return an error for POST /contribution", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatal("It should return status created.", res.StatusCode)
	}
}

func TestHealth(t *testing.T) {
	server := kit.NewServer(New())

	expected, _ := json.Marshal("ok")
	headers := map[string]string{
		"origin": "https://financial-calculator.glitch.me",
	}
	res, body, err := httpRequest(server, endpointHealth, http.MethodGet, nil, headers)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(body)) != string(expected) {
		t.Fatal("It should return ok", strings.TrimSpace(string(body)), string(expected))
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("It should return 200", res.StatusCode)
	}
	if res.Header.Get("Access-Control-Allow-Origin") != headers["origin"] {
		t.Fatal("It should return the expected CORS Access-Control-Allow-Origin", res.Header.Get("Access-Control-Allow-Origin"))
	}
}
