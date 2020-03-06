package service

import (
	"net/http"
	"testing"

	"github.com/justindfuller/financial"
)

func TestHealth(t *testing.T) {
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
	}

	runTests(t, tests)
}
