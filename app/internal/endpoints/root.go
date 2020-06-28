package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jeffizhungry/workflows/app/internal/cadenceadapter"
	"github.com/jeffizhungry/workflows/app/internal/workflows"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/cadence/client"
)

// RegisterRootHandler registers root route handlers
func RegisterRootHandler(e *echo.Echo, adapter *cadenceadapter.CadenceAdapter) {
	root := rootHandler{adapter}
	e.GET("/", root.home)
	e.GET("/start", root.startWorkflow)
	e.GET("/continue", root.continueWorkflow)
}

type rootHandler struct {
	adapter *cadenceadapter.CadenceAdapter
}

func (h *rootHandler) home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func (h *rootHandler) startWorkflow(c echo.Context) error {
	// Parse input
	accountID := c.QueryParam("accountId")
	if accountID == "" {
		return c.String(http.StatusBadRequest, "`accountId` query param is required")
	}

	// Run
	wo := client.StartWorkflowOptions{
		TaskList:                     workflows.TaskListName,
		ExecutionStartToCloseTimeout: time.Hour * 24,
	}
	execution, err := h.adapter.CadenceClient.StartWorkflow(context.Background(), wo, workflows.Workflow, accountID)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error starting workflow: %w", err))
	}
	logrus.WithFields(logrus.Fields{
		"workflowId": execution.ID,
		"runId":      execution.RunID,
	}).Info("Started workflow!")
	return c.JSONPretty(http.StatusOK, execution)
}

func (h *rootHandler) continueWorkflow(c echo.Context) error {
	// Parse input
	workflowID := c.QueryParam("workflowId")
	if workflowID == "" {
		return c.String(http.StatusBadRequest, "`workflowId` query param is required")
	}
	ageRaw := c.QueryParam("age")
	if ageRaw == "" {
		return c.String(http.StatusBadRequest, "`age` query param is required")
	}
	age, err := strconv.Atoi(ageRaw)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid `age` query param %v: %w", ageRaw, err))
	}

	// Run
	if err = h.adapter.CadenceClient.SignalWorkflow(context.Background(), workflowID, "", workflows.SignalName, age); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error signalling workflow: %w", err))
	}
	logrus.WithFields(logrus.Fields{
		"workflowId": workflowID,
		"age":        age,
	}).Info("Signalled workflow!")
	return c.String(http.StatusOK, "Success")
}
