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

    assert.Equal(t, len(list), totalCount)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=vegas", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)


    require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
    assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenOk(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow",nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusOK)
    assert.NotEmpty(t, responseRecorder.Body)
}
