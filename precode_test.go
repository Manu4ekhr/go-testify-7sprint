package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// тест на корректность запроса и непустое тело ответа
// тут если статус не 200, думаю нет смысла тело смотреть.
// поэтому статус тестил через require
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// ожидаемый статус (200)
	status := http.StatusOK

	// фактический статус
	answerStatus := responseRecorder.Code

	// проверка статуса
	require.Equal(t, status, answerStatus)

	// проверка, что тело ответа не пустое
	require.NotEmpty(t, responseRecorder.Body)
}

// тест на город
// если статус не 400, но решил, что сообщение можно не смотреть.
// статус тестил через require
func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=klin&count=10", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// фактический стату ответа
	answerStatus := responseRecorder.Code

	// ожидаемый статус ответа (400)
	status := http.StatusBadRequest

	// проверка, что код ответа 400
	require.Equal(t, status, answerStatus)

	// фактическое сообщение
	answerMessage := responseRecorder.Body.String()

	// ожидаемое сообщение
	message := "wrong city value"

	// проверка на правильность сообщения
	assert.Equal(t, message, answerMessage)
}

// тест на количество кафешек: запрашиавем больше, чем есть. Вернётся весь список.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	// ожидаемое кол-во кафешек в Москве
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// ответ в виде строки
	answer := responseRecorder.Body.String()

	// фактическое количество
	answerCount := len(strings.Split(answer, ","))

	// проверка, что вернулось ожидаемое кол-во кафешек
	assert.Equal(t, totalCount, answerCount)
}
