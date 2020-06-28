package endpoints

import (
	"net/http"

	"github.com/jeffizhungry/workflows/app/internal/cadenceadapter"
	"github.com/labstack/echo/v4"
)

// RegisterRootHandler registers root route handlers
func RegisterRootHandler(e *echo.Echo, adapter *cadenceadapter.CadenceAdapter) {
	root := rootHandler{adapter}
	e.GET("/", root.home)
	e.POST("/start", root.startWorkflow)
	e.POST("/continue", root.continueWorkflow)
}

type rootHandler struct {
	adapter *cadenceadapter.CadenceAdapter
}

func (h *rootHandler) home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func (h *rootHandler) startWorkflow(c echo.Context) error {
	// accountID := c.QueryParam("accountId")
	return c.String(http.StatusOK, "Hello, World")
}

func (h *rootHandler) continueWorkflow(c echo.Context) error {
	// workflowID := c.QueryParam("workflowId")
	// ageRaw := c.QueryParam("age")
	// age, err := strconv.Atoi(ageRaw)
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid `age` query param, got %v", ageRaw))
	// }
	return c.String(http.StatusOK, "Hello, World")
}
