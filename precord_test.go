package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestCafeAPISuite struct {
	suite.Suite
	testServer *httptest.Server
	client     *http.Client

	okRequest       *http.Request
	badCityRequest  *http.Request
	badCountRequest *http.Request
}

func (suite *TestCafeAPISuite) SetupSuite() {
	suite.testServer = httptest.NewServer(http.HandlerFunc(mainHandle))
	suite.client = suite.testServer.Client()

	var serverURL = suite.testServer.URL
	var unexistCity = "unexistCity"
	var existCity = "moscow"
	var existCityCafeCount = len(cafeList[existCity])

	suite.okRequest, _ = http.NewRequest(
		"GET", fmt.Sprintf("%s/cafe?city=%s&count=%d", serverURL, existCity, existCityCafeCount), nil,
	)

	suite.badCityRequest, _ = http.NewRequest(
		"GET", fmt.Sprintf("%s/cafe?city=%s&count=%d", serverURL, unexistCity, existCityCafeCount+1), nil,
	)

	suite.badCountRequest, _ = http.NewRequest(
		"GET", fmt.Sprintf("%s/cafe?city=moscow&count=%d", serverURL, existCityCafeCount+1), nil,
	)
}

func (suite *TestCafeAPISuite) TestMainHandlerWhenOk() {
	var test = suite.T()

	var response, _ = suite.client.Do(suite.okRequest)

	require.Equal(test, 200, response.StatusCode)

	defer response.Body.Close()

	var body, _ = io.ReadAll(response.Body)

	assert.NotEmpty(test, body)
}

func (suite *TestCafeAPISuite) TestMainHandlerWhenCityIsWrong() {
	var test = suite.T()

	var response, _ = suite.client.Do(suite.badCityRequest)

	assert.Equal(test, 400, response.StatusCode)

	defer response.Body.Close()

	var body, _ = io.ReadAll(response.Body)
	assert.Equal(test, "wrong city value", string(body))
}

func (suite *TestCafeAPISuite) TestMainHandlerWhenCountMoreThanTotal() {
	var test = suite.T()
	var cafeCount, _ = strconv.Atoi(
		suite.okRequest.URL.Query().Get("count"),
	)

	var response, _ = suite.client.Do(suite.badCountRequest)

	assert.Equal(test, 200, response.StatusCode)

	defer response.Body.Close()

	var body, _ = io.ReadAll(response.Body)
	var cafes = strings.Split(string(body), ",")

	assert.Len(test, cafes, cafeCount)
}

func TestCafeAPI(t *testing.T) {
	suite.Run(t, new(TestCafeAPISuite))
}
