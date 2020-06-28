package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.uber.org/cadence/worker"
)

func main() {
	runCadenceWorkers()
	runHTTPServer()
}

func runCadenceWorkers() {
	// Configure worker params
	var (
		domain               = ""
		workflowTaskListName = ""
		options              = worker.Options{}
	)
	cadenceadapter{}

	// Start worker
	w := worker.New(nil, domain, workflowTaskListName, options)
	if err := w.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start workers")
	}
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
