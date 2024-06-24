package main

import (
	"net/http"
	"strconv"
	"strings"
)

// cafeList - это карта, которая хранит список кафе для каждого поддерживаемого города.
var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

// mainHandle - это основная функция-обработчик для конечной точки "/cafe".

func mainHandle(w http.ResponseWriter, req *http.Request) {
	// Получаем значение параметра "count" из URL запроса.
	countStr := req.URL.Query().Get("count")
	// Если параметр "count" отсутствует, возвращаем код состояния Bad Request и сообщение об ошибке.
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	// Преобразуем значение параметра "count" в целое число.
	count, err := strconv.Atoi(countStr)
	// Если значение параметра "count" не является допустимым целым числом, возвращаем код состояния Bad Request и сообщение об ошибке.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	// Получаем значение параметра "city" из URL запроса.
	city := req.URL.Query().Get("city")

	// Получаем список кафе для указанного города из карты cafeList.
	cafe, ok := cafeList[city]
	// Если указанный город не поддерживается, возвращаем код состояния Bad Request и сообщение об ошибке.
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	// Если запрошенное количество кафе больше доступного количества, устанавливаем count равным доступному количеству.
	if count > len(cafe) {
		count = len(cafe)
	}

	// Объединяем первые "count" кафе в списке в строку, разделенную запятыми.
	answer := strings.Join(cafe[:count], ",")

	// Записываем код состояния OK и список кафе в ответ.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

// Функция main настраивает конечную точку "/cafe" для использования функции mainHandle в качестве обработчика и запускает HTTP-сервер на порту 8080.
func main() {
	http.HandleFunc("/cafe", mainHandle)
	http.ListenAndServe(":8080", nil)
}
