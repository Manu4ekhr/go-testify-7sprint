// precode.go
package main

import (
	"net/http"
	"strconv"
	"strings"
)

// Список кафе
var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

// Обработчик запросов
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

	// Проверка города
	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	// Если count больше, чем доступно, возвращаем все кафе
	if count > len(cafe) {
		count = len(cafe)
	}

	// Формируем ответ
	answer := strings.Join(cafe[:count], ",")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}
