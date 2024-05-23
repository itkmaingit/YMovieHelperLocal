import { Box, Grid, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import * as React from "react";

import InputFile from "../utils/InputFile";

import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";

type InputYMMPProps = {
  setStep: React.Dispatch<React.SetStateAction<number>>;
};

const InputYMMP = ({ setStep }: InputYMMPProps) => {
  const [selectedYMMP, setSelectedYMMP] = React.useState<File | null>(null);

  const token = parseCookies().token;
  const router = useRouter();
  const { software_id, project_id } = router.query;
  const context = React.useContext(FunctionHandlersContext);

  if (!context) {
    throw new Error(
      "FunctionHandlersContext is undefined, please verify the Provider"
    );
  }
  const { setCode, setState } = context;

  const onDrop = async (acceptedFiles: File[]) => {
    if (!acceptedFiles[0]) return;

    const extension = acceptedFiles[0].name.split(".").pop();
    if (extension !== "ymmp") {
      setState("completed");
      setCode({
        message:
          "アップロード失敗！ファイル形式はymmpのものしか受け付けられません！",
        status: "warning",
      });
      return;
    }

    if (acceptedFiles[0].size > 3000000) {
      setState("completed");
      setCode({
        message:
          "アップロード失敗！ファイルサイズが大きすぎます！ボイスキャッシュが含まれていないか確認してください！",
        status: "warning",
      });
      return;
    }

    setSelectedYMMP(acceptedFiles[0]);
    const formData = new FormData();
    const fileName = acceptedFiles[0].name;
    const lastDotPosition = fileName.lastIndexOf(".");
    const fileNameWithoutExtension = fileName.slice(0, lastDotPosition);
    formData.append("file", acceptedFiles[0]);
    formData.append("movieName", fileNameWithoutExtension);

    const uploadConfig = {
      method: "post",
      url: `/softwares/${software_id}/projects/${project_id}/make_ymmp/resolve_ymmp`,
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${token}`,
      },
    };

    await clientAxios(uploadConfig)
      .then(async () => {
        setCode({ message: "アップロード完了", status: "success" });
        setStep(4);
      })
      .catch(async (error) => {
        setCode({
          message:
            error.response.data.userError ||
            "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(() => setSelectedYMMP(null));
  };
  const wrapOnDrop = useAsyncAndLoading(onDrop, setState);

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
              保存したymmpファイルを加工せず、そのままアップロードしてください。
            </Typography>
          </Box>
        </Grid>

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
              flexDirection: "column",
            }}
          >
            <InputFile selectedFile={selectedYMMP} onDrop={wrapOnDrop}>
              <></>
            </InputFile>
          </Box>
        </Grid>
      </Grid>
    </>
  );
};

export default InputYMMP;
