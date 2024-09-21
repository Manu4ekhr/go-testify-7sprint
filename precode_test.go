package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
    req := httptest.NewRequest("GET", "localhost:8080/test?count=2&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusOK, responseRecorder.Code)
    assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
    req := httptest.NewRequest("GET", "localhost:8080/test?city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
    assert.Equal(t, "count missing", responseRecorder.Body.String())
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
    req := httptest.NewRequest("GET", "localhost:8080/test?count=2&city=saint-petersburg", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
    assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "localhost:8080/test?count=5&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusOK, responseRecorder.Code)
    assert.NotEmpty(t, responseRecorder.Body.String())
    list := strings.Split(responseRecorder.Body.String(), ",")
    assert.Len(t, list, totalCount)
}