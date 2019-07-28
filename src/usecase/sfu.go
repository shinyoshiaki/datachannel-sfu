package usecase

import (
	"data-sfu/src/domain/sfu"
	"data-sfu/src/domain/store"
	"data-sfu/src/domain/webrtc/peer"
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Join(room string) (webrtc.SessionDescription, string, error) {
	offer, peer, dc := peer.CreatePeer()

	uu, err := store.SetPeer(peer, room)
	if err != nil {
		fmt.Println("error", err)
		return webrtc.SessionDescription{}, "", err
	}

	store.SetDatachannel(dc, room, uu)

	peer.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		if connectionState.String() == "disconnected" {
			store.DeletePeer(room, uu)
			store.DeleteDatachannel(room, uu)
		}
	})

	sfu.Publish(dc, room, uu)

	return offer, uu, nil
}

func Answer(room string, uu string, TYPE webrtc.SDPType, SDP string) {
	sdp := webrtc.SessionDescription{Type: TYPE, SDP: SDP}
	peer := store.GetPeer(room, uu)
	err := peer.SetRemoteDescription(sdp)
	if err != nil {
		panic(err)
	}
}
