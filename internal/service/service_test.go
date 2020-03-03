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

func TestService(t *testing.T) {
	server := kit.NewServer(New())

	var user financial.UserResponse
	request := &financial.PostUserRequest{
		Data: &financial.PostUserData{
			Email: "service_test@example.com",
		},
	}
	res, err := protoRequest(server, endpointUser, http.MethodPost, request, &user)
	if err != nil {
		t.Fatal("It should not return an error on POST /user.", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatal("It should return Status 201.", res.StatusCode)
	}
	if user.Id == 0 {
		t.Fatal("It should return a non-zero Id.", user.Id)
	}
	if user.Email != request.Data.Email {
		t.Fatal("It should return the email that was given.", request.Data.Email, user.Email)
	}

	var user2 financial.UserResponse
	request2 := &financial.PostUserRequest{
		Data: &financial.PostUserData{
			Email: "service_test2@example.com",
		},
	}
	res, err = protoRequest(server, endpointUser, http.MethodPost, request2, &user2)
	if err != nil {
		t.Fatal("It should not return an error on POST /user.", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatal("It should return Status 201.", res.StatusCode)
	}
	if user.Id == user2.Id {
		t.Fatal("It should not create two users with the same Id.", user.Id)
	}

	var responseErr financial.Error
	res, err = protoRequest(server, endpointUser, http.MethodPost, request, &responseErr)
	if err != nil {
		t.Fatal("It should not return an error on POST /user.", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("It should return status 400.", res.StatusCode)
	}
	if responseErr.Message == "" || responseErr.Message != messageAlreadyExists {
		t.Fatal("It should return an already exists message.", responseErr.Message)
	}

	missingEmailRequest := &financial.PostUserRequest{}
	res, err = protoRequest(server, endpointUser, http.MethodPost, missingEmailRequest, &responseErr)
	if err != nil {
		t.Fatal("It should not return an error on POST /user.", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("It should return status 400.", res.StatusCode)
	}
	if responseErr.Message == "" || responseErr.Message != messageMissingEmail {
		t.Fatal("It should return a missing email message.", responseErr.Message)
	}

	missingEmailRequest = &financial.PostUserRequest{
		Data: &financial.PostUserData{},
	}
	res, err = protoRequest(server, endpointUser, http.MethodPost, missingEmailRequest, &responseErr)
	if err != nil {
		t.Fatal("It should not return an error on POST /user.", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("It should return status 400.", res.StatusCode)
	}
	if responseErr.Message == "" || responseErr.Message != messageMissingEmail {
		t.Fatal("It should return a missing email message.", responseErr.Message)
	}

	var user3 financial.UserResponse
	getRequest := &financial.GetUserRequest{
		Data: &financial.GetUserData{
			Email: user.Email,
		},
	}
	res, err = protoRequest(server, endpointUser, http.MethodGet, getRequest, &user3)
	if err != nil {
		t.Fatal("It should not return an error on GET /user.", err)
	}
	if !proto.Equal(&user3, &user) {
		t.Fatal("It should return the same user that was created.", user3, user)
	}

	getRequestNotFound := &financial.GetUserRequest{
		Data: &financial.GetUserData{
			Email: "not even a real email",
		},
	}
	res, err = protoRequest(server, endpointUser, http.MethodGet, getRequestNotFound, &responseErr)
	if err != nil {
		t.Fatal("It should not return an error on GET /user.", err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Fatal("It should return an http status 404.", res.StatusCode)
	}
	if responseErr.Message == "" || responseErr.Message != messageNotFound {
		t.Fatal("It should return a not found message.", responseErr.Message)
	}

	var accountResponse financial.PostAccountResponse
	postAccountRequest := &financial.PostAccountRequest{
		Data: &financial.Account{
			Name:    "Savings",
			UserId:  user3.Id,
			Balance: 27585.45,
			Mode:    financial.Mode_INVESTMENTS,
		},
	}
	res, err = protoRequest(server, endpointAccount, http.MethodPost, postAccountRequest, &accountResponse)
	if err != nil {
		t.Fatal("It should not return an error on POST /account.", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatal("It should return http status created.", res.StatusCode)
	}
	expected := &financial.PostAccountResponse{
		Id: 1,
	}
	if !proto.Equal(&accountResponse, expected) {
		t.Fatal("It should return a correctly created account.", &accountResponse, expected)
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
	res, err = protoRequest(server, endpointAccount, http.MethodPost, postAccountRequest2, &accountResponse2)
	if err != nil {
		t.Fatal("It should not return an error for POST /account", err)
	}
	expected = &financial.PostAccountResponse{
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
