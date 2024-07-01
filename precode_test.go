package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWithValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=kazan&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, http.StatusOK, responseRecorder.Code, "ожидался статус ОК")
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWithInvalidCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=invalid&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal (t, http.StatusBadRequest, responseRecorder.Code, "ожидался статус Bad Request")
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=kazan", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "ожидался статус ОК")

	list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, list, totalCount, "ожидалось правильное количество элементов")
}
