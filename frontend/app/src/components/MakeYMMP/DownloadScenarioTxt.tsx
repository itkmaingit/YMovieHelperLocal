import { Box, Button, Grid, Stack, Typography } from "@mui/material";
import saveAs from "file-saver";
import Image from "next/image";
import { useRouter } from "next/router";
import React, { useContext } from "react";

import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import { clientAxios } from "@/libs/axios";

type DownloadScenarioTxtProps = {
  setStep: React.Dispatch<React.SetStateAction<number>>;
};
const DownloadScenarioTxt = ({ setStep }: DownloadScenarioTxtProps) => {
  const router = useRouter();
  const { project_id } = router.query;

  const context = useContext(FunctionHandlersContext);

  if (!context) {
    throw new Error(
      "FunctionHandlersContext is undefined, please verify the Provider"
    );
  }

  const { setCode, setState } = context;
  const handleDownloadScenarioTxt = async () => {
    try {
      const response = await clientAxios.get(
        `/download/${project_id}/scenario_txt`,
        {
          responseType: "blob",
        }
      );
      const blob = new Blob([response.data]);
      saveAs(blob, "scenario.txt");
      setStep(3);
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
              次のリンクからscenario.txtというファイルをダウンロードし、YMM4で読み込み、保存しましょう。
            </Typography>
          </Box>
        </Grid>

        <Grid item md={12} lg={6}>
          <Stack alignItems="center" spacing={2}>
            <Box sx={{ width: "100%", height: "400px", position: "relative" }}>
              <Image
                src="/download.png"
                style={{ objectFit: "contain" }}
                alt="flow_chart"
                fill
              ></Image>
            </Box>
            <Button
              variant="contained"
              color="primary"
              onClick={handleDownloadScenarioTxt}
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

export default DownloadScenarioTxt;
