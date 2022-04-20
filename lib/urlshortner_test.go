package lib_test

import (
	"testing"

	"github.com/asfarsharief/URL-Shortener/lib"
	"github.com/asfarsharief/URL-Shortener/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetShortUrl(t *testing.T) {
	handler := new(mocks.StorageHandlerInterface)

	url := "google.com/search"
	expectedUrl := "www.shorturl.com/1"
	handler.Mock.On("ProcessUrl", url).Return(expectedUrl)
	service := lib.NewUrlShortnerService(handler)
	responseUrl, err := service.GetShortUrl(url)

	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, responseUrl)
}

func TestGetShortUrlErrorInvalidUrl(t *testing.T) {
	handler := new(mocks.StorageHandlerInterface)

	url := "invalid url"
	service := lib.NewUrlShortnerService(handler)
	_, err := service.GetShortUrl(url)

	assert.Error(t, err)
	assert.Equal(t, "Invalid URL. Please validate. - invalid url", err.Error())
}

func TestGetShortUrlWithHttp(t *testing.T) {
	handler := new(mocks.StorageHandlerInterface)

	url := "http://google.com/search"
	expectedUrl := "www.shorturl.com/1"
	handler.Mock.On("ProcessUrl", "google.com/search").Return(expectedUrl)
	service := lib.NewUrlShortnerService(handler)
	responseUrl, err := service.GetShortUrl(url)

	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, responseUrl)
}

func TestGetShortUrlWithWWW(t *testing.T) {
	handler := new(mocks.StorageHandlerInterface)

	url := "www.google.com/search"
	expectedUrl := "www.shorturl.com/1"
	handler.Mock.On("ProcessUrl", "google.com/search").Return(expectedUrl)
	service := lib.NewUrlShortnerService(handler)
	responseUrl, err := service.GetShortUrl(url)

	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, responseUrl)
}
