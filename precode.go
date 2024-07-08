package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
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

func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if !assert.Equal(t, http.StatusOK, responseRecorder.Code) {
		t.Fatalf("Expected %d, got %d", http.StatusOK, responseRecorder.Code)
	}
	if !assert.NotEmpty(t, responseRecorder.Body.String()) {
		t.Fatalf("response body is empty")
	}
}

// Другой город
func TestMainHandlerUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=saintpetersburg&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if !assert.Equal(t, http.StatusBadRequest, responseRecorder.Code) {
		t.Fatalf("Expected %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
	if !assert.Equal(t, "wrong city value", responseRecorder.Body.String()) {
		t.Fatalf("Waiting for: wrong city value, got %s", responseRecorder.Body.String())
	}
}

// Тест 3: Запрос с count больше, чем есть всего
func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	list := strings.Split(responseRecorder.Body.String(), ",")
	if !assert.Len(t, list, totalCount) {
		t.Fatalf("Waiting for all %d cafe's", totalCount)
	}
}
