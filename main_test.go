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

func TestMainHandlerValidRequest(t *testing.T) {
	// 1
	//фейк запрос
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	//пишем ответ сюда
	responseRecorder := httptest.NewRecorder()
	//переводим хэндлер в объект
	handler := http.HandlerFunc(mainHandle)
	//симулируем обработчик
	handler.ServeHTTP(responseRecorder, req)
	//testify
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected 200 OK")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Expected non-empty body")
}

func TestMainHandlerInvalidCity(t *testing.T) {
	// 2
	req := httptest.NewRequest("GET", "/cafe?count=2&city=tokyo", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected 400 Bad Request")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Expected 'wrong city value' message")
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	// 3
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected 200 OK")
	expectedCafes := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedCafes, responseRecorder.Body.String(), "Expected all cafes")
}
