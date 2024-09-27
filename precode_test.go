package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "Expected non-empty body")

	cafes := strings.Split(body, ",")
	assert.Len(t, cafes, 2, "Expected 2 cafes")
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code 400")

	expected := `count missing`
	assert.Equal(t, expected, responseRecorder.Body.String(), "Expected error message for missing count")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")

	expectedCafes := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "Expected non-empty body")

	cafes := strings.Split(body, ",")
	assert.Len(t, cafes, totalCount, "Expected all available cafes to be returned")

	for i, cafe := range cafes {
		assert.Equal(t, expectedCafes[i], cafe, "Expected correct cafe name")
	}
}