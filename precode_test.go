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
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	exceptedBody := strings.Join(cafeList["moscow"], ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")
	assert.Equal(t, exceptedBody, responseRecorder.Body.String(), "Expected all cafes in response body")
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), len(cafeList["moscow"]), "Expected length of cafes to match total available")
}

func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	exceptedBody := strings.Join(cafeList["moscow"], ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")
	assert.Equal(t, exceptedBody, responseRecorder.Body.String(), "Expected all cafes in response body")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Response body should not be empty")
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=samara", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code 400")
	assert.Equal(t, "wrong city valuee", responseRecorder.Body.String())
}
