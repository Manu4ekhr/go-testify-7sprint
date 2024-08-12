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
	req := httptest.NewRequest("GET", "/cafe?count=9&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusCode := responseRecorder.Code
	require.NotEqual(t, 0, statusCode)
	require.Equal(t, http.StatusOK, statusCode)

	responseBody := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, responseBody, totalCount)
}

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusCode := responseRecorder.Code

	require.Equal(t, http.StatusOK, statusCode)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	const WRONH_CITY_ERROR = "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=2&city=rostov", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusCode := responseRecorder.Code

	require.Equal(t, http.StatusBadRequest, statusCode)
	require.Equal(t, WRONH_CITY_ERROR, responseRecorder.Body.String())
}
