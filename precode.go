package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

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

func main() {
	http.HandleFunc("/cafe", mainHandle)
	http.ListenAndServe(":8080", nil)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// создаем запрос с count большим, чем количество кафе в списке
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	// проверяем, что ошибки нет
	assert.NoError(t, err)
	// создаем объект для записи ответа сервера
	responseRecorder := httptest.NewRecorder()
	// создаем обработчик запроса и вызываем его с запросом и объектом для записи ответа
	handler := http.HandlerFunc(mainHandle)
	// проверяем, что статус ответа 200 OK
	handler.ServeHTTP(responseRecorder, req)
	// проверяем, что статус ответа 200 OK
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	// проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())
	// проверяем, что в теле ответа содержатся все кафе из списка
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
}

func TestMainHandlerWithWrongCity(t *testing.T) {
	// создаем запрос с несуществующим городом
	req, err := http.NewRequest("GET", "/cafe?count=2&city=nonexistent", nil)
	// проверяем, что ошибки нет
	assert.NoError(t, err)
	// создаем объект для записи ответа сервера
	responseRecorder := httptest.NewRecorder()
	// создаем обработчик запроса и вызываем его с запросом и объектом для записи ответа
	handler := http.HandlerFunc(mainHandle)
	// проверяем, что статус ответа 400 Bad Request
	handler.ServeHTTP(responseRecorder, req)
	// проверяем, что статус ответа 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// проверяем, что тело ответа содержит сообщение об ошибке
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}

func TestMainHandlerWithMissingCount(t *testing.T) {
	// создаем запрос без параметра count
	req, err := http.NewRequest("GET", "/cafe?city=moscow", nil)
	// проверяем, что ошибки нет
	assert.NoError(t, err)
	// создаем объект для записи ответа сервера
	responseRecorder := httptest.NewRecorder()
	// создаем обработчик запроса и вызываем его с запросом и объектом для записи ответа
	handler := http.HandlerFunc(mainHandle)
	// проверяем, что статус ответа 400 Bad Request
	handler.ServeHTTP(responseRecorder, req)
	// проверяем, что статус ответа 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// проверяем, что тело ответа содержит сообщение об ошибке
	assert.Equal(t, "count missing", responseRecorder.Body.String())
}
