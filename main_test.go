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
	q := "/cafe?city=moscow&count=4"
	req := httptest.NewRequest(http.MethodGet, q, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	q := "/cafe?city=tyumen&count=4"
	req := httptest.NewRequest(http.MethodGet, q, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	expectedAnswer := "wrong city value"
	assert.Equal(t, responseRecorder.Body.String(), expectedAnswer)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	q := "/cafe?city=moscow&count=5"
	req := httptest.NewRequest(http.MethodGet, q, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body.String())

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), 4)
}
