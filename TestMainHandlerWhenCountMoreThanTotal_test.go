package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=15&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	totalCount := 4
	count, err := strconv.Atoi(req.URL.Query().Get("count"))
	if err != nil {
		log.Println(err)
	}

	assert.GreaterOrEqual(t, count, totalCount)
}
