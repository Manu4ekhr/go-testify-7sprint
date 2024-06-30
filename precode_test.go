package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	require, err := http.NewRequest(http.MethodGet, "/cafe?city=moscow&count=2", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, require)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerUnsupportedCity(t *testing.T) {
	require, err := http.NewRequest(http.MethodGet, "/cafe?city=unknown&count=2", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, require)

	expected := "wrong city value"
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerCountMoreThanAvailable(t *testing.T) {
	require, err := http.NewRequest(http.MethodGet, "/cafe?city=moscow&count=10", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, require)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
}
