package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	target := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount+1)
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}

func TestMainHandlerWhenCityDoesNotExist(t *testing.T) {
	totalCount := 4
	target := fmt.Sprintf("/cafe?count=%d&city=%s", totalCount, "Yelabuga")
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenEverythingIsOK(t *testing.T) {
	totalCount := 4
	target := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount)
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCountIsMissing(t *testing.T) {
	target := fmt.Sprintf("/cafe?count=%s&city=moscow", "")
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestMainHandlerWhenWrongCountValue(t *testing.T) {
	target := fmt.Sprintf("/cafe?count=%s&city=moscow", "test")
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
