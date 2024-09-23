package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	//totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // Создаём запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
}
