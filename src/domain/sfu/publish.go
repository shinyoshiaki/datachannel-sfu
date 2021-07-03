package sfu

import (
	"data-sfu/src/domain/store"
	"fmt"

	"github.com/pion/webrtc/v2"
)

func Publish(dc *webrtc.DataChannel, room string, uu string, onError func(error), onUpdate func([]byte)) {
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		groupe := store.GetDatachannels(room, dc.Label())
		if dc.Label() == "update" {
			onUpdate(msg.Data)
			return
		}
		fmt.Println("publish", string(msg.Data))
		for k, v := range groupe {
			if k != uu && v != nil {
				defer func() {
					if err := recover(); err != nil {
						onError(err.(error))
					}
				}()
				if msg.IsString == true {
					v.SendText(string(msg.Data))
				} else {
					v.Send(msg.Data)
				}

			}
		}
	})
}
