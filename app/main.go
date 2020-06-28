package main

import (
	"net/http"

	"github.com/jeffizhungry/workflows/app/internal/cadenceadapter"
	"github.com/jeffizhungry/workflows/app/internal/workflows"
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
	logrus.Info("runCadenceWorkers starting...")
	defer logrus.Info("runCadenceWorkers exited")

	// Configure worker params
	var (
		domain               = "simple-domain"
		workflowTaskListName = workflows.TaskListName
		options              = worker.Options{}
	)

	var adapter cadenceadapter.CadenceAdapter
	adapter.Setup(cadenceadapter.CadenceConfig{
		Domain:   domain,
		Service:  "cadence-frontend",
		HostPort: "127.0.0.1:7933",
	})

	// Start worker
	w := worker.New(adapter.ServiceClient, domain, workflowTaskListName, options)
	if err := w.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start workers")
	}
}

func runHTTPServer() {
	logrus.Info("runHTTPServer starting...")
	defer logrus.Info("runHTTPServer exited")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", home)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
