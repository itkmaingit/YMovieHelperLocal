export type UseAsyncAndLoadingProps = {
  setState: React.Dispatch<
    React.SetStateAction<"none" | "loading" | "completed">
  >;
};

const useAsyncAndLoading = <T, Args extends any[]>(
  asyncFunction: (...args: Args) => Promise<T>,
  setState: React.Dispatch<
    React.SetStateAction<"none" | "loading" | "completed">
  >
): ((...args: Args) => Promise<T>) => {
  const execute = async (...args: Args): Promise<T> => {
    setState("loading");
    const result = await asyncFunction(...args);
    setState("completed");
    return result;
  };

  return execute;
};

export default useAsyncAndLoading;
