package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerSuccessfulRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что тело ответа не пустое
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
}

func TestMainHandlerWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=someothercity", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус ответа
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем, что тело ответа содержит ожидаемую ошибку
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что тело ответа не пустое
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)

	// Разделяем ответ на элементы по запятым
	actualResponse := strings.Split(body, ",")

	// Получаем ожидаемый список кафе для города "moscow"
	expectedResponse := cafeList["moscow"]

	// Проверяем, что элементы возвращаемого списка совпадают с ожидаемыми
	assert.ElementsMatch(t, expectedResponse, actualResponse)

	// Проверяем, что длина списка соответствует ожидаемому количеству элементов
	totalCount := len(expectedResponse)
	assert.Len(t, actualResponse, totalCount)
}
