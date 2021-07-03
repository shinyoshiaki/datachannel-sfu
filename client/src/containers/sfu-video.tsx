import React, { FC } from "react";
import WebRTC, { getLocalVideo } from "webrtc4me";
import Videos from "../components/videos";
import Event from "rx.mini";
import { useAsyncEffect } from "../hooks/useAsyncEffect";

const SfuVideo: FC<{ peer: WebRTC }> = ({ peer }) => {
  const streamEvent = new Event<MediaStream>();

  useAsyncEffect(async () => {
    await peer.onConnect.asPromise();
    const stream = await getLocalVideo();
    peer.addTrack(stream.getVideoTracks()[0], stream);

    const { unSubscribe } = peer.onAddTrack.subscribe(stream =>
      streamEvent.execute(stream)
    );

    peer.onDisconnect.once(unSubscribe);
  }, []);

  return (
    <div>
      <p>sfu videos</p>
      <Videos streamEvent={streamEvent} />
    </div>
  );
};

export default SfuVideo;
