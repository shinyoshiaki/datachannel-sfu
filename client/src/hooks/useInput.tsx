import { useState } from "react";

export default function useInput(
  init?: string
): [string, (e: any) => void, () => void] {
  const [value, setvalue] = useState(init || "");
  const input = (e: any) => {
    setvalue(e.target.value);
  };

  const clear = () => {
    setvalue("");
  };

  return [value, input, clear];
}
