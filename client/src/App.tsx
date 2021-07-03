import React, { useRef, useState } from "react";
import axios from "axios";
import WebRTC from "webrtc4me";
import { useApi } from "./hooks/useApi";
import useInput from "./hooks/useInput";
import Send from "./components/send";
import SfuVideo from "./containers/sfu-video";

const url =
  process.env.NODE_ENV === "production"
    ? "http://34.73.92.59:8088"
    : "http://localhost:8088";

const App: React.FC = () => {
  const state = useRef({ peer: new WebRTC({ trickle: true }) });
  const [log, setlog] = useState("");
  const [room, setroom] = useInput("a");

  const { fetch, loading, error, result } = useApi(
    async () => {
      if (room) {
        const { peer } = state.current;

        if (peer.isConnected) return null;

        const res = await axios.post<{ sdp: any; uu: string }>(url + "/join", {
          room
        });
        const { sdp, uu } = res.data;
        peer.setSdp(sdp);
        await new Promise(r => {
          peer.onSignal.subscribe(({ sdp, type, ice }) => {
            sdp = sdp ? sdp : (ice!.candidate as any);
            console.log({ sdp });
            axios.post(url + "/signaling", { type, sdp, room, uu });
          });
          peer.onConnect.once(r);
        });
        return { uu };
      }
    },
    () => {
      console.log("connected");
      const { peer } = state.current;
      peer.onData.subscribe(v => {
        typeof v.data === "string" && setlog(v.data);
      });
    }
  );

  return (
    <div>
      <input placeholder="room" onChange={setroom} value={room} />
      <button onClick={fetch}>join</button>
      {loading && <p>loading</p>}
      {error && <p>error</p>}
      {result && <p>{result.uu}</p>}
      <Send peer={state.current.peer} />
      <p>{log}</p>
      <SfuVideo peer={state.current.peer} />
    </div>
  );
};

export default App;
