package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandleWhenRequestIsIncorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//сначала проверяем код ответа
	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)
	//потом содержимое
	body, _ := ioutil.ReadAll(req.Body)
	assert.NotEmpty(t, body)
}

func TestMainHandleWhenCityIsIncorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=city", nil)
	city := req.URL.Query().Get("city")
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code
	//сначала проверяем код ответа
	assert.NotEqual(t, status, http.StatusOK)
	//потом сам город
	assert.NoContains(t, city, "wrong city value", status)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Equal(t, totalCount, len(list))
}
