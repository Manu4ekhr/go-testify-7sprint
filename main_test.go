package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	prefix     = "/cafes?"
	paramCity  = "city="
	paramCount = "&count="
)

func makeRequestParams(city string, count int) string {
	return fmt.Sprintf("%s%s%s%s%d", prefix, paramCity, city, paramCount, count)
}

// единая функция
func makeRequest(t *testing.T, paramAll string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", paramAll, nil)
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return nil, err
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder, nil
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	count := 10
	paramAll := makeRequestParams(city, count)

	responseRecorder, err := makeRequest(t, paramAll)
	if err != nil {
		return
	}

	body := responseRecorder.Body.String()

	// проверяем, что в cafeList есть ключ "moscow" и получаем список кафе из этого ключа
	cafes, ok := cafeList["moscow"]
	assert.True(t, ok)

	// создаем expectedCafes как срез из первых totalCount элементов списка кафе
	expectedCafes := cafes[:count]
	// создаем expectedBody как строку, полученную объединением expectedCafes через запятую.
	expectedBody := strings.Join(expectedCafes, ",")

	actual := strings.Split(body, ",")
	// сравниваем фактический список кафе, полученный из тела ответа, с expectedCafes
	assert.ElementsMatch(t, actual, expectedCafes)
	// сравниваем фактическое тело ответа с expectedBody.
	assert.Equal(t, expectedBody, body)

}

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func Test200AndBodyNotEmpty(t *testing.T) {
	city := "moscow"
	count := 10
	paramAll := makeRequestParams(city, count)

	responseRecorder, err := makeRequest(t, paramAll)
	if err != nil {
		return
	}

	// ответ возвращает код 200 (ОК).
	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	// тело ответа не пустое.
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, err)
}

// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestNoCity(t *testing.T) {
	city := "orel"
	count := 10
	paramAll := makeRequestParams(city, count)

	responseRecorder, err := makeRequest(t, paramAll)
	if err != nil {
		return
	}

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}
