package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/golang/protobuf/proto"
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

	var user UserResponse
	request := &PostUserRequest{
		Data: &PostUserData{
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

	var user2 UserResponse
	request2 := &PostUserRequest{
		Data: &PostUserData{
			Email: "service_test2@example.com",
		},
	}
	res, err = makeRequest(server, endpointUser, http.MethodPost, request2, &user2)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatal("It should return Status 201.", res.StatusCode)
	}
	if user.Id == user2.Id {
		t.Fatal("It should not create two users with the same Id.", user.Id)
	}

	var responseErr Error
	res, err = makeRequest(server, endpointUser, http.MethodPost, request, &responseErr)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("It should return status 400.", res.StatusCode)
	}
	if responseErr.Message == "" || responseErr.Message != messageMissingEmail {
		t.Fatal("It should return a missing email message.", messageMissingEmail)
	}

	var user3 UserResponse
	getRequest := &GetUserRequest{
		Data: &GetUserData{
			Email: user.Email,
		},
	}
	res, err = makeRequest(server, endpointUser, http.MethodGet, getRequest, &user3)
	if err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(&user3, &user) {
		t.Fatal("It should return the same user that was created.", user3, user)
	}

	getRequestNotFound := &GetUserRequest{
		Data: &GetUserData{
			Email: "not even a real email",
		},
	}
	res, err = makeRequest(server, endpointUser, http.MethodGet, getRequestNotFound, &responseErr)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Fatal("It should return an http status 404.", res.StatusCode)
	}
	if responseErr.Message == "" || responseErr.Message != messageNotFound {
		t.Fatal("It should return a not found message.", responseErr.Message)
	}
}
