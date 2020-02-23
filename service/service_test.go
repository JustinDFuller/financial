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
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected Status %d Got %d", http.StatusCreated, res.StatusCode)
	}
	if user.Id == 0 {
		t.Fatal("Expected a user ID but got zero.")
	}
	if user.Email != request.Data.Email {
		t.Fatal("Got unexpected email.", user.Email)
	}

	var user2 UserResponse
	res, err = makeRequest(server, endpointUser, http.MethodPost, &PostUserRequest{
		Data: &PostUserData{
			Email: "service_test2@example.com",
		},
	}, &user2)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected Status %d Got %d", http.StatusCreated, res.StatusCode)
	}
	if user.Id == user2.Id {
		t.Fatal("Two users created with the same ID.", user.Id)
	}

	var responseErr Error
	res, err = makeRequest(server, endpointUser, http.MethodPost, request, &responseErr)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected Status %d Got %d", http.StatusBadRequest, res.StatusCode)
	}

	var user3 UserResponse
	res, err = makeRequest(server, endpointUser, http.MethodGet, &GetUserRequest{
		Data: &GetUserData{
			Email: user.Email,
		},
	}, &user3)
	if err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(&user3, &user) {
		t.Fatal("Expected user3 and user to be the same.")
	}
}
