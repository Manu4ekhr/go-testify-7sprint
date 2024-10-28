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
	req, _ := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Неверный код ответа")
	require.NotEmpty(t, responseRecorder.Body, "Тело ответа пустое")
}

func TestMainHandlerWrongCity(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cafe?city=london&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Неверный код ответа")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Неверное сообщение об ошибке")
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Неверный код ответа")
	assert.NotEmpty(t, responseRecorder.Body, "Тело ответа пустое")

	cafeNames := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafeNames, 4, "Неверное количество кафе в ответе")
}
