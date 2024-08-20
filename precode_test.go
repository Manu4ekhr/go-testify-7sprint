package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("Get", "/cafe?count=4&city=moscow", nil)
   
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
   
	resp := responseRecorder.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
		return
	}
   
	body := responseRecorder.Body.String()
	trueCount := len(strings.Split(body, ","))
	assert.Equal(t, totalCount, trueCount, "Status code:%d, wrong count value", responseRecorder.Code)
   }

func TestMainHandlerWhenStatusOkAndNotEmpty(t *testing.T) {
	status := http.StatusOK
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusResult := responseRecorder.Code
	resp := responseRecorder.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
		return
	}
	assert.Equal(t, status, statusResult, "expected status code:%d, got %d", http.StatusOK, statusResult) //не знаю что тут можно изменить
	assert.NotEmpty(t, responseRecorder.Body, "body is empty")
}

func TestMainHandlerWhenNotMoscow(t *testing.T) {
	city := "moscow"
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cityResult := req.URL.Query().Get("city")
	assert.Equal(t, city, cityResult, "Status code:%d, wrong city value", responseRecorder.Code)
}