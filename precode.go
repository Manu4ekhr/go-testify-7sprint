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
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmpty(t, responseRecorder.Code)

	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)

	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)

}

func TestMainHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmpty(t, responseRecorder.Code)

	assert.NotNil(t, responseRecorder.Body)

	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=chelyabinsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusBadRequest := http.StatusBadRequest
	require.Equal(t, statusBadRequest, status)

	body := responseRecorder.Body.String()
	wrongCity := "wrong city value"
	assert.Equal(t, wrongCity, body)
}
