package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=15&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	arr := strings.Split(body, ",")

	totalCount := 4
	count := len(arr)

	assert.GreaterOrEqual(t, count, totalCount)
	assert.Equal(t, []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}, arr)
}
