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

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.NotEmpty(t, responseRecorder.Code)

	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	assert.NotEqual(t, "wrong city value", body)

	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)
}
