package lib

import "fmt"

type UrlShortnerInterface interface {
	GetShortUrl(url string)
}
type urlShortnerService struct {
}

func NewUrlShortnerService() urlShortnerService {
	return urlShortnerService{}
}

func (us *urlShortnerService) GetShortUrl(url string) {
	fmt.Println("URL ..: ", url)
}
