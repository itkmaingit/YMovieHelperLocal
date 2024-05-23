import { Box, Button, Stack, Typography } from "@mui/material";
import Image from "next/image";
import { useRouter } from "next/router";
import { ReactElement } from "react";

import handlePageTransition from "@/feature/page_transition";
import NotDashboardLayout from "@/layouts/NotDashboardLayout";

import { NextPageWithLayout } from "./_app";

const Sorry: NextPageWithLayout = () => {
  const router = useRouter();
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
        500 Internet Server Error
      </Typography>
      <Stack
        direction="row"
        spacing={5}
        sx={{ alignItems: "center", justifyContent: "space-between" }}
      >
        <Stack spacing={5}>
          <Typography variant="h5">
            申し訳ありません...ただいまサーバーメンテナンス中です。
            <br />
            しばらく時間をおいてからお試しください。
          </Typography>
          <Button
            variant="contained"
            onClick={handlePageTransition(router, "/")}
          >
            トップページへ
          </Button>
        </Stack>
        <Image
          src="/sorry_reimu.gif"
          height="640"
          width="452"
          alt="sorry_reimu"
        ></Image>
      </Stack>
    </Box>
  );
};

Sorry.getLayout = function getLayout(page: ReactElement) {
  return <NotDashboardLayout>{page}</NotDashboardLayout>;
};

export default Sorry;
