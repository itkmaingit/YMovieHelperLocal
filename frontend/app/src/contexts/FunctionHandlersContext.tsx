import { createContext } from "react";

import { FunctionHandlers } from "@/components/utils/LoadingAndSnackbar";

const FunctionHandlersContext = createContext<FunctionHandlers | undefined>(
  undefined
);

export default FunctionHandlersContext;
