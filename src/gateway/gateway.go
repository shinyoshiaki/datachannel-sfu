package gateway

import (
	"data-sfu/src/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pion/webrtc/v2"
)

type CreateAnswerReq struct {
	TYPE webrtc.SDPType `json:"type"`
	SDP  string         `json:"sdp"`
}

type CreateAnswerRes struct {
	SDP webrtc.SessionDescription `json:"sdp"`
	UU  string                    `json:"uu"`
}

func CreateAnswerGateWay(c echo.Context) error {
	var req CreateAnswerReq
	c.Bind(&req)
	sdp, uu, err := usecase.Join(req.TYPE, req.SDP)
	if err != nil {
		fmt.Println("error", err)
		return c.String(http.StatusInternalServerError, "error")
	}
	res := &CreateAnswerRes{SDP: sdp, UU: uu}
	fmt.Println("res")
	return c.JSON(http.StatusOK, res)
}
