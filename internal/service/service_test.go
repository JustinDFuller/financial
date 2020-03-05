package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
		return nil, nil, fmt.Errorf("Unable to read body. %w", err)
	}

	return w.Result(), body, nil
}

func protoRequest(server *kit.Server, endpoint, httpMethod string, request, response proto.Message, requestHeaders map[string]string) (*http.Response, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	result, body, err := httpRequest(server, endpoint, httpMethod, data, requestHeaders)
	if err != nil {
		return nil, fmt.Errorf("Error making http request. %w", err)
	}

	err = proto.Unmarshal(body, response)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling proto response. %w", err)
	}

	return result, nil
}

type test struct {
	name            string
	endpoint        string
	httpMethod      string
	statusCode      int
	request         proto.Message
	response        proto.Message
	expected        proto.Message
	requestHeaders  map[string]string
	responseHeaders map[string]string
}

func TestService(t *testing.T) {
	server := kit.NewServer(New())

	tests := []*test{
		{
			name:       "GET /health",
			endpoint:   endpointHealth,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request:    &financial.GetHealthRequest{},
			response:   &financial.GetHealthResponse{},
			expected: &financial.GetHealthResponse{
				Ok: true,
			},
			requestHeaders: map[string]string{
				"origin": "https://financial-calculator.glitch.me",
			},
			responseHeaders: map[string]string{
				"Access-Control-Allow-Origin": "https://financial-calculator.glitch.me",
			},
		},
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
			name:       "POST /account savings",
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
		{
			name:       "POST /account credit card",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostAccountRequest{
				Data: &financial.Account{
					Name:    "Credit Card",
					UserId:  2,
					Balance: 3496.45,
					Mode:    financial.Mode_DEBT,
				},
			},
			response: &financial.PostAccountResponse{},
			expected: &financial.PostAccountResponse{
				Id: 2,
			},
		},
		{
			name:       "GET /accounts",
			endpoint:   endpointAccounts,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request: &financial.GetAccountsRequest{
				Data: &financial.GetAccountsData{
					UserId: 2,
				},
			},
			response: &financial.GetAccountsResponse{},
			expected: &financial.GetAccountsResponse{
				Accounts: []*financial.Account{
					{
						Id:      1,
						Name:    "Savings",
						UserId:  2,
						Balance: 27585.45,
						Mode:    financial.Mode_INVESTMENTS,
					},
					{
						Id:      2,
						Name:    "Credit Card",
						UserId:  2,
						Balance: 3496.45,
						Mode:    financial.Mode_DEBT,
					},
				},
			},
		},
		{
			name:       "GET /accounts not found",
			endpoint:   endpointAccounts,
			httpMethod: http.MethodGet,
			statusCode: http.StatusNotFound,
			request: &financial.GetAccountsRequest{
				Data: &financial.GetAccountsData{
					UserId: 222,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageNotFound,
			},
		},
		{
			name:       "POST /contribution credit card",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostContributionRequest{
				Data: &financial.PostContributionData{
					AccountId: 2,
					Amount:    500,
				},
			},
			response: &financial.PostContributionResponse{},
			expected: &financial.PostContributionResponse{
				Id: 1,
			},
		},
		{
			name:       "POST /contribution savings",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostContributionRequest{
				Data: &financial.PostContributionData{
					AccountId: 1,
					Amount:    500,
				},
			},
			response: &financial.PostContributionResponse{},
			expected: &financial.PostContributionResponse{
				Id: 2,
			},
		},
		{
			name:       "POST /contribution already exists",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostContributionRequest{
				Data: &financial.PostContributionData{
					AccountId: 1,
					Amount:    500,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageAlreadyExists,
			},
		},
		{
			name:       "POST /contribution bad request",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request:    &financial.PostContributionRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /contribution bad request",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostContributionRequest{
				Data: &financial.PostContributionData{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := protoRequest(server, tc.endpoint, tc.httpMethod, tc.request, tc.response, tc.requestHeaders)
			if err != nil {
				t.Fatal("It should not return an error.", err)
			}
			if res.StatusCode != tc.statusCode {
				t.Fatalf("It should return Status %d. Got %d.", tc.statusCode, res.StatusCode)
			}
			if !proto.Equal(tc.response, tc.expected) {
				t.Fatal("It should return the expected response. \n Got:", tc.response, "\n Expected:", tc.expected)
			}
			for key, val := range tc.responseHeaders {
				if got := res.Header.Get(key); got != val {
					t.Fatalf("It should return the expected %s header. Wanted %s. Got %s.", key, val, got)
				}
			}
		})
	}
}
