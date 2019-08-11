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

	setupDatachannel(dc, room, uu)

	peer.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		if connectionState.String() == "disconnected" {
			fmt.Println("disconnect", room, uu)
			store.DeletePeer(room, uu)
		}
		if connectionState.String() == "connected" {
			fmt.Println("connectedd", room, uu)
		}
	})

	peer.OnDataChannel(func(dc *webrtc.DataChannel) {
		dc.OnOpen(func() {
			setupDatachannel(dc, room, uu)
		})
	})

	return offer, uu, nil
}

func setupDatachannel(dc *webrtc.DataChannel, room string, uu string) {
	fmt.Println("dc opened", dc.ReadyState().String())
	store.SetDatachannel(dc, room, uu)
	sfu.Publish(dc, room, uu, func(err error) {
		fmt.Println("dc diconnect", err)
		store.DeletePeer(room, uu)
	})
}

func Answer(room string, uu string, TYPE string, SDP string) error {
	peer := store.GetPeer(room, uu)
	switch TYPE {
	case "candidate":
		ice := webrtc.ICECandidateInit{Candidate: SDP}
		err := peer.AddICECandidate(ice)
		if err != nil {
			return err
		}
	case "offer":
		sdp := webrtc.SessionDescription{Type: webrtc.SDPTypeOffer, SDP: SDP}
		err := peer.SetRemoteDescription(sdp)
		if err != nil {
			return err
		}
	case "answer":
		sdp := webrtc.SessionDescription{Type: webrtc.SDPTypeAnswer, SDP: SDP}
		err := peer.SetRemoteDescription(sdp)
		if err != nil {
			return err
		}
	}
	return nil
}
