import { Box } from "@mui/material";
import { useRouter } from "next/router";
import * as React from "react";

import { useAuthContext } from "./AuthProvider";

type Props = {
  children: React.ReactNode;
};

export const AuthGuard = ({ children }: Props) => {
  const { user } = useAuthContext();
  const router = useRouter();

  if (typeof user === "undefined") {
    return <Box>読み込み中...</Box>;
  }

  if (user === null) {
    router.replace("/login");
    return null;
  }

  return <>{children}</>;
};
