import { Alert, Snackbar } from "@mui/material";
import * as React from "react";
import { KeyedMutator } from "swr";

import FullScreenLoading from "./FullScreenLoading";

export type ResponseStatus = {
  message: string;
  status: "success" | "warning";
};

export type LoadingAndSnackbarProps = {
  code: ResponseStatus;
  state: "none" | "loading" | "completed";
  setCode: React.Dispatch<React.SetStateAction<ResponseStatus>>;
  setState: React.Dispatch<
    React.SetStateAction<"none" | "loading" | "completed">
  >;
};

export type FunctionHandlers = {
  mutate?: KeyedMutator<any>;
  setCode: React.Dispatch<React.SetStateAction<ResponseStatus>>;
  setState: React.Dispatch<
    React.SetStateAction<"none" | "loading" | "completed">
  >;
};

const LoadingAndSnackbar = ({
  code,
  state,
  setCode,
  setState,
}: LoadingAndSnackbarProps) => {
  const [loadingOpen, setLoadingOpen] = React.useState(false);
  const [snackbarOpen, setSnackbarOpen] = React.useState(false);

  const handleSnackbarClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }

    setSnackbarOpen(false);
    setState("none");
    setCode({ message: "", status: "success" });
  };

  React.useEffect(() => {
    if (state === "loading") {
      setLoadingOpen(true);
    } else if (state === "completed") {
      setLoadingOpen(false);
      setSnackbarOpen(true);
    }
  }, [state]);

  return (
    <>
      <FullScreenLoading open={loadingOpen}></FullScreenLoading>
      {code.message !== "" && (
        <Snackbar
          open={snackbarOpen}
          autoHideDuration={3000}
          onClose={handleSnackbarClose}
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert
            onClose={handleSnackbarClose}
            severity={code.status}
            sx={{ width: "100%" }}
          >
            {code.message}
          </Alert>
        </Snackbar>
      )}
    </>
  );
};

export default LoadingAndSnackbar;
