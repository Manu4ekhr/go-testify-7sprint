package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=11&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, totalCount, len(strings.Split(responseRecorder.Body.String(), ",")))
}
func TestMainHandlerWhenCityIsIncorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=city", nil)
	city := req.URL.Query().Get("city")
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.NotEqual(t, responseRecorder.Code, http.StatusOK)
	assert.NotContains(t, city, "wrong city value", responseRecorder.Code)
}

func TestMainHandlerWhenRequestIsIncorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)

}
