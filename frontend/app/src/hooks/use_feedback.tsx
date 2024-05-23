import { useState } from "react";

import {
  LoadingAndSnackbarProps,
  ResponseStatus,
} from "@/components/utils/LoadingAndSnackbar";

const useFeedback = (): [
  React.Dispatch<React.SetStateAction<ResponseStatus>>,
  React.Dispatch<React.SetStateAction<"none" | "loading" | "completed">>,
  LoadingAndSnackbarProps
] => {
  const [code, setCode] = useState<ResponseStatus>({
    message: "",
    status: "success",
  });
  const [state, setState] = useState<"none" | "loading" | "completed">("none");

  const snackbarProps: LoadingAndSnackbarProps = {
    code: code,
    state: state,
    setCode: setCode,
    setState: setState,
  };

  return [setCode, setState, snackbarProps];
};

export default useFeedback;
