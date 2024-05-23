import * as React from "react";

import { AuthGuard } from "./AuthGuard";
import { AuthProvider } from "./AuthProvider";
type LayoutProps = Required<{
  readonly children: React.ReactNode;
}>;

export default function Authentication({ children }: LayoutProps) {
  return (
    <>
      <AuthProvider>
        <AuthGuard>
          <main>{children}</main>
        </AuthGuard>
      </AuthProvider>
    </>
  );
}
