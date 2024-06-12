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
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "ожидался статус ОК")
	assert.NotEmpty(t, responseRecorder.Body.String(), "ожидалось непустое тело ответа")
}

func TestMainHandlerWithInvalidCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=invalid&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "ожидался статус Bad Request")
	assert.Equal(t, "неверное значение города", responseRecorder.Body.String(), "ожидалось сообщение об ошибке")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=kazan", nil) //
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, len(list), totalCount)
}
