package testify

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerCorrectReq(t *testing.T) {
	fmt.Println("*** Correct request TEST ***")
	handler := http.HandlerFunc(mainHandle)
	for cafeCount := 1; cafeCount <= len(cafeList["moscow"]); cafeCount++ {
		// можно конечно и в рамках этого цикла проверить кейс с "count">кафе
		reqStr := fmt.Sprintf("/cafe?city=moscow&count=%d", cafeCount)
		req := httptest.NewRequest("GET", reqStr, nil)
		responseRecorder := httptest.NewRecorder()
		handler.ServeHTTP(responseRecorder, req)
		fmt.Printf("request string: '%s'\ncode: %d\nbody: %s\n", reqStr, responseRecorder.Code, responseRecorder.Body)
		assert.Equal(t, http.StatusOK, responseRecorder.Code, "Responce code for correct request is not 200 OK")
		assert.NotEmpty(t, responseRecorder.Body.String(), "Response body for correct request is empty")
	}
}

func TestMainHandlerCityMissing(t *testing.T) {
	fmt.Println("*** Incorrect city request TEST ***")
	handler := http.HandlerFunc(mainHandle)
	// сгенерируем случайные названия городов
	// набор букв и цифр, потому что в закрытых городах, где куют щит Родины, тоже пьют кофе
	charList := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxy" + "0123456789")
	for range 10 {
		cityNameLength := rand.IntN(20) + 1
		var cityBytes strings.Builder
		for l := 0; l < cityNameLength; l++ {
			cityBytes.WriteRune(charList[rand.IntN(len(charList))])
		}
		cityName := cityBytes.String()
		_, cityExist := cafeList[cityName] // проверяем, что случайно сгенерированный город отсуствует в базе
		if cityExist {
			continue
		}
		reqStr := fmt.Sprintf("/cafe?city=%s&count=%d", cityName, rand.IntN(30)+1)
		req := httptest.NewRequest("GET", reqStr, nil)
		responseRecorder := httptest.NewRecorder()
		handler.ServeHTTP(responseRecorder, req)
		fmt.Printf("request string: '%s'\ncode: %d\nbody: %s\n", reqStr, responseRecorder.Code, responseRecorder.Body)
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Responce code for missing city should be 400")
		assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Responce body for missing city should be 'wrong city value'")
	}
}

func TestMainHandlerCorrectCityWrongCount(t *testing.T) {
	fmt.Println("*** Correct request but count > real cafe number TEST ***")
	handler := http.HandlerFunc(mainHandle)
	for range 10 {
		cafeCount := len(cafeList["moscow"]) + rand.IntN(10)
		reqStr := fmt.Sprintf("/cafe?city=moscow&count=%d", cafeCount)
		req := httptest.NewRequest("GET", reqStr, nil)
		responseRecorder := httptest.NewRecorder()
		handler.ServeHTTP(responseRecorder, req)
		fmt.Printf("request string: '%s'\ncode: %d\nbody: %s\n", reqStr, responseRecorder.Code, responseRecorder.Body)
		assert.Equal(t, http.StatusOK, responseRecorder.Code, "Responce code for correct request is not 200 OK")
		assert.NotEmpty(t, responseRecorder.Body.String(), "Response body for correct request is empty")
	}
}
