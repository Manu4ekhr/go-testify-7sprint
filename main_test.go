package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) { //Если в параметре count указано больше,
	// чем есть всего, должны вернуться все доступные кафе.
	totalCount := 4
	// здесь нужно создать запрос к сервису

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	// проверка количества полученных в ответе кафе с их действительным количеством, при условии что count
	// больше этого количества
	body := responseRecorder.Body.String()

	cafeList := strings.Split(body, ",")

	assert.Len(t, cafeList, totalCount)

}

func TestMainHandlerWhenResponseBodyNotEmpty(t *testing.T) { //запрос сформирован корректно,
	// сервис возвращает код ответа 200 и тело ответа не пустое
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверка кода ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	//проверка не пустое ли тело ответа
	assert.NotEmpty(t, responseRecorder.Body)

}

func TestMainHandlerWhenResponseCityNotSupported(t *testing.T) { // Город, который передаётся в параметре city,
	// не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа

	req := httptest.NewRequest("GET", "/cafe?count=2&city=rostov", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	city := req.URL.Query().Get("city")

	assert.NotContains(t, cafeList, city)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}
