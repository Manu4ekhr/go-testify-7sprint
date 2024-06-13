package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	// totalCount := 4

	type want struct {
		statusCode int
		answer     string
	}

	tests := []struct {
		name  string
		count string
		city  string
		want  want
	}{
		{
			name:  "simple test",
			count: "3",
			city:  "moscow",
			want: want{
				statusCode: 200,
				answer:     "Мир кофе,Сладкоежка,Кофе и завтраки",
			},
		},
		{
			name:  "city not exist",
			count: "10",
			city:  "Petegburg",
			want: want{
				statusCode: 400,
				answer: "wrong city value",
			},
		},
		{
			name:  "count more then exist",
			count: "10",
			city:  "moscow",
			want: want{
				statusCode: 200,
				answer:     "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент",
			},
		},
		{
			name:  "count missing",
			count: "",
			city:  "moscow",
			want: want{
				statusCode: 400,
				answer: "count missing",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			req:= httptest.NewRequest(http.MethodGet, "/", nil)
			
			q:=req.URL.Query()
			q.Add("count", tt.count)
			q.Add("city", tt.city)
			req.URL.RawQuery = q.Encode()

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(mainHandle)
			handler.ServeHTTP(responseRecorder, req)

			res := responseRecorder.Result()

			assert.Equal(t, tt.want.statusCode, res.StatusCode,tt.name) //проверяем статус запроса

			result, err := io.ReadAll(res.Body) //читаем тело ответа
			require.NoError(t, err)
			err = res.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, tt.want.answer, string(result),tt.name) //проверяем текст ответа
		})
	}
}

/*
{
    totalCount := 4
    req := ... // здесь нужно создать запрос к сервису

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // здесь нужно добавить необходимые проверки
}
*/
