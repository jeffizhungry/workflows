package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	runHTTPServer()
}

func runHTTPServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", home)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
