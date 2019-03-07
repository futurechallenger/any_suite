package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeController handle user's upload
type HomeController struct {
}

// HomeHandler will handle user's upload
func (home *HomeController) HomeHandler(c echo.Context) error {
	fmt.Println("Hello Uploader")
	return c.String(http.StatusOK, "Hello World!")
}
