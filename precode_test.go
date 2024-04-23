package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	baseUrl, err := url.Parse("/cafe")
	require.NoError(t, err)

	params := url.Values{}
	params.Add("count", "10")
	params.Add("city", "moscow")

	baseUrl.RawQuery = params.Encode()
	req := httptest.NewRequest("GET", baseUrl.String(), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	baseUrl, err := url.Parse("/cafe")
	require.NoError(t, err)

	params := url.Values{}
	params.Add("count", "3")
	params.Add("city", "moscow")

	baseUrl.RawQuery = params.Encode()
	req := httptest.NewRequest("GET", baseUrl.String(), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
}

func TestMainHandlerWhenCityUncorrect(t *testing.T) {
	errString := "wrong city value"

	baseUrl, err := url.Parse("/cafe")
	require.NoError(t, err)

	params := url.Values{}
	params.Add("count", "3")
	params.Add("city", "ryazan123")

	baseUrl.RawQuery = params.Encode()
	req := httptest.NewRequest("GET", baseUrl.String(), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	body := responseRecorder.Body.String()
	assert.Equal(t, errString, body)
}
