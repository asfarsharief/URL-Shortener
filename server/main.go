package main

import (
	"github.com/asfarsharief/URL-Shortener/handler"
	"github.com/asfarsharief/URL-Shortener/lib"
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
	urlShortnerservice := lib.NewUrlShortnerService()
	urlShortnerHandler := handler.NewUrlShortnerHandler(&urlShortnerservice)

	// Routes
	e.GET("/url/:url", urlShortnerHandler.UrlShortner)

	// Start server
	return e
}
