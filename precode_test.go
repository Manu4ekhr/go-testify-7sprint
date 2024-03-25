package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
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

	status := responseRecorder.Code
	statusExpected := http.StatusOK

	answer := responseRecorder.Body

	require.Equal(t, statusExpected, status)
	assert.NotEmpty(t, answer)
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusBadRequest

	answer := responseRecorder.Body.String()
	answerContains := "count missing"

	require.Equal(t, statusExpected, status)
	assert.Contains(t, answer, answerContains)
}

func TestMainHandlerWhenWrongCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusBadRequest

	answer := responseRecorder.Body.String()
	answerContains := "wrong count value"

	require.Equal(t, statusExpected, status)
	assert.Contains(t, answer, answerContains)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=UnExistCity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusBadRequest

	answer := responseRecorder.Body.String()
	answerContains := "wrong city value"

	require.Equal(t, statusExpected, status)
	assert.Contains(t, answer, answerContains)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := len(cafeList["moscow"])
	countMoreTotal := strconv.Itoa(totalCount + 10)
	req := httptest.NewRequest("GET", "/cafe?count="+countMoreTotal+"&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusOK

	answer := responseRecorder.Body.String()
	list := strings.Split(answer, ",")
	countExpected := totalCount

	require.Equal(t, statusExpected, status)
	assert.Len(t, list, countExpected)
}
