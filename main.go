package main

import (
	"net/http"
	"strconv"
	"strings"
)

// Прошу пояснить. У меня линтер в VScode говорит, что:
// redundant type from array, slice, or map composite literal
// Предлагает удалить []string перед списком названий кафе. Как быть?
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

	// Добавил проверку на наличие параметра city.
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("city not specified"))
		return
	}

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ", ") // Поправил разделение, чтобы после запятой был пробел.

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc(`/cafe`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
