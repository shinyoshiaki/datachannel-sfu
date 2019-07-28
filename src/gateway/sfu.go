package gateway

import (
	"data-sfu/src/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pion/webrtc/v2"
)

type JoinReq struct {
	ROOM string `json:"room"`
}

type JoinRes struct {
	SDP webrtc.SessionDescription `json:"sdp"`
	UU  string                    `json:"uu"`
}

func Join(c echo.Context) error {
	var req JoinReq
	c.Bind(&req)

	sdp, uu, err := usecase.Join(req.ROOM)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	res := &JoinRes{SDP: sdp, UU: uu}
	return c.JSON(http.StatusOK, res)
}

type AnswerReq struct {
	ROOM string         `json:"room"`
	UU   string         `json:"uu"`
	TYPE webrtc.SDPType `json:"type"`
	SDP  string         `json:"sdp"`
}

func Answer(c echo.Context) error {
	var req AnswerReq
	c.Bind(&req)

	usecase.Answer(req.ROOM, req.UU, req.TYPE, req.SDP)

	return c.NoContent(http.StatusOK)
}
