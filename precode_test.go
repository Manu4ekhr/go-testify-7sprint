package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

/*
Обработчик возвращает список кафе.
В запросе указано сколько вернуть кафе и из какого города.
Если какие-то параметры указаны некорректно (нет такого города, неправильно указано количество), обработчик вернёт ошибку.
Сервер будет ожидать обращение по пути /cafe.

В GET-параметрах ожидается:

	count — количество кафе, которые нужно вернуть
	city — город, в котором нужно найти кафе

Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
В сервисе будет только один город moscow, в котором будет всего 4 кафе.
*/
func mainHandle(w http.ResponseWriter, req *http.Request) {

	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

// Проверки должны осуществляться с помощью пакета testify.
//
// Нужно реализовать три теста:
// ...
// ==========================================================================================
// 1. Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
// ==========================================================================================
func TestMainHandlerStatusOkAndNotEmptyBody(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// ...
	// Проверка статуса ответа
	require.Equal(t, responseRecorder.Code, http.StatusOK, "Принятый статус: %d не соответствует %d. Тест прерван!", responseRecorder.Code, http.StatusOK)

	// Проверка содержимого тела
	rxBody := responseRecorder.Body.String()

	rxList := strings.Split(rxBody, ",")

	assert.NotEmpty(t, len(rxList), "Ошибка. Тело ответа пустое. Было запрошено: 4 позиции!")

}

// ===========================================================================================================================================
// 2. Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
// ===========================================================================================================================================
func TestMainHandlerNotExistCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=omsk", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// ...
	// Проверка статуса ответа
	require.Equal(t, responseRecorder.Code, http.StatusOK, "Принятый статус: %d не соответствует %d. Тест прерван!. Тело ответа: %s", responseRecorder.Code, http.StatusOK, responseRecorder.Body.String())

}

// ==============================================================================================
// 3. Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
// ==============================================================================================
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=20&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// ...
	// Проверка статуса ответа
	require.Equal(t, responseRecorder.Code, http.StatusOK, "Принятый статус: %d не соответствует %d. Тест прерван!", responseRecorder.Code, http.StatusOK)

	// Проверка содержимого тела
	rxBody := responseRecorder.Body.String()

	rxList := strings.Split(rxBody, ",")

	assert.NotEqual(t, len(rxList), totalCount, "Нет соответствия. Запрошено: 20 позиций, а получено: %d позиций", len(rxList))

}
