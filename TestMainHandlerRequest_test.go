package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := http.StatusOK
	actual := responseRecorder.Code

	require.Equal(t, expected, actual)
	require.NotEmpty(t, responseRecorder.Body.String())
}
