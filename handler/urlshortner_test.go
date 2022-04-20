package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asfarsharief/URL-Shortener/handler"
	"github.com/asfarsharief/URL-Shortener/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUrlShortner(t *testing.T) {
	service := new(mocks.UrlShortnerInterface)
	url := "www.google.com/search"
	expectedUrl := "www.shorturl.com/1"
	service.Mock.On("GetShortUrl", url).Return(expectedUrl, nil)
	handler := handler.NewUrlShortnerHandler(service)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/shorten?url=www.google.com/search", nil)
	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)
	handler.UrlShortner(c)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	responsestr := `{"shortened_url": "www.shorturl.com/1"}`
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(responsestr), &jsonMap)
	if err != nil {
		panic(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, jsonMap, response)
}

func TestUrlShortnerErrorNoUrl(t *testing.T) {
	service := new(mocks.UrlShortnerInterface)
	// url := "www.google.com/search"
	// expectedUrl := "www.shorturl.com/1"
	// service.Mock.On("GetShortUrl", url).Return(expectedUrl, nil)
	handler := handler.NewUrlShortnerHandler(service)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/shorten?url=", nil)
	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)
	handler.UrlShortner(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	responsestr := `{"error": "No URL sent"}`
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(responsestr), &jsonMap)
	if err != nil {
		panic(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, jsonMap, response)
}

func TestUrlShortnerError(t *testing.T) {
	service := new(mocks.UrlShortnerInterface)
	url := "www.google.com/search"

	service.Mock.On("GetShortUrl", url).Return("", errors.New("service error"))
	handler := handler.NewUrlShortnerHandler(service)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/shorten?url=www.google.com/search", nil)
	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)
	handler.UrlShortner(c)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	responsestr := `{"error": "service error"}`
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(responsestr), &jsonMap)
	if err != nil {
		panic(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, jsonMap, response)
}
