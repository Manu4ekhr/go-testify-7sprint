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
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
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

// Тест на проверку работы хендлера при действительном запросе
func TestMainHandlerValidRequst(t *testing.T) {
	// Создаем запрос к сервису с параметрами
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	// создаем для записи ответа сервиса
	responseRecorder := httptest.NewRecorder()
	// создаем хендлер
	handler := http.HandlerFunc(mainHandle)
	// вызываем хендлер с запросом и записью ответа
	handler.ServeHTTP(responseRecorder, req)
	// проверяем ответ сервиса
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	// проверка содержимого ответа
	expectedBody := "Мир кофе,Сладкоежка"
	// проверяем, что тело ответа сервиса не пустое
	assert.NotEmpty(t, expectedBody, responseRecorder.Body.String())
}

// тест на проверку хандлера при недействительном городе
func TestMainHandlerInvalidCity(t *testing.T) {
	// Создаем запрос к сервису с параметрами
	req := httptest.NewRequest("GET", "/cafe?count=2&city=novgorod", nil)
	// создаем для записи ответа сервиса
	responseRecorder := httptest.NewRecorder()
	// создаем хендлер
	handler := http.HandlerFunc(mainHandle)
	// вызываем хендлер с запросом и записью ответа
	handler.ServeHTTP(responseRecorder, req)
	// проверяем статус ответа
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// проверяем тело ответа сервиса
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

// тест на проверку хандлера при недействительном количестве кафе
func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	// Создаем запрос к сервису с параметрами
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	// создаем для записи ответа сервиса
	responseRecorder := httptest.NewRecorder()
	// создаем хендлер
	handler := http.HandlerFunc(mainHandle)
	// вызываем хендлер с запросом и записью ответа
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки
	// проверяем ответ
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	// получаем ответ сервиса
	body := responseRecorder.Body.String()
	// разделение строки на слайс
	list := strings.Split(body, ",")
	// прверяем количество кафе
	assert.Len(t, list, totalCount)
}
