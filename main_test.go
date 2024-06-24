package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body, err := io.ReadAll(responseRecorder.Body)

	require.NoError(t, err)
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	require.Greater(t, len(body), 0)
}

func TestMainHandlerWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=irkutsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	require.Equal(t, string(body), "wrong city value")
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)

	cafe := strings.Join([]string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}, ",")
	// здесь нужно добавить необходимые проверки
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	require.Equal(t, string(body), cafe)
}
