import BroadcastOnPersonalIcon from "@mui/icons-material/BroadcastOnPersonal";
import {
  Box,
  Divider,
  Drawer,
  Stack,
  Typography,
  useMediaQuery,
} from "@mui/material";
import NextLink from "next/link";
import { useRouter } from "next/router";
import PropTypes from "prop-types";

import { items } from "./config";
import { SideNavItem } from "./SideNavigationItem";

import theme from "@/theme";

type Props = {
  open: boolean;
  onClose: (value: boolean) => void;
};

export const SideNav = (props: Props) => {
  const router = useRouter();
  if (router.isFallback) {
    return <div>Loading...</div>;
  }

  const { open, onClose } = props;

  const xlUp = useMediaQuery(theme.breakpoints.up("xl"));

  const content = (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
        backgroundColor: "black",
      }}
    >
      <Box sx={{ p: 3 }}>
        <Box
          component={NextLink}
          href="/"
          sx={{
            display: "inline-flex",
            height: 32,
            width: 32,
            textDecoration: "none",
            color: "white",
          }}
        >
          <BroadcastOnPersonalIcon
            sx={{ color: "white", fontSize: 45 }}
          ></BroadcastOnPersonalIcon>
          <Typography
            variant="h5"
            sx={{
              transform: "translate(20%,20%)",
            }}
          >
            YMovieHelper
          </Typography>
        </Box>
        <Box
          sx={{
            alignItems: "center",
            backgroundColor: "rgba(255, 255, 255, 0.04)",
            borderRadius: 1,
            cursor: "pointer",
            display: "flex",
            justifyContent: "space-between",
            mt: 2,
            p: "12px",
          }}
        ></Box>
      </Box>
      <Divider sx={{ borderColor: "neutral.700" }} />
      <Box
        component="nav"
        sx={{
          flexGrow: 1,
          px: 2,
          py: 3,
        }}
      >
        <Stack
          component="ul"
          spacing={0.5}
          sx={{
            listStyle: "none",
            p: 0,
            m: 0,
            pl: 5,
          }}
        >
          {items.map((item) => {
            return (
              <SideNavItem
                key={item.title}
                icon={item.icon}
                path={item.path}
                title={item.title}
                belong={item.belong}
                isBlank={item.isBlank}
                isExternal={item.isExternal}
                router={router}
              />
            );
          })}
        </Stack>
      </Box>
      <Divider sx={{ borderColor: "neutral.700" }} />
    </Box>
  );

  if (xlUp) {
    return (
      <Drawer
        anchor="left"
        open
        PaperProps={{
          sx: {
            backgroundColor: "neutral.800",
            color: "common.white",
            width: 300,
          },
        }}
        variant="permanent"
      >
        {content}
      </Drawer>
    );
  }

  return (
    <Drawer
      anchor="left"
      onClose={onClose}
      open={open}
      PaperProps={{
        sx: {
          backgroundColor: "neutral.800",
          color: "common.white",
          width: 300,
        },
      }}
      variant="temporary"
    >
      {content}
    </Drawer>
  );
};

SideNav.propTypes = {
  onClose: PropTypes.func,
  open: PropTypes.bool,
};
