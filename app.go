package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"int_ecosys/controllers"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	uploader := &controllers.UploadController{}
	e.GET("/", uploader.UploadHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
