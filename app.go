package main

import (
	"flag"
	"fmt"
	"int_ecosys/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"int_ecosys/controllers"
	"int_ecosys/utils"
)

var BUILD_TYPE string

func main() {
	// Used in build, go build -ldflags "-X main.BUILD_TYPE=dev"
	fmt.Printf("BUILD TYPE: %s\n", BUILD_TYPE)

	// Used with run commana, go run app.go -env debug
	config.SetBuildEnv(flag.String("env", "debug", "runing in env `debug` or `release`"))

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
	e.GET("/home", home.HomeHandler)
	e.GET("/auth", home.AuthHandler)
	e.GET("/auth/callback", home.AuthCallbackHandler)

	e.GET("/docker/install", home.ContainerHandler)
	e.POST("/upload", uploader.UploadHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
