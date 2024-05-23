import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import { IconButton, TableCell, TableRow, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import { useContext, useState } from "react";

import EditableText from "../utils/EditableText";

import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

export type MultipleItem = {
  id: number;
  name: string;
  countOfItems: number;
};

const MultipleItemRecord = ({ id, name, countOfItems }: MultipleItem) => {
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
  const token = parseCookies().token;

  const router = useRouter();

  if (!router.isReady) return null;

  const { software_id, project_id } = router.query;

  const [updateName, setUpdateName] = useState(name);
  const [isEditing, setIsEditing] = useState(false);
  const handleUpdateNameBlur = async () => {
    setIsEditing(false);
    if (isEmptyOrWhitespace(updateName) || updateName.length > 20) {
      setUpdateName(name);
      setState("completed");
      setCode({
        message: "文字列は20文字以下で、空白のみのものにしないでください！",
        status: "warning",
      });
      return;
    }

    if (updateName === name) {
      return;
    }

    const putConfig = {
      method: "put",
      url: `/softwares/${software_id}/projects/${project_id}/items/multiple_item`,
      data: { id: id, name: updateName },
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(putConfig).catch(() => {
      setState("completed");
      setCode({
        message: "サーバーエラー！時間をおいてから試してください！",
        status: "warning",
      });
      setUpdateName(name);
    });
    await mutate();
  };

  const deleteClick = async () => {
    const deleteConfig = {
      method: "delete",
      url: `/softwares/${software_id}/projects/${project_id}/items/multiple_item`,
      data: { id: id },
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(deleteConfig)
      .catch(() => {
        setState("completed");
        setCode({
          message: "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(async () => {
        await mutate();
      });
  };

  const handleDeleteClick = useAsyncAndLoading(deleteClick, setState);
  return (
    <TableRow>
      <TableCell component="th" scope="row" align="center">
        <EditableText
          text={updateName}
          setText={setUpdateName}
          isEditing={isEditing}
          handleTextClick={() => setIsEditing(true)}
          handleBlur={handleUpdateNameBlur}
          textFieldStyle={{
            style: {
              fontSize: "1.0rem",
              padding: "2px",
              paddingLeft: 10,
            },
          }}
          typographyStyle={{
            fontSize: "1.2rem",
          }}
        ></EditableText>
      </TableCell>
      <TableCell align="center">
        <Typography>{countOfItems}</Typography>
      </TableCell>
      <TableCell align="center"></TableCell>
      <TableCell align="center">
        <IconButton onClick={handleDeleteClick}>
          <DeleteForeverIcon fontSize="large" color="error" />
        </IconButton>
      </TableCell>
    </TableRow>
  );
};

export default MultipleItemRecord;
