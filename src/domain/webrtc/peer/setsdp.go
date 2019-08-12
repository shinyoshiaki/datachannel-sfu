package peer

import (
	"github.com/pion/webrtc/v2"
)

func SetSDP(sdp *Sdp, peer *webrtc.PeerConnection) error {
	switch sdp.Type {
	case "candidate":
		ice := webrtc.ICECandidateInit{Candidate: sdp.Sdp}
		err := peer.AddICECandidate(ice)
		if err != nil {
			return err
		}
	case "offer":
		sdp := webrtc.SessionDescription{Type: webrtc.SDPTypeOffer, SDP: sdp.Sdp}
		err := peer.SetRemoteDescription(sdp)
		if err != nil {
			return err
		}
	case "answer":
		sdp := webrtc.SessionDescription{Type: webrtc.SDPTypeAnswer, SDP: sdp.Sdp}
		err := peer.SetRemoteDescription(sdp)
		if err != nil {
			return err
		}
	}
	return nil
}
