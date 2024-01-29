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

func getResponse(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenOk(t *testing.T) {
	responseRecorder := getResponse("/cafe?count=2&city=moscow")
	expectedCode := http.StatusOK
	expectedCount := 2

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	expectedBody := `wrong city value`
	responseRecorder := getResponse("/cafe?count=2&city=omsk")
	expectedCode := http.StatusBadRequest

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Equal(t, expectedBody, body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	responseRecorder := getResponse("/cafe?count=10&city=moscow")
	expectedCode := http.StatusOK
	expectedCount := 4

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}
