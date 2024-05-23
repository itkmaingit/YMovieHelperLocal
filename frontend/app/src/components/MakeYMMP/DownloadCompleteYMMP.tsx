import { Box, Button, Grid, Stack, Typography } from "@mui/material";
import saveAs from "file-saver";
import Image from "next/image";
import { useRouter } from "next/router";
import React, { useContext } from "react";

import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import { clientAxios } from "@/libs/axios";

type DownloadCompleteYMMPProps = {
  setStep: React.Dispatch<React.SetStateAction<number>>;
};
const DownloadCompleteYMMP = ({ setStep }: DownloadCompleteYMMPProps) => {
  const router = useRouter();
  const { project_id } = router.query;

  const context = useContext(FunctionHandlersContext);

  if (!context) {
    throw new Error(
      "FunctionHandlersContext is undefined, please verify the Provider"
    );
  }

  const { setCode, setState } = context;
  const handleDownloadCompleteYMMP = async () => {
    try {
      const response = await clientAxios.get(
        `/download/${project_id}/complete_ymmp`,
        {
          responseType: "blob",
        }
      );
      const blob = new Blob([response.data]);
      saveAs(blob, "complete.ymmp");
      setStep(0);
    } catch (error) {
      setState("completed");
      setCode({
        message: "サーバーエラー！時間をおいてから試して下さい！",
        status: "warning",
      });
    }
  };
  return (
    <>
      <Grid container>
        <Grid item md={12} lg={6}>
          <Box
            sx={{
              width: "100%",
              height: "400px",
              position: "relative",
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              pl: 10,
            }}
          >
            <Typography variant="h5">
              お疲れ様です。complete.ymmpをダウンロードして、ご自身のYMM4で読み込んでください。
            </Typography>
          </Box>
        </Grid>

        <Grid item md={12} lg={6}>
          <Stack alignItems="center" spacing={2}>
            <Box sx={{ width: "100%", height: "400px", position: "relative" }}>
              <Image
                src="/complete.png"
                style={{ objectFit: "contain" }}
                alt="flow_chart"
                fill
              ></Image>
            </Box>
            <Button
              variant="contained"
              color="primary"
              onClick={handleDownloadCompleteYMMP}
              sx={{ width: "70%" }}
              size="large"
            >
              Download
            </Button>
          </Stack>
        </Grid>
      </Grid>
    </>
  );
};

export default DownloadCompleteYMMP;
