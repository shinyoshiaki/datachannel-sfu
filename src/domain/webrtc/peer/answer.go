package peer

import (
	"fmt"

	"github.com/pion/webrtc/v2"
)

func createAnswerPeer() *webrtc.PeerConnection {
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

	return peerConnection
}

func SetOffer(sdp webrtc.SessionDescription) (webrtc.SessionDescription, *webrtc.PeerConnection) {
	peer := createAnswerPeer()
	err := peer.SetRemoteDescription(sdp)
	if err != nil {
		panic(err)
	}
	answer, err := peer.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("setoffer")
	return answer, peer
}
