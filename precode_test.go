precode_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountIsOK(t *testing.T) {
req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)
handler := http.HandlerFunc(main.MainHandle)
assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
assert.Equal(t, "count missing", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountIsNotInt(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=aaa", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(main.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong count value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(main.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "city missing", responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsUnknown(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=unknown&count=1", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(main.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

func TestMainHandlerWhenCountIsOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(main.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())

	cafeList := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafeList, 2)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(main.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())

cafeList := strings.Split(responseRecorder.Body.String(), ",")
assert.Len(t, cafeList, 4)
}
