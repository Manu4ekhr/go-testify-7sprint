package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerStatusOk(t *testing.T) {
	cafeExpected := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	assert.NotEmpty(t, responseRecorder.Body.String())

	cafe := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, cafe, cafeExpected)
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=kazan", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	expected := `wrong city value`
	assert.Equal(t, responseRecorder.Body.String(), expected)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	cafeExpected := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cafe := strings.Split(responseRecorder.Body.String(), ",")
	lenCafe := len(cafe)

	assert.Equal(t, lenCafe, totalCount)

	assert.Equal(t, cafe, cafeExpected)

}
