package main_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mainHandle(w http.ResponseWriter, r *http.Request) {
	// Логика обработки запроса
}

func TestMainHandlerWithValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=kazan&count=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "ожидался статус ОК")
	assert.NotEmpty(t, responseRecorder.Body.String(), "ожидалось непустое тело ответа")
}

func TestMainHandlerWithInvalidCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=invalid&count=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "ожидался статус Bad Request")
	assert.Equal(t, "неверное значение города", responseRecorder.Body.String(), "ожидалось сообщение об ошибке")
}

func TestMainHandlerWithCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=kazan&count=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "ожидался статус ОК")

	expectedBody := strings.Join(cafeList["kazan"], ",")
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "ожидались все кафе")
}
