package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Если в параметре `count` указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Запрос сформирован корректно
	// сервис возвращает код ответа 200
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	body := responseRecorder.Body.String()
	// Тело ответа не пустое
	require.NotEmpty(t, body)
	list := strings.Split(body, ",")
	assert.Equal(t, totalCount, len(list))
}

// Город, который передаётся в параметре `city`, не поддерживается.
func TestMainHandlerWhenCityNotCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/?count=4&city=unknown", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
	list := strings.Split(body, ",")
	if !assert.Contains(t, body, list[0]) {
		t.Error("wrong city value")
	}
}
