package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandler(t *testing.T) {
	t.Run("when request ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)
		status := responseRecorder.Code

		assert.Equal(t, status, http.StatusOK)
	})

	t.Run("when city not allowed ok", func(t *testing.T) {
		expected := "wrong city value"
		req := httptest.NewRequest("GET", "/cafe?count=2&city=NewYork", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)
		status := responseRecorder.Code
		response := responseRecorder.Body.String()

		assert.Equal(t, status, http.StatusBadRequest)
		assert.Equal(t, response, expected)
	})

	t.Run("when count more than total", func(t *testing.T) {
		totalCount := 4
		req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)
		status := responseRecorder.Code
		response := responseRecorder.Body.String()
		cafesSlice := strings.Split(response, ",")

		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, totalCount, len(cafesSlice))
	})
}
