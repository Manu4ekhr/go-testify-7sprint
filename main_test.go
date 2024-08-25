package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenRequestValid(t *testing.T) {
	count := 3
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=%d&city=moscow", count), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
	expected := strings.Join(cafeList["moscow"][:count], ", ")
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityWrong(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=krasnodar", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
	expected := "wrong city value"
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMissing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
	expected := "count missing"
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountZero(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=0&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
	expected := "wrong count value"
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountLessThanOne(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=-2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
	expected := "wrong count value"
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanCafes(t *testing.T) {
	count := 10
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=%d&city=moscow", count), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
	expected := strings.Join(cafeList["moscow"], ", ")
	assert.Equal(t, expected, responseRecorder.Body.String())
}
