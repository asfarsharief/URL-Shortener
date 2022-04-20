package storage_test

import (
	"testing"

	"github.com/asfarsharief/URL-Shortener/storage"
	"github.com/stretchr/testify/assert"
)

func TestProcessUrl(t *testing.T) {
	storageHandler := storage.NewStorageHandler()
	url := "google.com/search"
	responseUrl := storageHandler.ProcessUrl(url)

	assert.Equal(t, "www.shorturl.com/1", responseUrl)
}

func TestProcessUrlReturnDifferentResponse(t *testing.T) {
	storageHandler := storage.NewStorageHandler()
	url := "google.com/search"
	responseUrl := storageHandler.ProcessUrl(url)
	assert.Equal(t, "www.shorturl.com/1", responseUrl)

	url = "google.com/find"
	responseUrl = storageHandler.ProcessUrl(url)
	assert.Equal(t, "www.shorturl.com/2", responseUrl)
}

func TestProcessUrlReturnSameUrl(t *testing.T) {
	storageHandler := storage.NewStorageHandler()
	url := "google.com/search"
	responseUrl := storageHandler.ProcessUrl(url)
	assert.Equal(t, "www.shorturl.com/1", responseUrl)

	responseUrl = storageHandler.ProcessUrl(url)
	assert.Equal(t, "www.shorturl.com/1", responseUrl)
}
