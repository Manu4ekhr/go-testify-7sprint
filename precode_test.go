package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	//Я НЕ ПОНИМАЮ КАК И ЧТО МНЕ ЕЩЕ НАПИСАТЬ. В ПРЕКОДЕ - ВОТ ОНО ЕДИНТВЕННОЕ МЕСТО, КУДА ПИСАТЬ ПРОВЕРКИ
	//В УРОКЕ ПО TESTIFY ТОЛЬКО ТАК И ДЕЛАЛОСЬ, В ОДНУ ФУНКЦИЮ ПИХАЛИ НЕСКОЛЬКО ПРОВЕРОК
	//ПРИ НЕВЫПОЛНЕНИИ ОДНОЙ ПРОВЕРКИ ТЕСТИРОВЩИК ТАК И ПАДАЕТ, ВСЕ РАБОТАЕТ ТАК КАК И ДОЛЖНО

	//Проверка, пришел ли ответ
	require.NotEmpty(t, responseRecorder.Code)

	//Проверка, что статус 200
	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)

	//Проверка, что мы отправили нормальный город
	body := responseRecorder.Body.String()
	assert.NotEqual(t, "wrong city value", body)

	//Проверка, что в ответе все 4 кафе
	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)

}
