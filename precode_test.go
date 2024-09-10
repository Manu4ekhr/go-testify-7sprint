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
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	cafeList := strings.Split(body, ",")
	assert.Len(t, cafeList, totalCount, "Response should contain %d cafes when count is %d", totalCount, len(cafeList))

}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=3", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status: %d", http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body, "Response body should not be empty")

}

func TestMainHandlerWhenUnknownCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=kazan&count=3", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status: %d", http.StatusBadRequest)
	assert.Equal(t, "wrong city value", responseRecorder.Code, "wrong city value")

}
