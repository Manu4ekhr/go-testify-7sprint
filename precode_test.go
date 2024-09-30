package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestIsValid(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=moscow&count=4", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=london&count=4", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountIsMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, len(cafeList["moscow"]), len(cafes))
}
