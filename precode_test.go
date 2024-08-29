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
	req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	responseBody := responseRecorder.Body.String()
	cafes := strings.Split(responseBody, ",")
	assert.ElementsMatch(t, cafeList["moscow"], cafes, "Все кафе из списка должны быть в ответе")
}

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Статус-код должен быть 200 OK")

	require.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsNotValid(t *testing.T) {
	expectedText := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=3&city=paris", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Статус-код должен быть 400 Bad Request")

	responseBody := responseRecorder.Body.String()
	assert.Equal(t, expectedText, responseBody)
}
