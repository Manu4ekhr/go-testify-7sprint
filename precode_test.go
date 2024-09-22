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

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenOk(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    status := responseRecorder.Code

    require.Equal(t, status, http.StatusOK)
    require.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=Kirov", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    status := responseRecorder.Code
    expected := "wrong city value"

    assert.Equal(t, status, http.StatusBadRequest)
    assert.Equal(t, expected, responseRecorder.Body.String())
}