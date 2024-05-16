package main

import (
	"fmt"
	"io/ioutil"
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

	//Город, передаваемый в city, не поддерживается
	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	//Если количество кафе более 4, вернуть все доступные кафе
	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

// Тестирование на корректность запроса
func TestMainHandleWhenRequestIsIncorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	_, error := ioutil.ReadAll(req.Body)
	if error != nil {
		status := http.StatusBadRequest
		fmt.Println(status)
	}
}

// Тест на наличие города в списке
func TestMainHandleWhenCityIsIncorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=city", nil)
	city := req.URL.Query().Get("city")
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	_, ok := cafeList[city]
	if !ok {
		status := http.StatusBadRequest
		assert.NotContains(t, city, "wrong city value", status)
	}
}

// Тест на указанное количество городов
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	if len(list) != totalCount {
		assert.Equal(t, totalCount, len(list))
	}
}
