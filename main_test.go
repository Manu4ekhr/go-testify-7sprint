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
    totalCount := len(cafeList["moscow"])

    req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    list := strings.Split(responseRecorder.Body.String(), ",")

    require.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)
    assert.Equal(t, totalCount, len(list))
}


func TestMainHandlerWhenWrongCityValue(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=Imperial-City", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusBadRequest, responseRecorder.Result().StatusCode)
    assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}
func TestMainHandlerWhenCorrectRequest(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    body := responseRecorder.Body.String()

    require.Equal(t, responseRecorder.Code, http.StatusOK)
    assert.NotEmpty(t, body)
}
