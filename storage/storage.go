package storage

import (
	"fmt"
	"math/rand"

	"github.com/asfarsharief/URL-Shortener/constants"
)

var urlMap = make(map[string]string)

type StorageHandlerInterface interface {
	ProcessUrl(url string) (string, error)
}
type storageHandler struct {
}

func NewStorageHandler() storageHandler {
	return storageHandler{}
}

func (sh *storageHandler) ProcessUrl(url string) (string, error) {
	if val, ok := urlMap[url]; ok {
		fmt.Println("URL already processed. Returning existing value")
		return fmt.Sprintf("%s/%s", constants.UrlBasePath, val), nil
	}
	uniquePath := make([]rune, 6)
	for i := range uniquePath {
		uniquePath[i] = constants.RandomLetterRune[rand.Intn(len(constants.RandomLetterRune))]
	}
	urlMap[url] = string(uniquePath)
	fmt.Println("URL successfully processed. ", uniquePath)
	return fmt.Sprintf("%s/%s", constants.UrlBasePath, string(uniquePath)), nil
}
