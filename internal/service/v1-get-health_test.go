package service

import (
	"net/http"
	"os"
	"testing"

	"github.com/justindfuller/financial"
)

func TestHealth(t *testing.T) {
	os.Setenv("CORS_ALLOWED_ORIGIN", "localhost:3000")
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
				"origin": "http://localhost:3000",
			},
			responseHeaders: map[string]string{
				"Access-Control-Allow-Origin": "http://localhost:3000",
			},
		},
	}

	runTests(t, tests)
}
