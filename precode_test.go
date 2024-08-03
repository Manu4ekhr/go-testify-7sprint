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

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	//добавить проверку по городу. завершить работу

	assert.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.NotEmpty(t, body)

	require.Len(t, list, totalCount)

}
func TestMainHandlerWhenCityNotMoscow(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=mosc", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	//добавить проверку по городу. завершить работу

	require.Equal(t, status, http.StatusBadRequest)

}
