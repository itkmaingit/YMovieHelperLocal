import { Button, Stack } from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import * as React from "react";

import InputFile from "@/components/utils/InputFile";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

const MultipleItem = () => {
  const router = useRouter();
  const { software_id, project_id } = router.query;
  const [selectedFile, setSelectedFile] = React.useState<null | File>(null);
  const [uploadMultipleItemName, setUploadMultipleItemName] =
    React.useState<string>("");
  const context = React.useContext(FunctionHandlersContext);

  if (!context) {
    throw new Error(
      "FunctionHandlersContext is undefined, please verify the Provider"
    );
  }

  const { setCode, setState } = context;

  const buttonDisabled =
    selectedFile === null ||
    isEmptyOrWhitespace(uploadMultipleItemName) ||
    uploadMultipleItemName.length > 20;

  const onDrop = (acceptedFiles: File[]) => {
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
  };

  const uploadFunction = async () => {
    if (!selectedFile) return;

    const formData = new FormData();
    formData.append("file", selectedFile);
    formData.append("name", uploadMultipleItemName);

    const token = parseCookies().token;
    const uploadConfig = {
      method: "post",
      url: `/softwares/${software_id}/projects/${project_id}/items/upload_multiple_item`,
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${token}`,
      },
    };
    await clientAxios(uploadConfig)
      .then(() => {
        setCode({
          message: "アップロード成功",
          status: "success",
        });
        setUploadMultipleItemName("");
      })
      .catch(async (error) => {
        setCode({
          message:
            error.response.data.userError ||
            "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(() => {
        setSelectedFile(null);
      });
  };

  const handleUploadFunction = useAsyncAndLoading(uploadFunction, setState);

  return (
    <>
      <Stack spacing={3}>
        <InputFile
          displayName={"Multiple Itemの名前"}
          inputText={uploadMultipleItemName}
          setInputText={setUploadMultipleItemName}
          selectedFile={selectedFile}
          onDrop={onDrop}
        >
          <Button
            type="submit"
            variant="contained"
            color="primary"
            size="large"
            disabled={buttonDisabled}
            onClick={handleUploadFunction}
          >
            アップロード
          </Button>
        </InputFile>
      </Stack>
    </>
  );
};

export default MultipleItem;
