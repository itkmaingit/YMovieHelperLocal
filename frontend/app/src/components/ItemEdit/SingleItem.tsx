import { Button, Grid } from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import React from "react";

import SingleItemDescription, {
  SingleItemDescriptionProps,
} from "@/components/ItemEdit/SingleItemDescription";
import InputFile from "@/components/utils/InputFile";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

const SingleItem = () => {
  const router = useRouter();
  const { software_id, project_id } = router.query;

  const token = parseCookies().token;

  const [singleItems, setSingleItems] = React.useState<
    SingleItemDescriptionProps[]
  >([]);

  const [selectedFile, setSelectedFile] = React.useState<File | null>(null);
  const context = React.useContext(FunctionHandlersContext);

  if (!context) {
    throw new Error(
      "FunctionHandlersContext is undefined, please verify the Provider"
    );
  }

  const { setCode, setState } = context;

  //1つでも空白の要素や、長さが21文字以上のものがあればボタンを押させない
  const buttonDisabled = !(
    singleItems.length !== 0 &&
    singleItems.every((item: SingleItemDescriptionProps) => {
      if (item.Name === undefined || isEmptyOrWhitespace(item.Name)) {
        return false;
      }
      if (item.Name.length > 20) {
        return false;
      }
      return true;
    })
  );

  const uploadClick = async () => {
    if (singleItems != undefined && singleItems.length != 0) {
      const token = parseCookies().token;
      const uploadConfig = {
        method: "post",
        url: `/softwares/${software_id}/projects/${project_id}/items/create_single_item`,
        data: {
          items: singleItems,
        },
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      };

      await clientAxios(uploadConfig)
        .then(() => {
          setCode({ message: "アイテム作成完了！", status: "success" });
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
          setSingleItems([]);
        });
    }
  };

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

    if (acceptedFiles[0].size > 30000) {
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
      url: `/softwares/${software_id}/projects/${project_id}/items/resolve_single_item`,
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${token}`,
      },
    };

    await clientAxios(uploadConfig)
      .then((response) => {
        setSingleItems(response.data.singleItems);
        setCode({ message: "ファイルの解析成功", status: "success" });
      })
      .catch(async (error) => {
        setCode({
          message:
            error.response.data.userError ||
            "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      });
  };

  const wrapOnDrop = useAsyncAndLoading(onDrop, setState);
  const handleUploadClick = useAsyncAndLoading(uploadClick, setState);

  return (
    <>
      <InputFile selectedFile={selectedFile} onDrop={wrapOnDrop}>
        <Grid container spacing={3} sx={{ transform: "translateX(-24px)" }}>
          {singleItems &&
            singleItems.map((item, index) => (
              <SingleItemDescription
                key={item.ID}
                {...item}
                index={index}
                singleItems={singleItems}
                setSingleItems={setSingleItems}
              ></SingleItemDescription>
            ))}
          <Grid item xs={12}></Grid>
        </Grid>
        <Button
          onClick={handleUploadClick}
          variant="contained"
          color="primary"
          disabled={buttonDisabled}
        >
          アップロード
        </Button>
      </InputFile>
    </>
  );
};

export default SingleItem;
