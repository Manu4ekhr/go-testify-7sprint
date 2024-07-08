package main

import(
	"net/http"
	"strings"
	"net/http/httptest"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)


func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
		totalCount := 4
		req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)
	
		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(responseRecorder, req)
	
	    require.Equal(t, http.StatusOK, responseRecorder.Code)
	    require.NotEmpty(t, responseRecorder.Body)

	    body := responseRecorder.Body.String()
		arr := strings.Split(body, ",")

		assert.Equal(t, totalCount, len(arr))
}

func TestMainHadlerWhenOk(t *testing.T){
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotNil(t, responseRecorder.Body)

}


func TestMainHandlerNotOk(t *testing.T){
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	expected := "wrong city value"
	require.Equal(t, expected, responseRecorder.Body.String())
}
