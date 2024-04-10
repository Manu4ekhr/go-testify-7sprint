package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerIncorrectCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := "moscow"
	actual := req.URL.Query().Get("city")

	assert.Equal(t, expected, actual)
}
