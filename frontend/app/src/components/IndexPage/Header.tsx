import { Box, Button, Stack, Typography } from "@mui/material";
import { useRouter } from "next/router";

import handlePageTransition from "@/feature/page_transition";

const Header = () => {
  const router = useRouter();
  return (
    <Box
      component="header"
      sx={{
        backdropFilter: "blur(6px)",
        backgroundColor: "rgba(255, 255, 255,0.9)",
        position: "sticky",
        top: 0,
        width: "100%",
        zIndex: 1000,
      }}
    >
      <Stack
        alignItems="center"
        direction="row"
        justifyContent="space-between"
        spacing={1}
        sx={{
          minHeight: 100,
          px: "2vw",
        }}
      >
        {/* space-betweenのためのダミーStack */}
        <Stack alignItems="center" direction="row" spacing={2}></Stack>
        <Stack alignItems="center" direction="row">
          <Button size="large" onClick={handlePageTransition(router, "/login")}>
            <Typography sx={{ fontSize: { xs: "1.0rem", md: "1.3rem" } }}>
              Login
            </Typography>
          </Button>
          <Button
            size="large"
            onClick={handlePageTransition(router, "/login/signup")}
          >
            <Typography sx={{ fontSize: { xs: "1.0rem", md: "1.3rem" } }}>
              Sign up
            </Typography>
          </Button>
        </Stack>
      </Stack>
    </Box>
  );
};

export default Header;
