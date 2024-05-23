import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import TagFacesIcon from "@mui/icons-material/TagFaces";
import {
  Divider,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
} from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import { useContext, useState } from "react";

import EditableText from "@/components/utils/EditableText";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

export type Emotion = {
  ID: number;
  Name: string;
};

type EmotionListItemProps = {
  emotion: Emotion;
  characterID: number;
};

const EmotionListItem = ({ emotion, characterID }: EmotionListItemProps) => {
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

  const [text, setText] = useState(emotion.Name);
  const [isEditing, setIsEditing] = useState(false);

  const handleBlur = async () => {
    if (isEmptyOrWhitespace(text) || text.length > 20) {
      setText(emotion.Name);
      setState("completed");
      setCode({
        message: "文字列は20文字以下で、空白のみのものにしないでください！",
        status: "warning",
      });
      return;
    }
    setIsEditing(false);
    const data = { id: emotion.ID, name: text };
    const putConfig = {
      method: "put",
      url: `/softwares/${software_id}/characters/${characterID}/emotions`,
      data: data,
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(putConfig)
      .then(() => setCode({ message: "名前変更完了", status: "success" }))
      .catch(() =>
        setCode({
          message: "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        })
      );
    await mutate();
  };

  const deleteClick = async () => {
    const deleteConfig = {
      method: "delete",
      url: `/softwares/${software_id}/characters/${characterID}/emotions`,
      headers: { Authorization: `Bearer ${token}` },
      data: { id: emotion.ID },
    };

    await clientAxios(deleteConfig)
      .catch(() => {
        setCode({
          message: "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(async () => await mutate());
  };

  const handleDeleteClick = useAsyncAndLoading(deleteClick, setState);
  return (
    <>
      <Divider></Divider>
      <ListItem
        secondaryAction={
          <ListItemButton onClick={handleDeleteClick}>
            <DeleteForeverIcon color="error" />
          </ListItemButton>
        }
      >
        <ListItemIcon>
          <TagFacesIcon></TagFacesIcon>
        </ListItemIcon>
        <ListItemText
          primary={
            <EditableText
              text={text}
              setText={setText}
              isEditing={isEditing}
              handleTextClick={() => setIsEditing(true)}
              handleBlur={handleBlur}
              textFieldStyle={{
                style: {
                  fontSize: "1.2rem",
                  padding: "2px",
                  paddingLeft: 10,
                },
              }}
              typographyStyle={{ fontSize: "1.2rem" }}
            />
          }
        >
          <Typography>{emotion.Name}</Typography>
        </ListItemText>
      </ListItem>
    </>
  );
};

export default EmotionListItem;
