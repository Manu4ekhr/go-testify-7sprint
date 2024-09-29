// Для проверки равенства подойдёт функция Equal().
// Чтобы убедиться, что тело не пустое, можно использовать NotEmpty().
// Проверить, что длина соответствует ожидаемой, можно с помощью Len().
// Если надо немедленно остановить выполнение теста - пакет require.
// Для остальных проверок подойдёт пакет assert.
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// здесь нужно создать запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, list, totalCount)
}
func TestMainHandlerWhenCityWrongValue(t *testing.T) {
	expectedErrorMessage := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=1&city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
	assert.Contains(t, body, expectedErrorMessage)
}
func TestMainHandlerWhenCountNormal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	body := responseRecorder.Body
	assert.NotEmpty(t, body)
	bodyString := responseRecorder.Body.String()
	assert.NotEmpty(t, bodyString)
}
