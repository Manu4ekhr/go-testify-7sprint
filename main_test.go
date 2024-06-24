package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMainHandlerWhenCountMoreThanTotal тестирует сценарий, когда количество запрашиваемых кафе превышает общее количество кафе.

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/cafe?count="+strconv.Itoa(totalCount+1)+"&city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, totalCount, len(strings.Split(responseRecorder.Body.String(), ",")))
}

// TestMainHandlerWhenCityNotSupported тестирует сценарий, когда параметр city не поддерживается.
func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=1&city=spb", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

// TestMainHandlerWhenRequestIsCorrect тестирует сценарий, когда запрос является корректным.ю

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, 2, len(strings.Split(responseRecorder.Body.String(), ",")))
}
