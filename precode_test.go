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
    req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // здесь нужно добавить необходимые проверки

    require.Equal(t, responseRecorder.Code, http.StatusOK)

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    require.NotEmpty(t, list)
    assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenOk(t *testing.T) {
    expectedLen := 2
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusOK)

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    require.NotEmpty(t, list)
    assert.Len(t, list, expectedLen)

}

func TestMainHandlerWhenWrongCity(t *testing.T) {
    expectedError := "wrong city value"
    req := httptest.NewRequest("GEt", "/cafe?count=1&city=novgorod", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)

    body := responseRecorder.Body.String()
    assert.Equal(t, body, expectedError)

}