package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	moreThenTotal := len(cafeList[city]) + 1
	url := fmt.Sprintf("/cafe?count=%d&city=%s", moreThenTotal, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	data := strings.Join(cafeList[city], ",")
	statusCode := responseRecorder.Code
	body := responseRecorder.Body.String()

	assert.Equal(t, statusCode, http.StatusOK)
	assert.NotEmpty(t, body)
	assert.Equal(t, body, data)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	city := "TLT"
	totalCount := len(cafeList[city])
	url := fmt.Sprintf("/cafe?count=%d&city=%s", totalCount, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	data := "wrong city value"
	statusCode := responseRecorder.Code
	body := responseRecorder.Body.String()

	assert.Equal(t, statusCode, http.StatusBadRequest)
	assert.Equal(t, body, data)
}
func TestMainHandlerWhenRightTotal(t *testing.T) {
	city := "moscow"
	totalCount := len(cafeList[city])
	url := fmt.Sprintf("/cafe?count=%d&city=%s", totalCount, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	data := strings.Join(cafeList[city], ",")
	statusCode := responseRecorder.Code
	body := responseRecorder.Body.String()

	assert.Equal(t, statusCode, http.StatusOK)
	assert.NotEmpty(t, body)
	assert.Equal(t, body, data)
}
