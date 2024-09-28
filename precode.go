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

func TestMainHandlerCorrectRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Неверный код ответа")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа пустое")
}

func TestMainHandlerWrongCity(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cafe?city=london&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Неверный код ответа")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Неверное сообщение об ошибке")
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Неверный код ответа")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа пустое")

	cafeNames := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, 4, len(cafeNames), "Неверное количество кафе в ответе")

	for _, cafeName := range cafeList["moscow"] {
		assert.Contains(t, cafeNames, cafeName, "Кафе %s не найдено в ответе", cafeName)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
