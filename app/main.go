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

const defaultDomain = "simple-domain"

func runCadenceWorkers(adapter *cadenceadapter.CadenceAdapter) {
	logrus.Info("runCadenceWorkers starting...")
	defer logrus.Info("runCadenceWorkers exited")

	// Configure worker params
	var (
		workflowTaskListName = workflows.TaskListName
		options              = worker.Options{}
	)

	// Start worker
	w := worker.New(adapter.ServiceClient, defaultDomain, workflowTaskListName, options)
	if err := w.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start workers")
	}
}

func runHTTPServer(adapter *cadenceadapter.CadenceAdapter) {
	logrus.Info("runHTTPServer starting...")
	defer logrus.Info("runHTTPServer exited")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", homeEndpoint)
	e.POST("/start", startWorkflowEndpoint)
	e.POST("/continue", continueWorkflowEndpoint)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}

func homeEndpoint(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func startWorkflowEndpoint(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func continueWorkflowEndpoint(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func main() {
	var adapter cadenceadapter.CadenceAdapter
	adapter.Setup(cadenceadapter.CadenceConfig{
		Domain:   defaultDomain,
		Service:  "cadence-frontend",
		HostPort: "127.0.0.1:7933",
	})

	runCadenceWorkers(&adapter)
	runHTTPServer(&adapter)
}
