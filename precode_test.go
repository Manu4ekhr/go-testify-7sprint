package main

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	src := rand.NewSource(time.Now().Unix())
	randomNumber := src.Int63()
	testCount := strconv.Itoa(totalCount + int(randomNumber%10) + 1)

	req := httptest.NewRequest("GET", "/cafe?count="+testCount+"&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	// Проверка на корректность запроса
	// Возвращает код 200 и не пустое тело ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, body)

	// Проверка на неподдерживаемый город в параметре city
	// Возвращает код 400 и wrong city value в теле ответа
	wrongCity := "russianGotham"
	wrongCityUrl := "/cafe?count=" + testCount + "&city=" + wrongCity
	errorMessage := "wrong city value"
	assert.HTTPStatusCode(t, handler, http.MethodGet, wrongCityUrl, nil, http.StatusBadRequest)
	assert.HTTPBodyContains(t, handler, http.MethodGet, wrongCityUrl, nil, errorMessage)

	// Проверка, когда count больше, чем общее число кафе
	// Возвращает все доступные кафе
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)
}
