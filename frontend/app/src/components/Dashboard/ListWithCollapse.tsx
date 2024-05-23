/** @jsxImportSource @emotion/react */

import { ExpandLess, ExpandMore } from "@mui/icons-material";
import AddToPhotosIcon from "@mui/icons-material/AddToPhotos";
import ComputerIcon from "@mui/icons-material/Computer";
import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import {
  Button,
  Collapse,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Stack,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import React, { useContext, useState } from "react";

import EditableText from "../utils/EditableText";

import Projects, { Project } from "@/components/Dashboard/Projects";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import { clientAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

export type ListWithCollapseProps = {
  softwareID: number;
  name: string;
  projects: Project[];
};

export default function ListWithCollapse({
  softwareID,
  name,
  projects,
}: ListWithCollapseProps) {
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

  const [dialogOpen, setDialogOpen] = useState(false);
  const [projectDialogOpen, setProjectDialogOpen] = useState(false);
  const [collapseOpen, setCollapseOpen] = React.useState(false);

  const [createProjectName, setCreateProjectName] = React.useState("");
  const addProjectClick = async () => {
    if (
      isEmptyOrWhitespace(createProjectName) ||
      createProjectName.length > 20
    ) {
      setCreateProjectName("");
      setState("completed");
      setCode({
        message: "文字列は20文字以下で、空白のみのものにしないでください！",
        status: "warning",
      });
      return;
    }

    const postConfig = {
      method: "post",
      url: `/softwares/${softwareID}/projects`,
      data: { name: createProjectName },
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(postConfig)
      .catch(() => {
        setCode({
          message: "サーバーエラー！時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(async () => {
        await mutate();
        setProjectDialogOpen(false);
      });
  };

  const handleInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCreateProjectName(e.target.value);
  };

  const handleOpenCharactersPage = async () => {
    await router.push(`/mypage/${softwareID}/characters`);
  };

  const [isEditing, setIsEditing] = useState(false);
  const [updateSoftwareName, setUpdateSoftwareName] = useState(name);
  const handleBlur = async () => {
    if (
      isEmptyOrWhitespace(updateSoftwareName) ||
      updateSoftwareName.length > 20
    ) {
      setUpdateSoftwareName(name);
      setState("completed");
      setCode({
        message: "文字列は20文字以下で、空白のみのものにしないでください！",
        status: "warning",
      });
      return;
    }

    setIsEditing(false);

    if (updateSoftwareName === name) {
      return;
    }

    const putConfig = {
      method: "put",
      url: `/softwares/${softwareID}`,
      data: { id: softwareID, name: updateSoftwareName },
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(putConfig).catch(() => {
      setCode({
        message: "サーバーエラー！時間をおいてから試してください！",
        status: "warning",
      });
      setUpdateSoftwareName(name);
    });
    await mutate();
  };

  const deleteClick = async () => {
    const deleteConfig = {
      method: "delete",
      url: `/softwares/${softwareID}`,
      headers: { Authorization: `Bearer ${token}` },
    };

    await clientAxios(deleteConfig)
      .catch(() => {
        setCode({
          message: "サーバーエラー！少し時間をおいてから試してください！",
          status: "warning",
        });
      })
      .finally(async () => await mutate());
  };

  const handleDeleteClick = useAsyncAndLoading(deleteClick, setState);
  const handleAddProjectClick = useAsyncAndLoading(addProjectClick, setState);

  return (
    <React.Fragment key={softwareID}>
      <ListItem
        secondaryAction={
          // aria-controlsのとMenuのidが対応する
          <Stack spacing={3} direction="row">
            <Tooltip title="プロジェクトの追加" arrow>
              <ListItemButton onClick={() => setProjectDialogOpen(true)}>
                <AddToPhotosIcon fontSize="large" />
              </ListItemButton>
            </Tooltip>
            <Tooltip title="キャラクター" arrow>
              <ListItemButton onClick={handleOpenCharactersPage}>
                <PeopleAltIcon fontSize="large" />
              </ListItemButton>
            </Tooltip>
            <Tooltip title="ソフトウェアの削除" arrow>
              <ListItemButton onClick={() => setDialogOpen(true)}>
                <DeleteForeverIcon fontSize="large" color="error" />
              </ListItemButton>
            </Tooltip>
          </Stack>
        }
        sx={{ padding: 2, px: 5 }}
      >
        <Stack direction="row" spacing={3} alignItems="center">
          <ListItemButton onClick={() => setCollapseOpen(!collapseOpen)}>
            {collapseOpen ? (
              <ExpandLess fontSize="large" />
            ) : (
              <ExpandMore fontSize="large" />
            )}
          </ListItemButton>
          <ListItemIcon>
            <ComputerIcon fontSize="large" />
          </ListItemIcon>
          <ListItemText
            primary={
              <EditableText
                text={updateSoftwareName}
                setText={setUpdateSoftwareName}
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
        </Stack>
      </ListItem>
      <Collapse
        in={collapseOpen}
        timeout="auto"
        unmountOnExit
        key={`collapse-${softwareID}`}
      >
        <List component="div" disablePadding>
          {projects &&
            projects.map((project: Project) => (
              <Projects
                key={project.projectID}
                softwareID={softwareID}
                project={project}
              ></Projects>
            ))}
        </List>
      </Collapse>

      {/* ----------------プロジェクトpost用のDialog-------------- */}
      <Dialog
        onClose={() => setProjectDialogOpen(false)}
        open={projectDialogOpen}
        maxWidth="sm"
        fullWidth
      >
        <DialogTitle>プロジェクトの作成</DialogTitle>
        <DialogContent>
          <Stack spacing={3}>
            <TextField
              sx={{ mt: 2 }}
              label="プロジェクトの名前"
              fullWidth
              value={createProjectName}
              onChange={handleInput}
            />

            <Button
              type="submit"
              variant="contained"
              color="primary"
              size="large"
              disabled={
                createProjectName.length === 0 || createProjectName.length > 20
              }
              onClick={handleAddProjectClick}
            >
              Upload
            </Button>
          </Stack>
        </DialogContent>
      </Dialog>
      {/* ----------------プロジェクトpost用のDialog-------------- */}

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
          >{`「${name}」の削除`}</Typography>
        </DialogTitle>
        <DialogContent>
          <Typography variant="h6">
            ソフトウェアを削除すると、復元することができません。本当に削除しますか？
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
}
