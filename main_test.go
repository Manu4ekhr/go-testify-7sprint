package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Функция для выполнения запроса и возврата рекордера ответа
func retRequest(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)
	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(respRec, req)
	return respRec
}
func TestMainHSuccessTests(t *testing.T) {
	expectedCount := 4
	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"

	respRec := retRequest("/cafe?count=4&city=moscow")
	require.Equal(t, http.StatusOK, respRec.Code)

	body := respRec.Body.String()
	require.NotEmpty(t, body)

	list := strings.Split(body, ",")
	assert.Len(t, list, expectedCount)
	assert.Equal(t, expectedBody, body)
}
func TestMainHandlerMissingCount(t *testing.T) {
	respRec := retRequest("/cafe?city=moscow")
	require.Equal(t, http.StatusBadRequest, respRec.Code)

	body := respRec.Body.String()
	assert.Equal(t, "count missing", body)
}
func TestMainHandlerWrongCountValue(t *testing.T) {
	respRec := retRequest("/cafe?count=abc&city=moscow")
	require.Equal(t, http.StatusBadRequest, respRec.Code)

	body := respRec.Body.String()
	assert.Equal(t, "wrong count value", body)
}
func TestMainHandlerWrongCityValue(t *testing.T) {
	respRec := retRequest("/cafe?count=4&city=unknown")
	require.Equal(t, http.StatusBadRequest, respRec.Code)

	body := respRec.Body.String()
	assert.Equal(t, "wrong city value", body)
}
