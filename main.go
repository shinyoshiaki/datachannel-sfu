package main

import (
	"data-sfu/src/gateway"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.POST("/join", gateway.Join)
	e.POST("/answer", gateway.Answer)
	e.Logger.Fatal(e.Start(":8080"))

}
