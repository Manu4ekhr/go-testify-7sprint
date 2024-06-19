package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code) // проверяем код ответа
	body := responseRecorder.Body.String()                // получаем тело ответа
	assert.NotEmpty(t, body)                              //проверяем, что тело не пустое
}

func TestMainHandlerWhenMissingCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=6&city=saratov", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code) // проверяем код ответа, если не 400, завершаем тест

	expected := `wrong city value`
	body := responseRecorder.Body.String()
	assert.Equal(t, expected, body) // сверяем тело ответа с expected
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=8&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code) // проверяем код ответа, если не 200, завершаем тест

	body := responseRecorder.Body.String() // получаем тело ответа
	list := strings.Split(body, ",")       // создаем слайс из тела ответа
	assert.Len(t, list, totalCount)        // сверяем длину
}
