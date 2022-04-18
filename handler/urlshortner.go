package handler

import (
	"github.com/asfarsharief/URL-Shortener/lib"
	"github.com/labstack/echo/v4"
)

type urlShortnerHandler struct {
	urlShortnerService lib.UrlShortnerInterface
}

func NewUrlShortnerHandler(urlShortnerService lib.UrlShortnerInterface) urlShortnerHandler {
	return urlShortnerHandler{
		urlShortnerService: urlShortnerService,
	}
}

func (us *urlShortnerHandler) UrlShortner(c echo.Context) error {
	originalUrl := c.Param("url")

	us.urlShortnerService.GetShortUrl(originalUrl)
	return nil
}
