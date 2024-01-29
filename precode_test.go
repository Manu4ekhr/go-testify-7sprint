package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getCountAndCity(count string, city string) (*http.Request, *httptest.ResponseRecorder) {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=%s&city=%s", count, city), nil), httptest.NewRecorder()
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, responseRecorder := getCountAndCity("5", "moscow")
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, []byte(strings.Join(cafeList["moscow"][:totalCount], ",")), responseRecorder.Body.Bytes())

}

func TestMainHandlerWhenOk(t *testing.T) {
	req, responseRecorder := getCountAndCity("2", "moscow")
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestWhenWrongCity(t *testing.T) {
	req, responseRecorder := getCountAndCity("2", "london")
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
}
