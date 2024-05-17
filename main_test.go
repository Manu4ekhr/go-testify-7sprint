package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Функция для работы с рекордером ответа
func retRequest(t *testing.T, url string, expectedStatus int) string {
	req := httptest.NewRequest("GET", url, nil)
	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(respRec, req)
	require.Equal(t, expectedStatus, respRec.Code)
	return respRec.Body.String()
}
func TestMainHandlerSuccessTestsEverethingWentGood(t *testing.T) {
	body := retRequest(t, "/cafe?count=4&city=moscow", http.StatusOK)
	require.NotEmpty(t, body)
	list := strings.Split(body, ",")

	expectedCount := 4
	assert.Len(t, list, expectedCount)

	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, body)
}
func TestMainHandlerMissingCount(t *testing.T) {
	body := retRequest(t, "/cafe?city=moscow", http.StatusBadRequest)
	assert.Equal(t, "count missing", body)
}
func TestMainHandlerWrongCountValue(t *testing.T) {
	body := retRequest(t, "/cafe?count=fff&city=moscow", http.StatusBadRequest)
	assert.Equal(t, "wrong count value", body)
}
func TestMainHandlerMoreThanAvailableCount(t *testing.T) {
	body := retRequest(t, "/cafe?count=666&city=moscow", http.StatusOK)
	require.NotEmpty(t, body)

	list := strings.Split(body, ",")
	expectedCount := 4
	assert.Len(t, list, expectedCount)

	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, body)
}

func TestMainHandlerWrongCityValue(t *testing.T) {
	body := retRequest(t, "/cafe?count=4&city=Mordor", http.StatusBadRequest)
	assert.Equal(t, "wrong city value", body)
}
