package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RunnerController runs something
type RunnerController struct {
}

// RunnerHandler handles run directive
func (runner *RunnerController) RunnerHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello, this is runner")
}
