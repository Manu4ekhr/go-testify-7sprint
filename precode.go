package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест на корректный запрос: ожидается код 200 и непустое тело ответа
func TestMainHandle_ValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.NotEmpty(t, rr.Body.String(), "Expected non-empty response body")
}

// Тест на неверный город: ожидается код 400 и сообщение об ошибке
func TestMainHandle_InvalidCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=unknowncity", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status code 400")
	assert.Contains(t, rr.Body.String(), "wrong city value", "Expected error message 'wrong city value'")
}

// Тест на превышение количества кафе: возвращаем все доступные кафе
func TestMainHandle_CountExceedsTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.Len(t, strings.Split(rr.Body.String(), "\n"), 4, "Expected all available cafes to be returned")
}
