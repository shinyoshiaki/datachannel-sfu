import React, { useRef, useState } from "react";
import axios from "axios";
import WebRTC from "webrtc4me";
import { useApi } from "./hooks/useApi";

const target = "http://localhost:8080";

const App: React.FC = () => {
  const state = useRef({ peer: new WebRTC() });
  const [log, setlog] = useState("");

  const { fetch, loading, error, result } = useApi(
    async () => {
      const { peer } = state.current;

      if (peer.isConnected) return null;

      peer.makeOffer();
      const sdp = await peer.onSignal.asPromise();
      const res = await axios.post(target + "/join", {
        type: sdp.type,
        sdp: sdp.sdp
      });
      console.log({ res });
      return res.data as { sdp: any; uu: string };
    },
    async ({ sdp }) => {
      const { peer } = state.current;
      peer.setSdp(sdp);
      await peer.onConnect.asPromise();
      setlog("connected");
    }
  );

  return (
    <div>
      <button onClick={fetch}>offer</button>
      {loading && <p>loading</p>}
      {error && <p>error</p>}
      <p>{log}</p>
      {result && <p>{result.uu}</p>}
    </div>
  );
};

export default App;
