package controller

import (
	"data-sfu/src/usecase"

	"github.com/pion/webrtc/v2"
)

type JoinReq struct {
	ROOM string `json:"room"`
}

type JoinRes struct {
	SDP webrtc.SessionDescription `json:"sdp"`
	UU  string                    `json:"uu"`
}

func Join(req JoinReq) (*JoinRes, error) {
	sdp, uu, err := usecase.Join(req.ROOM)

	if err != nil {
		return nil, err
	}

	res := &JoinRes{SDP: sdp, UU: uu}
	return res, nil
}

type SignalingReq struct {
	ROOM string `json:"room"`
	UU   string `json:"uu"`
	TYPE string `json:"type"`
	SDP  string `json:"sdp"`
}

func Signaling(req SignalingReq) error {
	err := usecase.Signaling(req.ROOM, req.UU, req.TYPE, req.SDP)

	if err != nil {
		return err
	}

	return nil
}
