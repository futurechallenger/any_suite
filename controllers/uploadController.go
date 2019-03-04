package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// UploadController handle user's upload
type UploadController struct {
}

// UploadHandler will handle user's upload
func (uploader *UploadController) UploadHandler(c echo.Context) error {
	fmt.Println("Hello Uploader")
	return nil
}
