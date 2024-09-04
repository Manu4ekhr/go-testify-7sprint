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

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code: %d, got %d", http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	cafeList := strings.Split(body, ",")
	assert.Len(t, cafeList, totalCount, "Response should contain %d cafes when count is %d", totalCount, len(cafeList))

	expectedCafes := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}
	assert.Equal(t, expectedCafes, cafeList, "Expected list of cafes did not match")

}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=3", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status: %d", http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body.String(), "Response body should not be empty")

}

func TestMainHandlerWhenUnknownCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=kazan&count=3", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status: %d", http.StatusBadRequest)
	assert.Equal(t, "wrong city value", responseRecorder.Code, "wrong city value")

}
