import React, { useRef, useState } from "react";
import axios from "axios";
import WebRTC from "webrtc4me";
import { useApi } from "./hooks/useApi";
import useInput from "./hooks/useInput";
import Send from "./components/send";

const target = "http://localhost:8080";

const App: React.FC = () => {
  const state = useRef({ peer: new WebRTC({ trickle: true }) });
  const [log, setlog] = useState("");
  const [room, setroom] = useInput();

  const { fetch, loading, error, result } = useApi(
    async () => {
      if (room) {
        const { peer } = state.current;

        if (peer.isConnected) return null;

        const res = await axios.post<{ sdp: any; uu: string }>(
          target + "/join",
          {
            room
          }
        );
        const { sdp, uu } = res.data;
        console.log({ sdp, uu });
        peer.setSdp(sdp);
        await new Promise(r => {
          peer.onSignal.subscribe(({ sdp, type, ice }) => {
            sdp = sdp ? sdp : (ice as any);
            axios.post(target + "/answer", { type, sdp, room, uu });
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
        console.log({ v });
        typeof v.data === "string" && setlog(v.data);
      });
    }
  );

  return (
    <div>
      <input placeholder="room" onChange={setroom} />
      <button onClick={fetch}>join</button>
      {loading && <p>loading</p>}
      {error && <p>error</p>}
      {result && <p>{result.uu}</p>}
      <Send peer={state.current.peer} />
      <p>{log}</p>
    </div>
  );
};

export default App;
