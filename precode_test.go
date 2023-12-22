package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCorrectQuery(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe", nil)
	q := req.URL.Query()
	q.Add("count", "5")
	q.Add("city", "moscow")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
}
func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe", nil)
	q := req.URL.Query()
	q.Add("count", "5")
	q.Add("city", "Rostov-on-Don")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "wrong city value", responseRecorder.Body.String())

}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe", nil)
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

func TestMainHandlerWhenWrongCountValue(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe", nil)
	q := req.URL.Query()
	q.Add("count", "xxx")
	q.Add("city", "moscow")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.NotEmpty(t, "wrong count value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMissing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe", nil)
	q := req.URL.Query()
	q.Add("city", "moscow")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.NotEmpty(t, "count missing", responseRecorder.Body.String())
}