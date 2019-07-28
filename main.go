package main

import (
	"data-sfu/src/gateway"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.POST("/join", gateway.CreateAnswerGateWay)
	e.POST("/ping", func(c echo.Context) error {
		type Ping struct {
			Msg string `json:"msg"`
		}
		res := &Ping{Msg: "pong"}
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":8080"))

}
