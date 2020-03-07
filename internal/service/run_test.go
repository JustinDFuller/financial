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
)

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

func runTests(t *testing.T, tests []*test) {
	server := kit.NewServer(New())

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
