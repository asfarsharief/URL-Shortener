package main

import (
	"github.com/asfarsharief/URL-Shortener/handler"
	"github.com/asfarsharief/URL-Shortener/lib"
	"github.com/asfarsharief/URL-Shortener/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := createServer()
	server.Logger.Fatal(server.Start(":3030"))
}

func createServer() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//handler
	storageHandler := storage.NewStorageHandler()
	urlShortnerservice := lib.NewUrlShortnerService(&storageHandler)
	urlShortnerHandler := handler.NewUrlShortnerHandler(&urlShortnerservice)

	// Routes
	e.GET("/shorten", urlShortnerHandler.UrlShortner)

	// Start server
	return e
}
