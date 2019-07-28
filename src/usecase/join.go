package usecase

import (
	"data-sfu/src/domain/sfu"
	"data-sfu/src/domain/store"
	"data-sfu/src/domain/webrtc/peer"
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Join(TYPE webrtc.SDPType, SDP string, room string) (webrtc.SessionDescription, string, error) {
	sdp := webrtc.SessionDescription{Type: TYPE, SDP: SDP}
	answer, peer := peer.SetOffer(sdp)

	fmt.Println("room", room)

	uu, err := store.SetPeer(peer, room)

	if err != nil {
		fmt.Println("error", err)
		return webrtc.SessionDescription{}, "", err
	}

	sfu.Publish(peer, room, uu)

	return answer, uu, nil
}
