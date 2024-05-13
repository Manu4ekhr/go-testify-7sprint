package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupRequestAndRecorder(method, url string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, nil)
	recorder := httptest.NewRecorder()
	return req, recorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, recorder := setupRequestAndRecorder("GET", "/?count=10&city=moscow")
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	require.NotEmpty(t, recorder.Body.String())
	list := strings.Split(recorder.Body.String(), ",")
	assert.Equal(t, totalCount, len(list))
}

func TestMainHandlerWhenCityNotCorrect(t *testing.T) {
	req, recorder := setupRequestAndRecorder("GET", "/?count=4&city=unknown")
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	require.NotEmpty(t, recorder.Body.String())
	list := strings.Split(recorder.Body.String(), ",")
	assert.Contains(t, recorder.Body.String(), list[0], "wrong city value")
}
