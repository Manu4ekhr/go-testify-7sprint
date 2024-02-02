package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func serveHttpRequest(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	responseRecorder := serveHttpRequest("/cafe?count=2&city=moscow")
	expectedCode := http.StatusOK
	expectedCount := 2

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}

func TestMainHandlerWhenCityCorrect(t *testing.T) {
	expectedBody := `wrong city value`
	responseRecorder := serveHttpRequest("/cafe?count=2&city=omsk")
	expectedCode := http.StatusBadRequest

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Equal(t, expectedBody, body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	responseRecorder := serveHttpRequest("/cafe?count=10&city=moscow")
	expectedCode := http.StatusOK
	expectedCount := 4

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}
