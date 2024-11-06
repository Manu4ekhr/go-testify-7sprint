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

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4 // Указываем, что всего доступно 4 кафе

	// Создаем запрос с count больше, чем totalCount
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err)

	// Ответ на запрос
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 200 (успешный запрос)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body.String())

	// Проверяем, что количество кафе в ответе не превышает totalCount
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)

	// Проверяем, что в ответе возвращены все доступные кафе, так как count больше, чем доступные кафе
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=spb", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 400
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем, что сообщение об ошибке корректное
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMissing(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 400
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем, что сообщение об ошибке корректное
	assert.Equal(t, "count missing", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=invalid&city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что код ответа 400
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Проверяем, что сообщение об ошибке корректное
	assert.Equal(t, "wrong count value", responseRecorder.Body.String())
}
