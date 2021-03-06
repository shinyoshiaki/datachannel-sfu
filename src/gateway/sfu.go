package gateway

import (
	"data-sfu/src/controller"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Join(c echo.Context) error {
	var req controller.JoinReq
	c.Bind(&req)
	fmt.Println(req)

	res, err := controller.Join(req)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func Answer(c echo.Context) error {
	var req controller.AnswerReq
	c.Bind(&req)
	fmt.Println(req)

	err := controller.Answer(req)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
