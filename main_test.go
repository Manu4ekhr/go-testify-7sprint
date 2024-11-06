package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestMainHandlerValidRequest(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
    require.NoError(t, err)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
    assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerInvalidCity(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?count=2&city=unknown", nil)
    require.NoError(t, err)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
    assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil) 
    require.NoError(t, err)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)


    assert.Equal(t, http.StatusOK, responseRecorder.Code)
    

    expectedResponse := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
    assert.Equal(t, expectedResponse, responseRecorder.Body.String())

    assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}