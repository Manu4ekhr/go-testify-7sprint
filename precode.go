package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

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
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
	require.NoError(t, err, "Ошибка при создании запроса")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Код ответа  200")
	require.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа не должно быть пустым")

	responseBody := responseRecorder.Body.String()
	cafes := strings.Split(responseBody, ",")

	require.Len(t, cafes, 4, "Количество возвращенных кафе должно быть равно 4")
}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=3&city=unsupported", nil)
	require.NoError(t, err, "Ошибка при создании запроса")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Код ответа 400")
	require.Contains(t, responseRecorder.Body.String(), "wrong city value", "Тело ответа содержать ошибку wrong city value!!!!")
}

func TestMainHandlerWhenCountExceedsTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err, "Ошибка при создании запроса")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Код ответа  200")
	require.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа не должно быть пустым")

	responseBody := responseRecorder.Body.String()
	cafes := strings.Split(responseBody, ",")

	require.Len(t, cafes, totalCount, "Количество возвращенных кафе должно равно общему количеству доступных кафе")
}
