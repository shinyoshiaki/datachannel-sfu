package peer

import (
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Listen(peer *webrtc.PeerConnection, sdpReady chan Sdp) {
	sdp := <-sdpReady
	fmt.Println("onsdp", sdp.Type)
	if err := SetSDP(&sdp, peer); err != nil {
		fmt.Println("SetSDP", err)
	}
}
