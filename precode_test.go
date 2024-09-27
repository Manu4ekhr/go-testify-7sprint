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
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	responseBody := responseRecorder.Body.String()
	require.NotEmpty(t, responseBody)

	assert.Len(t, strings.Split(responseBody, ","), totalCount)
}

func TestMainHandlerValidRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	responseBody := responseRecorder.Body.String()
	assert.NotEmpty(t, responseBody)

	assert.Len(t, strings.Split(responseBody, ","), 2)
}

func TestMainHandlerCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=paris", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expected := "wrong city value"
	assert.Equal(t, expected, responseRecorder.Body.String())
}
