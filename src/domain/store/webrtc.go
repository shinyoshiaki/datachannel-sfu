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
	delete(peers[room], uu)
	if len(peers[room]) == 0 {
		delete(peers, room)
	}

	deleteDatachannel(room, uu)
}

func GetPeers(room string) map[string]*webrtc.PeerConnection {
	groupe := peers[room]
	return groupe
}

func GetPeer(room string, uu string) *webrtc.PeerConnection {
	peer := peers[room][uu]
	return peer
}

var datachannels = make(map[string]map[string]map[string]*webrtc.DataChannel)

func SetDatachannel(dc *webrtc.DataChannel, room string, label string, uu string) {
	_, exist := datachannels[room]
	if exist == false {
		datachannels[room] = make(map[string]map[string]*webrtc.DataChannel)
	}
	_, exist = datachannels[room][label]
	if exist == false {
		datachannels[room][label] = make(map[string]*webrtc.DataChannel)
	}

	datachannels[room][label][uu] = dc
}

func GetDatachannels(room string, label string) map[string]*webrtc.DataChannel {
	groupe := datachannels[room][label]
	return groupe
}

func deleteDatachannel(room string, uu string) {
	for label := range datachannels[room] {
		delete(datachannels[room][label], uu)
		if len(datachannels[room][label]) == 0 {
			delete(datachannels[room], label)
		}
	}
	if len(datachannels[room]) == 0 {
		delete(datachannels, room)
	}
}
