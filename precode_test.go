package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotalOne(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(cd)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
}

func TestMainHandlerWhenCityNotCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=kaliningrad", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(cd)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
	expectedErrorMessage := "city not supported"
	assert.Contains(t, body, expectedErrorMessage)

}

func TestMainHandlerWhenCountMoreThanTotalTwo(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(cd)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, len(list), totalCount)
	assert.NotEmpty(t, body)
}
