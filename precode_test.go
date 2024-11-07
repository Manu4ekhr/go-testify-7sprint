package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=3&city=moscow", nil)
	require.NoError(t, err)

	// Создаем рекордер для записи ответа
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем тело ответа
	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки"
	assert.Equal(t, expectedBody, responseRecorder.Body.String())
}

func TestMainHandlerWrongCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=3&city=spb", nil)
	require.NoError(t, err)

	// Создаем рекордер для записи ответа
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус код
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем тело ответа
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerCountMoreThanAvailable(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err)

	// Создаем рекордер для записи ответа
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что вернулось не больше 4 кафе
	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, responseRecorder.Body.String())
}
