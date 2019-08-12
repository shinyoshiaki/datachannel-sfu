package peer

import (
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Listen(peer *webrtc.PeerConnection, sdpReady chan Sdp) {
	sdp := <-sdpReady
	fmt.Println("onsdp", sdp.Type)
	SetSDP(&sdp, peer)
}
