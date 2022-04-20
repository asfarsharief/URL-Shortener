package storage

import (
	"fmt"

	"github.com/asfarsharief/URL-Shortener/constants"
)

type StorageHandlerInterface interface {
	ProcessUrl(url string) string
}
type storageHandler struct {
	urlMap    map[string]int
	nextValue int
}

func NewStorageHandler() storageHandler {
	return storageHandler{
		urlMap:    make(map[string]int),
		nextValue: 1,
	}
}

func (sh *storageHandler) ProcessUrl(url string) string {
	if val, ok := sh.urlMap[url]; ok {
		fmt.Println("URL already processed. Returning existing value")
		return fmt.Sprintf("%s/%s", constants.UrlBasePath, encodeBase62(val))
	}

	sh.urlMap[url] = sh.nextValue
	sh.nextValue++

	return fmt.Sprintf("%s/%s", constants.UrlBasePath, encodeBase62(sh.urlMap[url]))
}

func encodeBase62(uniqueID int) string {
	uniqueRune := []rune{}
	for uniqueID > 0 {
		div := uniqueID % 62
		uniqueRune = append(uniqueRune, constants.RandomLetterRune[div])
		uniqueID = uniqueID / 62
	}
	for i, j := 0, len(uniqueRune)-1; i < j; i, j = i+1, j-1 {
		uniqueRune[i], uniqueRune[j] = uniqueRune[j], uniqueRune[i]
	}
	return string(uniqueRune)
}
