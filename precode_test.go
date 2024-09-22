package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWithCorrectRequest(t *testing.T) {
	city := "moscow"
	count := 4
	req := httptest.NewRequest("GET", "/cafe?city="+city+"&count="+strconv.Itoa(count), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWithUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=invalidcity&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 10
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count="+strconv.Itoa(totalCount), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	expectedCafes := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedCafes, responseRecorder.Body.String())
}
