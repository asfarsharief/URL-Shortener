package lib

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/asfarsharief/URL-Shortener/constants"
	"github.com/asfarsharief/URL-Shortener/storage"
)

type UrlShortnerInterface interface {
	GetShortUrl(url string) (string, error)
}
type urlShortnerService struct {
	storageHandler storage.StorageHandlerInterface
}

func NewUrlShortnerService(storageHandler storage.StorageHandlerInterface) urlShortnerService {
	return urlShortnerService{
		storageHandler: storageHandler,
	}
}

func (us *urlShortnerService) GetShortUrl(url string) (string, error) {
	fmt.Println("Given URL : ", url)

	if ok, _ := regexp.MatchString(constants.UrlValidatorRegex, url); !ok {
		str := fmt.Sprintf("Invalid URL. Please validate. - %s", url)
		fmt.Println(str)
		return "", errors.New(str)
	}

	return us.storageHandler.ProcessUrl(url)
}
