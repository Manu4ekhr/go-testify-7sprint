package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandleWhenRequestIsIncorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusOK, responseRecorder.Code) // равно
	assert.NotEmpty(t, responseRecorder.Body.String())    // не пусто
}

func TestMainHandleWhenCityIsIncorrect(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?city", nil)
	city := req.URL.Query().Get("city")
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.NotEqual(t, responseRecorder.Code, http.StatusOK)               // не равно
	assert.NotContains(t, city, "wrong city value", responseRecorder.Code) // содержание
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	countCafe := len(strings.Split(responseRecorder.Body.String(), ","))
	assert.Equal(t, countCafe, totalCount) // равно
}
