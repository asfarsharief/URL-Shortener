package handler

import (
	"net/http"

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
	originalUrl := c.QueryParam("url")
	if originalUrl == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No URL sent"})
	}
	shortnedURL, err := us.urlShortnerService.GetShortUrl(originalUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"shortened_url": shortnedURL})
}
