package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/golang/protobuf/proto"
	"github.com/justindfuller/financial"
)

func makeRequest(server *kit.Server, endpoint, httpMethod string, request, response proto.Message) (*http.Response, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(httpMethod, endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, r)
	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return w.Result(), nil
}

func TestService(t *testing.T) {
	server := kit.NewServer(New())

	var user financial.UserResponse
	request := &financial.PostUserRequest{
		Data: &financial.PostUserData{
			Email: "service_test@example.com",
		},
	}
	res, err := makeRequest(server, endpointUser, http.MethodPost, request, &user)
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
	res, err = makeRequest(server, endpointUser, http.MethodPost, request2, &user2)
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
	res, err = makeRequest(server, endpointUser, http.MethodPost, request, &responseErr)
	if err != nil {
		t.Fatal("It should not return an error on POST /user.", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("It should return status 400.", res.StatusCode)
	}
	if responseErr.Message == "" || responseErr.Message != messageMissingEmail {
		t.Fatal("It should return a missing email message.", messageMissingEmail)
	}

	var user3 financial.UserResponse
	getRequest := &financial.GetUserRequest{
		Data: &financial.GetUserData{
			Email: user.Email,
		},
	}
	res, err = makeRequest(server, endpointUser, http.MethodGet, getRequest, &user3)
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
	res, err = makeRequest(server, endpointUser, http.MethodGet, getRequestNotFound, &responseErr)
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
		Data: &financial.PostAccountData{
			Name:    "Savings",
			UserId:  user3.Id,
			Balance: 27585.45,
			Mode:    financial.Mode_INVESTMENTS,
		},
	}
	res, err = makeRequest(server, endpointAccount, http.MethodPost, postAccountRequest, &accountResponse)
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
		Data: &financial.PostAccountData{
			Name:    "Credit Card",
			UserId:  user3.Id,
			Balance: 3496.45,
			Mode:    financial.Mode_DEBT,
		},
	}
	res, err = makeRequest(server, endpointAccount, http.MethodPost, postAccountRequest2, &accountResponse2)
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
		Accounts: []*financial.GetAccountsResponse_AccountsMessage{
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
	res, err = makeRequest(server, endpointAccounts, http.MethodGet, getAccountsRequest, &getAccountsResponse)
	if err != nil {
		t.Fatal("It should not return an error for GET /accounts", err)
	}
	if !proto.Equal(&getAccountsResponse, expectedGetAccountsResponse) {
		t.Fatal("It should return the expected getAccountsResponse", &getAccountsResponse, expectedGetAccountsResponse)
	}
}
