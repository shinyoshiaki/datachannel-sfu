import { useState } from "react";

type Head<T extends any[], D = never> = T extends [infer X, ...any[]] ? X : D;

export function useApi<T, A>(
  promise: (obj: A) => Promise<T>,
  onResult?: (res: T) => void
) {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);
  const [result, setresult] = useState<T>();

  const fetch = async (p: Head<Parameters<typeof promise>>) => {
    setLoading(true);
    const res = p ? await promise(p as any) : await promise(undefined as any);
    if (!res) setError(true);

    setLoading(false);
    setresult(res);
    onResult && onResult(res);
    return res;
  };

  return { loading, fetch, error, result };
}
