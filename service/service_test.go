package service

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NYTimes/gizmo/server/kit"
)

func makeRequest(server *kit.Server, endpoint, httpMethod string, body io.Reader) *http.Response {
	r, _ := http.NewRequest(httpMethod, endpoint, body)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, r)
	return w.Result()
}

func TestService(t *testing.T) {
	server := kit.NewServer(New())

	res := makeRequest(server, endpointUser, http.MethodPost, strings.NewReader(`{"name": "Justin"}`))
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected Status %d Got %d", http.StatusCreated, res.StatusCode)
	}

}
