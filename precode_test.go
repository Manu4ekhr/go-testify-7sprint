package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerCorrectRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	require.NoError(t, err, "creating request should not result in an error")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "expecting status code to be 200")

	assert.NotEmpty(t, responseRecorder.Body.String(), "response body should not be empty")
}

// Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerUnsupportedCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=1&city=unknown", nil)
	require.NoError(t, err, "creating request should not result in an error")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expecting status code to be 400")

	assert.Contains(t, responseRecorder.Body.String(), "wrong city value", "response body should contain 'wrong city value'")
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerCountMoreThanAvailable(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err, "creating request should not result in an error")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "expecting status code to be 200")

	assert.NotEmpty(t, responseRecorder.Body.String(), "response body should not be empty")

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), 4, "expecting 4 cafes in the response because that's all that is available")
}
