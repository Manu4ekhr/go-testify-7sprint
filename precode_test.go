package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.NotEmpty(t, rr.Body.String())

	cafes := strings.Split(rr.Body.String(), ",")
	require.Len(t, cafes, 2)
	assert.Contains(t, cafes, "Мир кофе")
	assert.Contains(t, cafes, "Сладкоежка")
}

func TestMainHandlerInvalidCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=spb", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, "wrong city value", rr.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	cafes := strings.Split(rr.Body.String(), ",")
	require.Len(t, cafes, len(cafeList["moscow"]))

	for _, cafe := range cafeList["moscow"] {
		assert.Contains(t, cafes, cafe)
	}
}
