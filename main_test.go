package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := http.StatusOK                                      // Ожидаемое сообщение о статусе.
	expectedReply := `Мир кофе, Сладкоежка`                        // Ожидаемое тело ответа.
	assert.Equal(t, expected, responseRecorder.Code)               // Проверяем, что статус 200 корректный.
	assert.Equal(t, expectedReply, responseRecorder.Body.String()) // Сравниваем сообщение с ожидаемым.
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=murmansk&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := `wrong city value`                                // Ожидаемое сообщение.
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest) // Проверяем, что статус 400 корректный.
	assert.Equal(t, expected, responseRecorder.Body.String())     // Проверяем, что сообщение правильное.
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	expectedList := `Мир кофе, Сладкоежка, Кофе и завтраки, Сытый студент` // Ожидаемый ответ.
	req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()
	list := strings.Split(body, ", ")

	expected := http.StatusOK                                     // Ожидаемое сообщение о статусе.
	assert.Equal(t, expected, responseRecorder.Code)              // Проверяем, что статус 200 корректный.
	assert.Equal(t, expectedList, responseRecorder.Body.String()) // Проверка правильности ответа.
	assert.Equal(t, totalCount, len(list))                        // Проверка правильной обработки больших count.

}
