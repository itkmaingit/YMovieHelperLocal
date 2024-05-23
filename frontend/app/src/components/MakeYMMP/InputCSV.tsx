import { Box, Grid, Link, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import * as React from "react";

import InputFile from "../utils/InputFile";

import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";

type InputCSVProps = {
  setStep: React.Dispatch<React.SetStateAction<number>>;
};

const InputCSV = ({ setStep }: InputCSVProps) => {
  const [selectedCSV, setSelectedCSV] = React.useState<File | null>(null);

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
    if (extension !== "csv") {
      setState("completed");
      setCode({
        message:
          "アップロード失敗！ファイル形式はcsvのものしか受け付けられません！",
        status: "warning",
      });
      return;
    }

    if (acceptedFiles[0].size > 100000) {
      setState("completed");
      setCode({
        message: "アップロード失敗！ファイルサイズが大きすぎます！",
        status: "warning",
      });
      return;
    }

    setSelectedCSV(acceptedFiles[0]);
    const formData = new FormData();
    formData.append("file", acceptedFiles[0]);

    const uploadConfig = {
      method: "post",
      url: `/softwares/${software_id}/projects/${project_id}/make_ymmp/resolve_scenario`,
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${token}`,
      },
    };

    await clientAxios(uploadConfig)
      .then(async () => {
        setCode({ message: "アップロード完了", status: "success" });
        setStep(2);
      })
      .catch(async (error) => {
        setCode({
          message:
            error.response.data.userError ||
            "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(() => setSelectedCSV(null));
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
              作成した台本をアップロードしましょう。
              どのような台本をアップロードすればいいか分からない時は、
              <Link href="/how_to" target="_blank">
                How toページ
              </Link>
              で確認してください。
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
            <InputFile selectedFile={selectedCSV} onDrop={wrapOnDrop}>
              <></>
            </InputFile>
          </Box>
        </Grid>
      </Grid>
    </>
  );
};

export default InputCSV;
