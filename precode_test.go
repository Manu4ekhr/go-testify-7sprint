package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createResponseRecoder(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	currentStatus := http.StatusOK
	responseRecorder := createResponseRecoder("/cafe?city=moscow&count=2")

	assert.Equal(t, currentStatus, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCityDontExist(t *testing.T) {
	currentStatus := http.StatusBadRequest
	currentBody := "wrong city value"
	responseRecorder := createResponseRecoder("/cafe?city=krasnoyarsk&count=3")

	assert.Equal(t, currentStatus, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Equal(t, currentBody, body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	currentCount := 4
	responseRecorder := createResponseRecoder("/cafe?city=moscow&count=8")

	assert.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, currentCount)
}
