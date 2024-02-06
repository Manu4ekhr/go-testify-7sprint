package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmpty(t, responseRecorder.Code)

	//Проверка, что статус 200
	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)
}

func TestMainHandWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=mocsow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//Проверка, что мы отправили нормальный город
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.NotEmpty(t, responseRecorder.Code)

	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEqual(t, "wrong city value", body)

	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)

}
