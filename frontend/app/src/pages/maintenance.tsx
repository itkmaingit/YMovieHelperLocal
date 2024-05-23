import { Box, Stack, Typography } from "@mui/material";
import { ReactElement } from "react";

import NotDashboardLayout from "@/layouts/NotDashboardLayout";

import { NextPageWithLayout } from "./_app";

const Sorry: NextPageWithLayout = () => {
  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        mt: 10,
        flexDirection: "column",
      }}
    >
      <Typography variant="h3" mb={10}>
        503 Service Temporarily Unavailable
      </Typography>
      <Stack
        direction="row"
        spacing={5}
        sx={{ alignItems: "center", justifyContent: "space-between" }}
      >
        <img
          src="https://d36fgadpjsedj7.cloudfront.net/static/sorry_marisa.gif"
          alt="sorry_marisa"
          height="640"
          width="452"
        ></img>
        <Stack spacing={5}>
          <Typography variant="h5">
            申し訳ありません...ただいまサーバーメンテナンス中です。
            <br />
            しばらく時間をおいてからお試しください。
          </Typography>
        </Stack>
      </Stack>
    </Box>
  );
};

Sorry.getLayout = function getLayout(page: ReactElement) {
  return <NotDashboardLayout>{page}</NotDashboardLayout>;
};

export default Sorry;
