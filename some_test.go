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
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Проверка статус кода, отличный от 200
	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)

	//Получаем список кафе
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	//Проверяем, что длина полученных кафе равна ожидаемой
	assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenOkAndBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Проверка статус кода, отличный от 200
	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)

	//Проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//Проверка статус кода, отличный от 400
	status := responseRecorder.Code
	require.Equal(t, http.StatusBadRequest, status)

	//Проверка ожидаемого ответа сервера
	expected := "count missing"
	assert.Equal(t, expected, responseRecorder.Body.String())
}
