package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
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

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	// Создаем новый HTTP-запрос с count больше, чем общее количество кафе
	req, err := http.NewRequest("GET", "/?count=10&city=moscow", nil)
	assert.NoError(t, err)

	// Создаем ResponseRecorder для записи ответа сервера
	responseRecorder := httptest.NewRecorder()

	// Вызываем обработчик хендлера с созданным запросом и записываем ответ в ResponseRecorder
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Проверяем, что тело ответа содержит все доступные кафе
	expectedAnswer := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedAnswer, responseRecorder.Body.String())

	// Проверяем, что count в ответе равен общему количеству кафе
	actualCount, err := strconv.Atoi(responseRecorder.Header().Get("Count"))
	assert.NoError(t, err)
	assert.Equal(t, totalCount, actualCount)
}
