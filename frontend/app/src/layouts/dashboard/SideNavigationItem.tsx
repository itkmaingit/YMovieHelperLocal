import { Box, ButtonBase } from "@mui/material";
import { NextRouter } from "next/router";
import PropTypes from "prop-types";
import * as React from "react";

import { Belong } from "./config";

import handlePageTransition from "@/feature/page_transition";

type Props = {
  icon: React.ReactNode;
  path: string;
  title: string;
  belong: Belong;
  isBlank: boolean;
  isExternal: boolean;
  router: NextRouter;
};

type ButtonState = "Focused" | "Disabled" | "Enabled";

export const SideNavItem = ({
  icon,
  path,
  title,
  belong,
  router,
  isBlank,
  isExternal,
}: Props) => {
  const currentURL = router.pathname;
  const softwareID = router.query.software_id as string;
  const projectID = router.query.project_id as string;
  // 今、どこにいるかを確かめる
  let currentLocation: Belong = "None";
  if (softwareID === undefined && projectID === undefined) {
    currentLocation = "None";
  } else if (projectID === undefined) {
    currentLocation = "Software";
  } else {
    currentLocation = "Project";
  }

  // Belongを見て、自分より下の存在にはアクセスできないようにする
  let clickDenied = false;
  if (belong === "Software") {
    if (currentLocation === "None") {
      clickDenied = true;
    }
  } else if (belong === "Project") {
    if (currentLocation === "None" || currentLocation === "Software") {
      clickDenied = true;
    }
  }

  // URLを見て、今自分がFocusが当たっているかどうかを確かめる
  const isFocused = currentURL.includes(path);

  // Belongから、URLを動的に生成
  let builtURL = "";
  if (belong === "None") {
    builtURL = path;
  } else if (belong === "Software") {
    builtURL = `/mypage/${softwareID}/${path}`;
  } else if (belong === "Project") {
    builtURL = `/mypage/${softwareID}/${projectID}/${path}`;
  }

  // 種々の条件から、ボタンの可否を判定
  let state: ButtonState = "Enabled";
  if (clickDenied) {
    state = "Disabled";
  }
  if (isFocused) {
    state = "Focused";
  }

  let buttonColor, fontColor, backgroundColor, hoverColor: string;
  if (state === "Focused") {
    buttonColor = "primary.main";
    fontColor = "neutral.400";
    backgroundColor = "rgba(255, 255, 255, 0.2)";
    hoverColor = "rgba(255, 255, 255, 0.2)";
  } else if (state === "Enabled") {
    buttonColor = "neutral.400";
    fontColor = "neutral.400";
    backgroundColor = "";
    hoverColor = "rgba(255, 255, 255, 0.2)";
  } else if (state === "Disabled") {
    buttonColor = "grey";
    fontColor = "grey";
    backgroundColor = "";
    hoverColor = "";
  } else {
    buttonColor = "";
    fontColor = "";
    backgroundColor = "";
    hoverColor = "";
  }

  const handleClick = handlePageTransition(
    router,
    builtURL,
    isBlank,
    isExternal
  );

  return (
    <li>
      <ButtonBase
        disabled={clickDenied}
        sx={{
          alignItems: "center",
          borderRadius: 1,
          display: "flex",
          justifyContent: "flex-start",
          pl: "16px",
          pr: "16px",
          py: "6px",
          my: 0.5,
          textAlign: "left",
          width: "100%",
          backgroundColor: backgroundColor,
          "&:hover": {
            backgroundColor: hoverColor,
          },
        }}
        onClick={handleClick}
      >
        {icon && (
          <Box
            component="span"
            sx={{
              alignItems: "center",
              display: "inline-flex",
              justifyContent: "center",
              mr: 2,
              color: buttonColor,
            }}
          >
            {icon}
          </Box>
        )}
        <Box
          component="span"
          sx={{
            flexGrow: 1,
            fontSize: 20,
            lineHeight: "24px",
            whiteSpace: "nowrap",
            color: fontColor,
          }}
        >
          {title}
        </Box>
      </ButtonBase>
    </li>
  );
};

SideNavItem.propTypes = {
  active: PropTypes.bool,
  disabled: PropTypes.bool,
  external: PropTypes.bool,
  icon: PropTypes.node,
  path: PropTypes.string,
  title: PropTypes.string.isRequired,
};
