package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOkAndNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) // Создаём запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Проверяем что запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body.String())

}

func TestMainHandlerWhenCityIsNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=rostov", nil) // Создаём запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Проверяем что город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа..
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")

}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // Создаём запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Проверяем что если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
	list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, list, totalCount)
	assert.Equal(t, cafeList["moscow"], list)

}
