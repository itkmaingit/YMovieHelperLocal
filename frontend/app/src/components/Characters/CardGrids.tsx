import { ExpandLess, ExpandMore } from "@mui/icons-material";
import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import FileUploadIcon from "@mui/icons-material/FileUpload";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  Collapse,
  FormControlLabel,
  Grid,
  IconButton,
  Stack,
  Switch,
} from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import React, { useContext, useState } from "react";
import { useDropzone } from "react-dropzone";

import EmotionListItem, {
  Emotion,
} from "@/components/Characters/EmotionListItem";
import EditableText from "@/components/utils/EditableText";
import InputFileAndUpload from "@/components/utils/InputFileAndUpload";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

export type Character = {
  ID: number;
  IsEmpty: boolean;
  Name: string;
  Emotions: Emotion[];
};

type setIsHeadsUpProps = {
  setIsHeadsUp: React.Dispatch<React.SetStateAction<boolean>>;
};

const CardGrids = ({
  ID,
  IsEmpty,
  Name,
  Emotions,
  setIsHeadsUp,
}: Character & setIsHeadsUpProps) => {
  const token = parseCookies().token;

  const router = useRouter();
  const { software_id } = router.query;
  const context = useContext(FunctionHandlersContext);

  if (!context) {
    throw new Error(
      "FunctionHandlersContext is undefined, please verify the Provider"
    );
  }

  const { mutate, setCode, setState } = context;
  if (mutate === undefined) {
    return null;
  }

  const [dialogOpen, setDialogOpen] = useState(false);
  const [collapseOpen, setCollapseOpen] = React.useState(false);

  const [text, setText] = useState<string>(Name);
  const [isEditing, setIsEditing] = useState<boolean>(false);

  const setImage = (name: string) => {
    if (name === "霊夢" || name === "れいむ" || name === "れーむ") {
      return "/reimu.gif";
    } else if (name === "魔理沙" || name === "まりさ") {
      return "/marisa.gif";
    }

    return "/show_unknown_character.svg";
  };
  const handleBlur = async () => {
    if (isEmptyOrWhitespace(text) || text.length > 20) {
      setText(Name);
      setState("completed");
      setCode({
        message: "文字列は20文字以下で、空白のみのものにしないでください！",
        status: "warning",
      });
      return;
    }

    setIsEditing(false);

    if (text === Name) {
      return;
    }
    const data = { id: ID, name: text };
    const putConfig = {
      method: "put",
      url: `/softwares/${software_id}/characters/${ID}`,
      data: data,
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(putConfig).catch(() =>
      setCode({
        message: "サーバーエラー！時間をおいてから試してください！",
        status: "warning",
      })
    );
    await mutate();
  };

  const [checked, setChecked] = React.useState(!IsEmpty);
  const switchChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const data = { id: ID, isEmpty: !event.target.checked };
    const putConfig = {
      method: "put",
      url: `/softwares/${software_id}/characters/${ID}`,
      data: data,
      headers: { Authorization: `Bearer ${token}` },
    };
    await clientAxios(putConfig)
      .then(() => {
        setChecked(!event.target.checked);
        setIsHeadsUp(true);
      })
      .catch(() =>
        setCode({
          message: "サーバーエラー。時間をおいてやり直してください。",
          status: "warning",
        })
      );
  };

  const deleteClick = async () => {
    const deleteConfig = {
      method: "delete",
      url: `/softwares/${software_id}/characters/${ID}`,
      headers: { Authorization: `Bearer ${token}` },
    };
    await clientAxios(deleteConfig)
      .then(async () => {
        setCode({ message: "削除完了", status: "success" });
      })
      .catch(() => {
        setCode({ message: "削除失敗", status: "warning" });
      })
      .finally(async () => await mutate());
  };

  const [selectedFile, setSelectedFile] = React.useState<null | File>(null);

  const fileUpload = async (file: File) => {
    const formData = new FormData();
    formData.append("file", file);
    const uploadConfig = {
      method: "post",
      url: `/softwares/${software_id}/characters/${ID}/emotions`,
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${token}`,
      },
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
        await mutate();
        setSelectedFile(null);
      });
  };

  const onDrop = async (acceptedFiles: File[]) => {
    const extension = acceptedFiles[0].name.split(".").pop();
    if (extension !== "ymmp") {
      setDialogOpen(false);
      setState("completed");
      setCode({
        message:
          "アップロード失敗！ファイル形式はymmpのものしか受け付けられません！",
        status: "warning",
      });
      return;
    }

    if (acceptedFiles[0].size > 300000) {
      setDialogOpen(false);
      setState("completed");
      setCode({
        message:
          "アップロード失敗！ファイルサイズが大きすぎます！(300KB以下にしてください！)",
        status: "warning",
      });
      return;
    }

    setSelectedFile(acceptedFiles[0]);
    if (acceptedFiles[0] != undefined || null) {
      await wrapFileUpload(acceptedFiles[0]);
    }
    setDialogOpen(false);
  };

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
  });

  const handleDeleteClick = useAsyncAndLoading(deleteClick, setState);
  const handleSwitchChange = useAsyncAndLoading(switchChange, setState);
  const wrapFileUpload = useAsyncAndLoading(fileUpload, setState);

  return (
    <React.Fragment key={ID}>
      <Grid item xs={12} xl={6} my={3}>
        <Card
          sx={{
            maxWidth: 345,
            boxShadow: "1px 1px 10px rgba(0,0,0,0.15)",
            transition: "0.3s",
            borderRadius: "15px",
          }}
        >
          <CardMedia
            component="img"
            height="250"
            image={setImage(Name)}
            alt={Name}
            sx={{
              objectPosition: "center 30px",
            }}
          />
          <CardContent>
            <Stack direction="row" justifyContent="space-around">
              <EditableText
                text={text}
                setText={setText}
                isEditing={isEditing}
                handleBlur={handleBlur}
                handleTextClick={() => setIsEditing(true)}
                textFieldStyle={{
                  style: {
                    fontSize: "1.6rem",
                    padding: "2px",
                    paddingLeft: 10,
                  },
                }}
                typographyStyle={{ fontSize: "2.0rem" }}
              ></EditableText>
              <FormControlLabel
                control={
                  <Switch checked={checked} onChange={handleSwitchChange} />
                }
                labelPlacement="bottom"
                label={checked ? "キャラ" : "削除"}
              />
            </Stack>
          </CardContent>
          <CardActions>
            <Stack sx={{ width: "100%" }}>
              <Stack direction="row" justifyContent="space-around">
                <IconButton onClick={() => setCollapseOpen(!collapseOpen)}>
                  {collapseOpen ? (
                    <ExpandLess fontSize="large" />
                  ) : (
                    <ExpandMore fontSize="large" />
                  )}
                </IconButton>
                <Button
                  size="large"
                  color="primary"
                  startIcon={<FileUploadIcon />}
                  onClick={() => setDialogOpen(true)}
                >
                  Add Emotion
                </Button>
                <Button
                  size="large"
                  color="error"
                  startIcon={<DeleteForeverIcon />}
                  onClick={handleDeleteClick}
                >
                  Delete
                </Button>
              </Stack>
            </Stack>
          </CardActions>
          <Collapse in={collapseOpen} timeout="auto" unmountOnExit>
            {Emotions &&
              Emotions.map((emotion) => {
                return (
                  <EmotionListItem
                    key={`emotion-${emotion.ID}`}
                    emotion={emotion}
                    characterID={ID}
                  ></EmotionListItem>
                );
              })}
          </Collapse>
        </Card>
      </Grid>

      {/* ------------------Dialog---------------- */}

      <InputFileAndUpload
        dialogOpen={dialogOpen}
        setDialogOpen={setDialogOpen}
        selectedFile={selectedFile}
        getRootProps={getRootProps}
        getInputProps={getInputProps}
        isDragActive={isDragActive}
        title={`${Name}の表情追加`}
      ></InputFileAndUpload>
      {/* ------------------Dialog---------------- */}
    </React.Fragment>
  );
};

export default CardGrids;
