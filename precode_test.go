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
	// здесь нужно создать запрос к сервису
	reqStatusOK := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)

	// здесь нужно добавить необходимые проверки
	handler.ServeHTTP(responseRecorder, reqStatusOK)
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, len(strings.Split(responseRecorder.Body.String(), ",")), totalCount)
}

func TestMainHandlerWhenCityBad(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=tambov", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}
