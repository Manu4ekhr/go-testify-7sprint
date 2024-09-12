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

	expected := http.StatusOK                                                   // Ожидаемое сообщение о статусе.
	warningMissingCount := `count missing`                                      // Сообщение об отсутствии количества.
	warningWrongCountValue := `wrong count value`                               // Сообщение о неверном количестве.
	warningCityNotSpecified := `city not specified`                             // Сообщение о том, что город не указан.
	warningWrongCityValue := `wrong city value`                                 // Сообщение о неверном городе.
	assert.Equal(t, expected, responseRecorder.Code)                            // Проверяем, что статус 200 корректный.
	assert.NotEmpty(t, responseRecorder.Body.String())                          // Проверяем, что тело ответа не пустое.
	assert.NotEqual(t, warningMissingCount, responseRecorder.Body.String())     // Проверка правильности запроса по ответу.
	assert.NotEqual(t, warningWrongCountValue, responseRecorder.Body.String())  // Проверка правильности запроса по ответу.
	assert.NotEqual(t, warningCityNotSpecified, responseRecorder.Body.String()) // Проверка правильности запроса по ответу.
	assert.NotEqual(t, warningWrongCityValue, responseRecorder.Body.String())   // Проверка правильности запроса по ответу.
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
	totalList := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}
	req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()
	list := strings.Split(body, ", ")

	expected := http.StatusOK                                     // Ожидаемое сообщение о статусе.
	expectedList := strings.Join(totalList[:totalCount], ", ")    // Ожидаемый ответ.
	assert.Equal(t, expected, responseRecorder.Code)              // Проверяем, что статус 200 корректный.
	assert.NotEmpty(t, responseRecorder.Body.String())            // Проверяем, что тело ответа не пустое.
	assert.Equal(t, expectedList, responseRecorder.Body.String()) // Проверка правильности ответа.
	assert.Equal(t, totalCount, len(list))                        // Проверка правильной обработки больших count.

}
