package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	actualCity := req.URL.Query().Get("city")
	for city, _ := range cafeList {
		assert.Equal(t, city, actualCity)
	}
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.NotEmpty(t, body)
	assert.Len(t, list, totalCount)
	// здесь нужно добавить необходимые проверки
}
