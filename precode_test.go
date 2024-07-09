package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Статус 200
func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if !assert.Equal(t, http.StatusOK, responseRecorder.Code) {
		t.Fatalf("Expected %d, got %d", http.StatusOK, responseRecorder.Code)
	}
	if !assert.NotEmpty(t, responseRecorder.Body.String()) {
		t.Fatalf("response body is empty")
	}
}

// Другой город
func TestMainHandlerUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=saintpetersburg&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if !assert.Equal(t, http.StatusBadRequest, responseRecorder.Code) {
		t.Fatalf("Expected %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
	if !assert.Equal(t, "wrong city value", responseRecorder.Body.String()) {
		t.Fatalf("Waiting for: wrong city value, got %s", responseRecorder.Body.String())
	}
}

// Запрос с count больше
func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	list := strings.Split(responseRecorder.Body.String(), ",")
	if !assert.Len(t, list, totalCount) {
		t.Fatalf("Waiting for all %d cafe's", totalCount)
	}
}
