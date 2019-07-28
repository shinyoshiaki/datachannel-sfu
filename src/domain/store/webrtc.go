package store

import (
	"github.com/google/uuid"
	"github.com/pion/webrtc/v2"
)

var peers = make(map[string]map[string]*webrtc.PeerConnection)

func SetPeer(peer *webrtc.PeerConnection, room string) (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uu := u.String()

	_, exist := peers[room]
	if exist == false {
		peers[room] = make(map[string]*webrtc.PeerConnection)
	}

	peers[room][uu] = peer

	return uu, nil
}

func DeletePeer(room string, uu string) {
	peers[room][uu] = nil
}

func GetPeers(room string) map[string]*webrtc.PeerConnection {
	groupe := peers[room]
	return groupe
}

func GetPeer(room string, uu string) *webrtc.PeerConnection {
	peer := peers[room][uu]
	return peer
}

var datachannels = make(map[string]map[string]*webrtc.DataChannel)

func SetDatachannel(dc *webrtc.DataChannel, room string, uu string) {
	_, exist := datachannels[room]
	if exist == false {
		datachannels[room] = make(map[string]*webrtc.DataChannel)
	}

	datachannels[room][uu] = dc
}

func GetDatachannels(room string) map[string]*webrtc.DataChannel {
	groupe := datachannels[room]
	return groupe
}

func DeleteDatachannel(room string, uu string) {
	datachannels[room][uu] = nil
}
