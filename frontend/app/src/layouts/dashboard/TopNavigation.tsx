import Bars3Icon from "@heroicons/react/24/solid/Bars3Icon";
import { Box, IconButton, Stack, SvgIcon, useMediaQuery } from "@mui/material";
import PropTypes from "prop-types";

import theme from "@/theme";

const SIDE_NAV_WIDTH = 300;
const TOP_NAV_HEIGHT = 64;

type Props = {
  onNavOpen: () => void;
};
export const TopNav = (props: Props) => {
  const { onNavOpen } = props;
  const xlUp = useMediaQuery(theme.breakpoints.up("xl"));

  return (
    <>
      <Box
        component="header"
        sx={{
          position: "sticky",
          left: {
            xl: `${SIDE_NAV_WIDTH}px`,
          },
          top: 0,
        }}
      >
        <Stack
          alignItems="center"
          direction="row"
          justifyContent="space-between"
          spacing={2}
          sx={{
            minHeight: TOP_NAV_HEIGHT,
            px: 2,
          }}
        >
          <Stack alignItems="center" direction="row" spacing={2}>
            {!xlUp && (
              <IconButton onClick={onNavOpen}>
                <SvgIcon fontSize="large">
                  <Bars3Icon />
                </SvgIcon>
              </IconButton>
            )}
          </Stack>
        </Stack>
      </Box>
    </>
  );
};

TopNav.propTypes = {
  onNavOpen: PropTypes.func,
};
