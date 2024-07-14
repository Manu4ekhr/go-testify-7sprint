package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	count := 3
	city := "moscow"

	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", count, city), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerIncorrectCity(t *testing.T) {
	count := 3
	city := "snezhinsk"

	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", count, city), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, 400, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerLargeNumber(t *testing.T) {
	count := 30
	city := "moscow"

	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", count, city), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, len(strings.Split(responseRecorder.Body.String(), ",")), len(cafeList[city]))
}

func TestMainHandlerAllInOne(t *testing.T) {
	type args struct {
		count int
		city  string
	}

	type want struct {
		status int
		body   string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"correct",
			args{3, "moscow"},
			want{200, "[not empty]"},
		},
		{
			"incorrect city",
			args{3, "Snezhinsk"},
			want{400, "wrong city value"},
		},
		{
			"big number",
			args{30, "moscow"},
			want{200, "[not empty]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", tt.args.count, tt.args.city), nil)

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(mainHandle)
			handler.ServeHTTP(responseRecorder, req)

			assert.Equal(t, tt.want.status, responseRecorder.Code)
			if len(tt.want.body) > 0 {
				require.NotEmpty(t, string(responseRecorder.Body.String()))
			}
			if responseRecorder.Code == 400 {
				assert.Equal(t, tt.want.body, responseRecorder.Body.String())
			}
		})
	}
}
