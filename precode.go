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

    // здесь нужно добавить необходимые проверки

    require.Equal(t, responseRecorder.Code, http.StatusOK)

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    require.NotEmpty(t, list)
    assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenOk(t *testing.T) {
    expectedLen := 2
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusOK)

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    require.NotEmpty(t, list)
    assert.Len(t, list, expectedLen)

}

func TestMainHandlerWhenWrongCity(t *testing.T) {
    expectedError := "wrong city value"
    req := httptest.NewRequest("GEt", "/cafe?count=1&city=novgorod", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)

    body := responseRecorder.Body.String()
    assert.Equal(t, body, expectedError)

}