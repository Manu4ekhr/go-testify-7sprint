package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	totalCount := len(cafeList[city])
	moreThenTotal := totalCount + 1
	url := fmt.Sprintf("/cafe?count=%d&city=%s", moreThenTotal, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusCode := responseRecorder.Code
	countedRes := len(strings.Split(responseRecorder.Body.String(), ","))

	require.Equal(t, statusCode, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, countedRes, totalCount)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	city := "TLT"
	totalCount := len(cafeList[city])
	url := fmt.Sprintf("/cafe?count=%d&city=%s", totalCount, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	data := "wrong city value"
	statusCode := responseRecorder.Code
	body := responseRecorder.Body.String()

	require.Equal(t, statusCode, http.StatusBadRequest)
	assert.Equal(t, body, data)
}
func TestMainHandlerWhenRightTotal(t *testing.T) {
	city := "moscow"
	totalCount := len(cafeList[city])
	url := fmt.Sprintf("/cafe?count=%d&city=%s", totalCount, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusCode := responseRecorder.Code
	countedRes := len(strings.Split(responseRecorder.Body.String(), ","))

	require.Equal(t, statusCode, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, countedRes, totalCount)
}
