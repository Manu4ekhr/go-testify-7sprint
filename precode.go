package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")
	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

// тест будет работать независимо от того, какие кафе и сколько их определено в cafeList.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req, err := http.NewRequest("GET", "/cafes?city=moscow&count=10", nil)
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
		return
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// ответ возвращает код 200 (ОК).
	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	// тело ответа не пустое.
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, err)

	// проверяем, что в cafeList есть ключ "moscow" и получаем список кафе из этого ключа
	cafes, ok := cafeList["moscow"]
	assert.True(t, ok)

	// создаем expectedCafes как срез из первых totalCount элементов списка кафе
	expectedCafes := cafes[:totalCount]
	// создаем expectedBody как строку, полученную объединением expectedCafes через запятую.
	expectedBody := strings.Join(expectedCafes, ",")

	actual := strings.Split(body, ",")
	// сравниваем фактический список кафе, полученный из тела ответа, с expectedCafes
	assert.ElementsMatch(t, actual, expectedCafes)
	// сравниваем фактическое тело ответа с expectedBody.
	assert.Equal(t, expectedBody, body)

}
