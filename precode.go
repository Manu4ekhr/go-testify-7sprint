package main

import (
	"fmt"
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

func TestMainHandlerWhenCityWrongValue(t *testing.T) {
	city := "novosibirsk"

	cityRequest := fmt.Sprintf("/cafe?count=4&city=%s", city)
	req := httptest.NewRequest("GET", cityRequest, nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Result().StatusCode)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenRequestOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)
	assert.NotEmpty(t, responseRecorder.Body)

}

// Направляем два запроса к серверу для возможности более удобной настройки теста при изменении данных
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	count := 6

	countRequest := fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount)
	greaterCountRequest := fmt.Sprintf("/cafe?count=%d&city=moscow", count)

	req := httptest.NewRequest("GET", countRequest, nil)
	greaterReq := httptest.NewRequest("GET", greaterCountRequest, nil)

	responseRecorder := httptest.NewRecorder()
	greaterResponseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)

	handler.ServeHTTP(responseRecorder, req)
	countResponse := fmt.Sprintln(responseRecorder.Body)

	handler.ServeHTTP(greaterResponseRecorder, greaterReq)
	greaterCountResponse := fmt.Sprintln(greaterResponseRecorder.Body)

	assert.Equal(t, countResponse, greaterCountResponse)
}
