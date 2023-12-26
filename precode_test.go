package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	q := req.URL.Query()
	q.Add("count", "5")
	q.Add("city", "moscow")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	res := responseRecorder.Body.String()
	arr := strings.Split(res, ",")
	assert.Equal(t, totalCount, len(arr))
}

func TestWrongCity(t *testing.T) {
	expectedError := "wrong city value"
	expectedCode := 400
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	q := req.URL.Query()
	q.Add("count", "4")
	q.Add("city", "omsk")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	resCode := responseRecorder.Code
	require.Equal(t, expectedCode, resCode)
	res := responseRecorder.Body.String()
	require.Equal(t, res, expectedError)
}

func TestCorrectQuery(t *testing.T) {
	expectedCode := 200
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}
	q := req.URL.Query()
	q.Add("count", "3")
	q.Add("city", "moscow")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	resCode := responseRecorder.Code
	require.Equal(t, expectedCode, resCode)
	require.NotEmpty(t, responseRecorder.Body)
}
