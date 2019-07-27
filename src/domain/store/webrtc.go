package store

import (
	"github.com/google/uuid"
	"github.com/pion/webrtc/v2"
)

var store = make(map[string]*webrtc.PeerConnection)

func SetPeer(peer *webrtc.PeerConnection) (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uu := u.String()
	store[uu] = peer

	return uu, nil
}
