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

	// Public
	e.Static("/", "public")

	// Routes
	uploader := &controllers.UploadController{}
	// home := &controllers.HomeController{}

	// e.GET("/", home.HomeHandler)
	e.POST("/upload", uploader.UploadHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
