import { useState } from "react";

const useAsync = <T,>(
  asyncFunction: (...args: any[]) => Promise<T>
): [(...args: any[]) => Promise<T>, boolean] => {
  const [loading, setLoading] = useState(false);

  const execute = async (...args: any[]): Promise<T> => {
    setLoading(true);
    const result = await asyncFunction(...args);
    setLoading(false);
    return result;
  };

  return [execute, loading];
};

export default useAsync;
