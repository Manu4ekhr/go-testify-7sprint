package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCorrectRequest(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) //

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	assert.NotEmpty(t, responseRecorder.Body.String())

}

func TestMainHandlerBadCitryRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=saratov", nil) //

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) //

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, len(list), totalCount)
}
