package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainStatus200(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка на статус 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=notcity", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус 400
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем запись при вводе несуществующего города
	expectedErrorMessage := "wrong city value"
	assert.Equal(t, expectedErrorMessage, responseRecorder.Body.String())
}

func TestMainСities(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка на статус 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// Конвертируем полученную строку в слайс
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	// Проверка на общее количество
	assert.Len(t, list, totalCount)
}
