package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	// Тестируем случай, когда запрос корректен
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус-код 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200 OK")

	// Проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerUnsupportedCity(t *testing.T) {
	// Тестируем случай, когда передан неподдерживаемый город
	req := httptest.NewRequest("GET", "/cafe?city=unknown&count=2", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка кода ответа 400 Bad Request
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code 400 Bad Request")

	// Проверка сообщения об ошибке
	expectedError := "wrong city value"
	assert.Equal(t, expectedError, responseRecorder.Body.String(), "Expected error message for unsupported city")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// Тестируем случай, когда count больше, чем доступное количество кафе
	totalCount := 4 // Объявляем totalCount для проверки

	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка кода ответа 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200 OK")

	// Проверка, что количество кафе в ответе равно totalCount
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount, "The number of cafes should match totalCount")
}
