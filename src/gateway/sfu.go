package gateway

import (
	"data-sfu/src/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Join(c echo.Context) error {
	var req controller.JoinReq
	c.Bind(&req)

	res, err := controller.Join(req)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func Signaling(c echo.Context) error {
	var req controller.SignalingReq
	c.Bind(&req)

	err := controller.Signaling(req)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
