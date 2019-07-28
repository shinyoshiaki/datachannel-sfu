import React, { FC } from "react";
import useInput from "../hooks/useInput";
import WebRTC from "webrtc4me";

const Send: FC<{ peer: WebRTC }> = ({ peer }) => {
  const [msg, setmsg, clear] = useInput();

  const send = () => {
    if (peer.isConnected) {
      peer.send(msg);
      clear();
    }
  };

  return (
    <div>
      <input placeholder="send" onChange={setmsg} value={msg} />
      <button onClick={send}>send</button>
    </div>
  );
};

export default Send;
