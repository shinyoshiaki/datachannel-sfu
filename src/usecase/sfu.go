package usecase

import (
	"data-sfu/src/domain/sfu"
	"data-sfu/src/domain/store"
	"data-sfu/src/domain/webrtc/peer"
	"encoding/json"
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Join(room string) (webrtc.SessionDescription, string, error) {

	create := make(chan peer.CreatePeerRes)
	dcOpen := make(chan bool)
	go peer.CreatePeer(create, dcOpen)
	res := <-create

	uu, err := store.SetPeer(res.Peer, room)
	if err != nil {
		fmt.Println("error", err)
		return webrtc.SessionDescription{}, "", err
	}

	res.Peer.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		if connectionState.String() == "disconnected" {
			fmt.Println("disconnect", room, uu)
			store.DeletePeer(room, uu)
		}
		if connectionState.String() == "connected" {
			fmt.Println("connectedd", room, uu)
		}
	})

	setupDatachannel(res.DC, res.Peer, room, uu)
	res.DC.OnOpen(func() {
		dcOpen <- true
	})

	res.Peer.OnDataChannel(func(dc *webrtc.DataChannel) {
		setupDatachannel(dc, res.Peer, room, uu)
	})

	return res.Offer, uu, nil
}

func setupDatachannel(dc *webrtc.DataChannel, pc *webrtc.PeerConnection, room string, uu string) {
	store.SetDatachannel(dc, room, dc.Label(), uu)

	onsdp := make(chan peer.Sdp)
	go peer.Listen(pc, onsdp)

	sfu.Publish(dc, room, uu,
		func(err error) {
			fmt.Println("dc diconnect", err)
			store.DeletePeer(room, uu)
		},
		func(data []byte) {
			var sdp peer.Sdp
			json.Unmarshal(data, &sdp)
			onsdp <- sdp
		})
}

func Answer(room string, uu string, TYPE string, SDP string) error {
	pc := store.GetPeer(room, uu)
	sdp := &peer.Sdp{Type: TYPE, Sdp: SDP}
	err := peer.SetSDP(sdp, pc)

	if err != nil {
		return err
	}

	return nil
}
