package usecase

import (
	"data-sfu/src/domain/store"
	"data-sfu/src/domain/webrtc/peer"
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Join(TYPE webrtc.SDPType, SDP string) (webrtc.SessionDescription, string, error) {
	sdp := webrtc.SessionDescription{}
	sdp.Type = TYPE
	sdp.SDP = SDP
	answer, peer := peer.SetOffer(sdp)
	uu, err := store.SetPeer(peer)
	if err != nil {
		fmt.Println("error", err)
		return webrtc.SessionDescription{}, "", err
	}
	fmt.Println("answer", uu)
	return answer, uu, nil
}
