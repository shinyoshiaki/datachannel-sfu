import React, { useRef, useState } from "react";
import axios from "axios";
import WebRTC from "webrtc4me";
import { useApi } from "./hooks/useApi";
import useInput from "./hooks/useInput";
import Send from "./components/send";

const target = "http://localhost:8080";

const App: React.FC = () => {
  const state = useRef({ peer: new WebRTC() });
  const [log, setlog] = useState("");
  const [room, setroom] = useInput();

  const { fetch, loading, error, result } = useApi(
    async () => {
      if (room) {
        const { peer } = state.current;

        if (peer.isConnected) return null;

        peer.makeOffer();
        const { type, sdp } = await peer.onSignal.asPromise();
        const res = await axios.post(target + "/join", { type, sdp, room });

        return res.data as { sdp: any; uu: string };
      }
    },
    async ({ sdp }) => {
      const { peer } = state.current;
      peer.setSdp(sdp);
      await peer.onConnect.asPromise();
      setlog("connected");
      peer.onData.subscribe(v => typeof v.data === "string" && setlog(v.data));
    }
  );

  return (
    <div>
      <input placeholder="room" onChange={setroom} />
      <button onClick={fetch}>offer</button>
      {loading && <p>loading</p>}
      {error && <p>error</p>}
      <p>{log}</p>
      {result && <p>{result.uu}</p>}
      <Send peer={state.current.peer} />
    </div>
  );
};

export default App;
