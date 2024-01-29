package cafehandler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	city := "moscow"
	count := 2
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?city=%s&count=%d", city, count), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	res := responseRecorder.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	assert.NotEmpty(t, data)
}

func TestMainHandlerWhenCityNotSupport(t *testing.T) {
	city := "AgurdJ2Bq9S6"
	count := 1
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?city=%s&count=%d", city, count), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	res := responseRecorder.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Equal(t, []byte("wrong city value"), data)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	count := 5
	totalCount := len(cafeList[city]) // 4 это было в прекоде)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?city=%s&count=%d", city, count), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	res := responseRecorder.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	assert.Equal(t, totalCount, len(strings.Split(string(data), ",")))
}
