package sfu

import (
	"data-sfu/src/domain/store"
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Publish(peer *webrtc.PeerConnection, room string, uu string) {
	peer.OnDataChannel(func(d *webrtc.DataChannel) {
		store.SetDatachannel(d, room, uu)

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Println(msg)
			groupe := store.GetDatachannels(room)
			for k, v := range groupe {
				if k != uu {
					if msg.IsString == true {
						v.SendText(string(msg.Data))
					} else {
						v.Send(msg.Data)
					}

				}
			}
		})
	})
}
