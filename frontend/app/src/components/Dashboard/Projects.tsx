import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import FolderIcon from "@mui/icons-material/Folder";
import GavelIcon from "@mui/icons-material/Gavel";
import SlideshowIcon from "@mui/icons-material/Slideshow";
import UploadFileIcon from "@mui/icons-material/UploadFile";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Stack,
  Tooltip,
  Typography,
} from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import React, { useContext, useState } from "react";

import EditableText from "../utils/EditableText";

import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

export type Project = {
  projectID: number;
  name: string;
};

type ProjectsProps = {
  softwareID: number;
  project: Project;
};

const Projects = ({ softwareID, project }: ProjectsProps) => {
  const router = useRouter();
  const token = parseCookies().token;
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
  const [updateProjectName, setUpdateProjectName] = useState(project.name);
  const [isEditing, setIsEditing] = useState(false);
  const [dialogOpen, setDialogOpen] = useState(false);

  const handleEditClick = async () => {
    await router.push(`/mypage/${softwareID}/${project.projectID}/items`);
  };

  const handleRuleClick = async () => {
    await router.push(`/mypage/${softwareID}/${project.projectID}/rules`);
  };

  const handleMakeYMMPClick = async () => {
    await router.push(`/mypage/${softwareID}/${project.projectID}/make_ymmp`);
  };

  const handleBlur = async () => {
    if (
      isEmptyOrWhitespace(updateProjectName) ||
      updateProjectName.length > 20
    ) {
      setUpdateProjectName(project.name);
      setState("completed");
      setCode({
        message: "文字列は20文字以下で、空白のみのものにしないでください！",
        status: "warning",
      });
      return;
    }

    setIsEditing(false);

    if (updateProjectName === project.name) {
      return;
    }

    const putConfig = {
      method: "put",
      url: `/softwares/${softwareID}/projects/${project.projectID}`,
      data: { id: project.projectID, name: updateProjectName },
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(putConfig).catch(() => {
      setCode({
        message: "サーバーエラー！時間をおいてから試してください！",
        status: "warning",
      });
      setUpdateProjectName(project.name);
    });
    await mutate();
  };

  const deleteClick = async () => {
    const deleteConfig = {
      method: "delete",
      url: `/softwares/${softwareID}/projects/${project.projectID}`,
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(deleteConfig)
      .catch(() => {
        setCode({
          message: "サーバーエラー！少し時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(async () => {
        await mutate();
        setDialogOpen(false);
      });
  };

  const handleDeleteClick = useAsyncAndLoading(deleteClick, setState);

  return (
    <React.Fragment key={`project-${softwareID}-${project.projectID}`}>
      <Divider></Divider>
      <ListItem
        secondaryAction={
          <Stack spacing={3} direction="row">
            <Tooltip title="アイテム" arrow>
              <ListItemButton onClick={handleEditClick}>
                <UploadFileIcon fontSize="large" />
              </ListItemButton>
            </Tooltip>
            <Tooltip title="ルール" arrow>
              <ListItemButton onClick={handleRuleClick}>
                <GavelIcon fontSize="large" />
              </ListItemButton>
            </Tooltip>
            <Tooltip title="動画の作成" arrow>
              <ListItemButton onClick={handleMakeYMMPClick}>
                <SlideshowIcon fontSize="large" />
              </ListItemButton>
            </Tooltip>
            <Tooltip title="プロジェクトの削除" arrow>
              <ListItemButton onClick={() => setDialogOpen(true)}>
                <DeleteForeverIcon fontSize="large" color="error" />
              </ListItemButton>
            </Tooltip>
          </Stack>
        }
        sx={{ padding: 2, paddingLeft: 25 }}
      >
        <ListItemIcon>
          <FolderIcon fontSize="large" />
        </ListItemIcon>
        <ListItemText
          primary={
            <EditableText
              text={updateProjectName}
              setText={setUpdateProjectName}
              isEditing={isEditing}
              handleTextClick={() => setIsEditing(true)}
              handleBlur={handleBlur}
              textFieldStyle={{
                style: {
                  fontSize: "1.6rem",
                  padding: "2px",
                  paddingLeft: 10,
                },
              }}
              typographyStyle={{
                fontSize: "2.0rem",
              }}
            ></EditableText>
          }
        />
      </ListItem>
      {/* ----------------ソフトウェアDelete用のDialog-------------- */}
      <Dialog
        onClose={() => setDialogOpen(false)}
        open={dialogOpen}
        maxWidth="md"
        fullWidth
      >
        <DialogTitle>
          <Typography
            variant="h4"
            gutterBottom
          >{`「${project.name}」の削除`}</Typography>
        </DialogTitle>
        <DialogContent>
          <Typography variant="h6">
            プロジェクトを削除すると、復元することができません。本当に削除しますか？
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() => setDialogOpen(false)}
            color="inherit"
            size="large"
          >
            キャンセル
          </Button>
          <Button
            onClick={handleDeleteClick}
            color="error"
            size="large"
            variant="contained"
          >
            削除
          </Button>
        </DialogActions>
      </Dialog>
      {/* ----------------ソフトウェアDelete用のDialog-------------- */}
    </React.Fragment>
  );
};

export default Projects;
