package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandler(t *testing.T) {
	tests := map[string]struct {
		reqMethod string
		reqCount  string
		reqCity   string
		expStatus int
		expCount  int
		expBody   string
	}{
		"when everything is correct": {
			reqMethod: http.MethodGet,
			reqCount:  "2",
			reqCity:   "moscow",
			expStatus: http.StatusOK,
			expCount:  2,
			expBody:   "Мир кофе,Сладкоежка",
		},
		"when the wrong city is specified": {
			reqMethod: http.MethodGet,
			reqCount:  "3",
			reqCity:   "sterlitamak",
			expStatus: http.StatusBadRequest,
			expBody:   "wrong city value",
		},
		"when count more than total": {
			reqMethod: http.MethodGet,
			reqCount:  "5",
			reqCity:   "moscow",
			expStatus: http.StatusOK,
			expCount:  4,
			expBody:   "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент",
		},
		"when count missing": {
			reqMethod: http.MethodPost,
			reqCount:  "",
			reqCity:   "moscow",
			expStatus: http.StatusBadRequest,
			expBody:   "count missing",
		},
		"when wrong count value": {
			reqMethod: http.MethodPost,
			reqCount:  "qq",
			reqCity:   "moscow",
			expStatus: http.StatusBadRequest,
			expBody:   "wrong count value",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(tt.reqMethod, "/cafe", nil)

			params := url.Values{}
			params.Add("city", tt.reqCity)
			params.Add("count", tt.reqCount)
			req.URL.RawQuery = params.Encode()

			responseRecorder := httptest.NewRecorder()

			handler := http.HandlerFunc(mainHandle)
			handler.ServeHTTP(responseRecorder, req)

			assert.Equal(t, tt.expStatus, responseRecorder.Code)

			responseBody := responseRecorder.Body.String()
			assert.NotEmpty(t, responseBody)
			assert.Equal(t, tt.expBody, responseBody)

			if responseRecorder.Code == http.StatusOK {
				list := strings.Split(responseBody, ",")
				assert.Equal(t, tt.expCount, len(list))
			}
		})
	}
}
