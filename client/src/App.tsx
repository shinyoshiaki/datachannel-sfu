import React, { useRef } from "react";
import axios from "axios";
import WebRTC from "webrtc4me";
import { useApi } from "./hooks/useApi";

const target = "http://localhost:8080";

const App: React.FC = () => {
  const state = useRef(
    new (class {
      peer = new WebRTC();
    })()
  );

  const { fetch, loading, error, result } = useApi(
    async () => {
      const { peer } = state.current;
      peer.makeOffer();
      const sdp = await peer.onSignal.asPromise();
      const res = await axios.post(target + "/join", {
        type: sdp.type,
        sdp: sdp.sdp
      });
      console.log({ res });
      return res.data as { sdp: any; uu: string };
    },
    async ({ sdp, uu }) => {
      const { peer } = state.current;
      console.log({ sdp, uu });
      peer.setSdp(sdp);
      await peer.onConnect.asPromise();
      console.log("connected");
    }
  );

  return (
    <div>
      <button onClick={fetch}>offer</button>
      {result && <p>{result.uu}</p>}
    </div>
  );
};

export default App;
