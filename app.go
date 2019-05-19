/**
TODO:
1. Uplaod file just accept uploaded files, then process files.
Run docker in the last
*/

package main

import (
	"any_suite/config"
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"any_suite/controllers"
	"any_suite/utils"
)

var BUILD_TYPE string

func main() {
	// Used in build, go build -ldflags "-X main.BUILD_TYPE=dev"
	fmt.Printf("BUILD TYPE: %s\n", BUILD_TYPE)

	// Used with run commana, go run app.go -env debug
	config.SetBuildEnv(flag.String("env", "debug", "runing in env `debug` or `release`"))

	// Start DB
	// data.NewIntEcoDB()

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
	e.POST("/upload/ret", uploader.UploadCompleteHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// TODO: Add upload indicator to handle upload success or failure
