package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err, "creating request should not result in an error")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "expecting status code to be 200")

	assert.NotEmpty(t, responseRecorder.Body.String(), "response body should not be empty")

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), 4, "expecting 4 cafes in the response")
}
