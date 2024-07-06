package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	prefix     = "/cafes?"
	paramCity  = "city="
	paramCount = "&count="
)

func makeRequestParams(city string, count int) string {
	return fmt.Sprintf("%s%s%s%s%d", prefix, paramCity, city, paramCount, count)
}

// единая функция
func makeRequest(t *testing.T, paramAll string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", paramAll, nil)
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return nil, err
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder, nil
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	count := 10
	paramAll := makeRequestParams(city, count)

	responseRecorder, err := makeRequest(t, paramAll)
	if err != nil {
		return
	}

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что в карте в принципе есть ключ "moscow" (ведь мы его явно задаем).
	cafes, ok := cafeList["moscow"]
	assert.True(t, ok)
	// Проверяем, что длина списка кафе равна нашему count.
	assert.Len(t, cafes, count)
}

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func Test200AndBodyNotEmpty(t *testing.T) {
	city := "moscow"
	count := 10
	paramAll := makeRequestParams(city, count)

	responseRecorder, err := makeRequest(t, paramAll)
	if err != nil {
		return
	}

	// ответ возвращает код 200 (ОК).
	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	// тело ответа не пустое.
	body := responseRecorder.Body.String()
	// верно, раз тело ответа должно быть непустым, имеет смысл сразу останавливать выполнение теста, если это условие не выполняется. Ведь дальнейшие проверки в этом случае будут бессмысленны.
	require.NotEmpty(t, body, err)
}

// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestNoCity(t *testing.T) {
	city := "orel"
	count := 10
	paramAll := makeRequestParams(city, count)

	responseRecorder, err := makeRequest(t, paramAll)
	if err != nil {
		return
	}

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	require.Equal(t, responseRecorder.Body.String(), "wrong city value")
}
