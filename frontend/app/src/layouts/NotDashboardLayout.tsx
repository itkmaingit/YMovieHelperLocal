import * as React from "react";

import CommonLayout from "@/layouts/CommonLayout";

type LayoutProps = Required<{
  readonly children: React.ReactNode;
}>;
const NotDashboardLayout = ({ children }: LayoutProps) => {
  return <CommonLayout>{children}</CommonLayout>;
};

export default NotDashboardLayout;
