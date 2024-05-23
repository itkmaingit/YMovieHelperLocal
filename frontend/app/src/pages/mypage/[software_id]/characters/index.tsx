import { Grid, Stack, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import React, { useState } from "react";
import { useDropzone } from "react-dropzone";
import useSWR from "swr";

import CardGrids, { Character } from "@/components/Characters/CardGrids";
import FullScreenLoading from "@/components/utils/FullScreenLoading";
import GlobalAddButton from "@/components/utils/GlobalAddButton";
import InputFileAndUpload from "@/components/utils/InputFileAndUpload";
import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import SectionTitle from "@/components/utils/SectionTitle";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import useFeedback from "@/hooks/use_feedback";
import { clientAxios } from "@/libs/axios";

export default function Characters() {
  const token = parseCookies().token;
  const router = useRouter();
  const { software_id } = router.query;
  const [isHeadsUp, setIsHeadsUp] = useState<boolean>(false);

  const fetcher = async (url: string) => {
    const getConfig = {
      method: "get",
      url: url,
      headers: { Authorization: `Bearer ${token}` },
    };
    const response = await clientAxios(getConfig).catch(
      (error) => error.response
    );
    return response.data;
  };
  const { data, mutate } = useSWR(
    `/softwares/${software_id}/characters`,
    fetcher
  );

  const [dialogOpen, setDialogOpen] = React.useState(false);
  const [selectedFile, setSelectedFile] = React.useState<null | File>(null);

  const dropFunction = async (acceptedFiles: File[]) => {
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

    if (acceptedFiles[0].size > 100000) {
      setState("completed");
      setCode({
        message: "アップロード失敗！ファイルサイズが大きすぎます！",
        status: "warning",
      });
      return;
    }
    setSelectedFile(acceptedFiles[0]);

    const formData = new FormData();
    formData.append("file", acceptedFiles[0]);

    const uploadConfig = {
      method: "post",
      url: `/softwares/${software_id}/characters`,
      data: formData,
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(uploadConfig)
      .then(async () => {
        setCode({ message: "アップロード完了", status: "success" });
      })
      .catch(async (error) => {
        setCode({
          message:
            error.response.data.userError ||
            "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(async () => {
        setDialogOpen(false);
        setSelectedFile(null);
        await mutate();
      });
  };

  const [setCode, setState, snackbarProps] = useFeedback();
  const onDrop = useAsyncAndLoading(dropFunction, setState);

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
  });

  return (
    <>
      <LoadingAndSnackbar {...snackbarProps} />
      <SectionTitle title="Characters"></SectionTitle>
      {/* -----------------------Main-------------------- */}
      <GlobalAddButton
        clickFunction={() => setDialogOpen(true)}
      ></GlobalAddButton>
      <Grid container spacing={3} sx={{ px: 10 }}>
        <FunctionHandlersContext.Provider
          value={{ mutate: mutate, setCode: setCode, setState: setState }}
        >
          {data && data.characters && data.characters.length > 0 ? (
            <>
              {isHeadsUp && (
                <Typography
                  variant="h6"
                  mb={5}
                  gutterBottom
                  color="error"
                  fontWeight="bold"
                >
                  キャラクターの属性を入れ替えた時には、必ずルールを確認し、保存まで行ってください！空読み文が正しく設定されません！
                </Typography>
              )}
              {data.characters.map((character: Character) => (
                <CardGrids
                  key={character.ID}
                  {...character}
                  setIsHeadsUp={setIsHeadsUp}
                ></CardGrids>
              ))}
              <Typography variant="h6" mt={10} gutterBottom>
                ※表示されるキャラクターの画像はYMM4に設定されているキャラクターの画像を表示しているわけではありません。
              </Typography>
            </>
          ) : data === undefined ? (
            <FullScreenLoading open={true}></FullScreenLoading>
          ) : (
            <>
              <Stack mt={5}>
                <Typography fontSize="1.2rem" component="div">
                  キャラクターがまだ作成されていないようです。
                </Typography>
                <Typography fontSize="1.2rem" component="div">
                  画面右下の+ボタンからキャラクターファイルをアップロードしてください。
                </Typography>
              </Stack>
            </>
          )}
        </FunctionHandlersContext.Provider>
      </Grid>
      {/* -----------------------Main-------------------- */}
      {/* -----------------------Dialog-------------------- */}

      <InputFileAndUpload
        dialogOpen={dialogOpen}
        setDialogOpen={setDialogOpen}
        selectedFile={selectedFile}
        getRootProps={getRootProps}
        getInputProps={getInputProps}
        isDragActive={isDragActive}
        title="キャラクターファイルのアップロード"
      ></InputFileAndUpload>
      {/* -----------------------Dialog-------------------- */}
    </>
  );
}
