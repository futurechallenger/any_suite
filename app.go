package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"int_ecosys/controllers"
	"int_ecosys/utils"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Renderer = utils.NewTemplateEngine()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public
	e.Static("/", "public")

	// Routes
	uploader := &controllers.UploadController{}
	home := &controllers.HomeController{}

	// test
	e.GET("/hello", home.HelloHandler)

	e.GET("/docker/install", home.HomeHandler)
	e.POST("/upload", uploader.UploadHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
