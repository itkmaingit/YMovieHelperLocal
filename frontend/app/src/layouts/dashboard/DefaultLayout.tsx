/** @jsxImportSource @emotion/react */

import { styled } from "@mui/material";
import * as React from "react";

import { SideNav } from "./SideNavigation";
import { TopNav } from "./TopNavigation";

import CommonLayout from "@/layouts/CommonLayout";
import Footer from "@/layouts/Footer";

const SIDE_NAV_WIDTH = 350;
const TOP_MARGIN_HEIGHT = 50;

const LayoutRoot = styled("div")(({ theme }) => ({
  display: "flex",
  flex: "1 1 auto",
  maxWidth: "85%",
  paddingLeft: 200,
  [theme.breakpoints.up("xl")]: {
    paddingLeft: SIDE_NAV_WIDTH + 200,
  },
  marginTop: TOP_MARGIN_HEIGHT,
  flexDirection: "column",
}));

const LayoutContainer = styled("div")({
  display: "flex",
  flex: "1 1 auto",
  flexDirection: "column",
});

type LayoutProps = Required<{
  readonly children: React.ReactNode;
}>;
const DefaultLayout = ({ children }: LayoutProps) => {
  const [openNav, setOpenNav] = React.useState(false);

  return (
    <CommonLayout>
      <TopNav onNavOpen={() => setOpenNav(true)} />
      <SideNav onClose={() => setOpenNav(false)} open={openNav} />
      <LayoutRoot>
        <LayoutContainer>{children}</LayoutContainer>
      </LayoutRoot>
      <Footer />
    </CommonLayout>
  );
};

export default DefaultLayout;
