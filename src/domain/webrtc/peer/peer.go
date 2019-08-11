package peer

import (
	"github.com/pion/webrtc/v2"
)

func CreatePeer() (webrtc.SessionDescription, *webrtc.PeerConnection, *webrtc.DataChannel) {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	dc, err := peerConnection.CreateDataChannel("datachannel", nil)
	if err != nil {
		panic(err)
	}

	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	return offer, peerConnection, dc
}
