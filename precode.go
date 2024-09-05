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

func TestMainHandlerWhenAllOk(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // здесь нужно добавить необходимые проверки
	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)

}

func TestMainHandlerWhenCityWrong(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=10&city=sochi", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // здесь нужно добавить необходимые проверки
	status := responseRecorder.Code
	require.Equal(t, http.StatusBadRequest, status)
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
	assert.Equal(t, "wrong city value", body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // здесь нужно добавить необходимые проверки
	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
	bodyList := strings.Split(body, ",")
	assert.Len(t, bodyList, totalCount)
}
