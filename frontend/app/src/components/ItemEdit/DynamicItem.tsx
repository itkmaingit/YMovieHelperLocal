import {
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
  Stack,
} from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import * as React from "react";

import InputFile from "@/components/utils/InputFile";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

const DynamicItem = () => {
  const router = useRouter();
  const { software_id, project_id } = router.query;
  const [selectedFile, setSelectedFile] = React.useState<null | File>(null);
  const [uploadDynamicItemName, setUploadDynamicItemName] =
    React.useState<string>("");
  const [itemSelect, setItemSelect] = React.useState("");
  const context = React.useContext(FunctionHandlersContext);

  if (!context) {
    throw new Error(
      "FunctionHandlersContext is undefined, please verify the Provider"
    );
  }

  const { setCode, setState } = context;

  const buttonDisabled =
    itemSelect === "" ||
    selectedFile === null ||
    isEmptyOrWhitespace(uploadDynamicItemName) ||
    uploadDynamicItemName.length > 20;

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

    if (acceptedFiles[0].size > 30000) {
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
    formData.append("itemType", itemSelect);
    formData.append("name", uploadDynamicItemName);

    const token = parseCookies().token;
    const uploadConfig = {
      method: "post",
      url: `/softwares/${software_id}/projects/${project_id}/items/upload_dynamic_item`,
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
        setItemSelect("");
        setUploadDynamicItemName("");
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

  const handleChange = (event: SelectChangeEvent) => {
    setItemSelect(event.target.value as string);
  };

  const handleUploadFunction = useAsyncAndLoading(uploadFunction, setState);

  return (
    <>
      <Stack spacing={5}>
        <InputFile
          displayName={"Dynamic Itemの名前"}
          inputText={uploadDynamicItemName}
          setInputText={setUploadDynamicItemName}
          selectedFile={selectedFile}
          onDrop={onDrop}
        >
          <FormControl fullWidth variant="filled">
            <InputLabel id="itemTypeSelect">アイテムの属性</InputLabel>
            <Select
              labelId="itemTypeSelect"
              id="itemTypeSelect"
              value={itemSelect}
              label="ItemType"
              onChange={handleChange}
            >
              <MenuItem value={"画像アイテム"}>画像アイテム</MenuItem>
              <MenuItem value={"ビデオアイテム"}>ビデオアイテム</MenuItem>
              <MenuItem value={"オーディオアイテム"}>
                オーディオアイテム
              </MenuItem>
            </Select>
          </FormControl>
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

export default DynamicItem;
