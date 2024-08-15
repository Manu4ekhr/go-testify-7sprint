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
	req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "The response code is not 200 OK")

	responseBody := responseRecorder.Body.String()
	cafesList := strings.Split(responseBody, ",")

	assert.Len(t, cafesList, totalCount, "The answer does not contain all cities")
}

func TestMainHandlerWhenRequestInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "The response code is not 200 OK")

	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCityInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=gotham", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "The response code is not 400 Bad Request")

	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}
