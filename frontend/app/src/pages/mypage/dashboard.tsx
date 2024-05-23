/** @jsxImportSource @emotion/react */

import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  List,
  Paper,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { parseCookies } from "nookies";
import React from "react";
import useSWR from "swr";

import ListWithCollapse, {
  ListWithCollapseProps,
} from "@/components/Dashboard/ListWithCollapse";
import FullScreenLoading from "@/components/utils/FullScreenLoading";
import GlobalAddButton from "@/components/utils/GlobalAddButton";
import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import SectionTitle from "@/components/utils/SectionTitle";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import useFeedback from "@/hooks/use_feedback";
import { clientAxios } from "@/libs/axios";

const Dashboard = () => {
  const fetcher = async (url: string) => {
    const token = parseCookies().token;
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
  const { data, mutate } = useSWR("/softwares", fetcher);

  const [setCode, setState, snackbarProps] = useFeedback();

  // ダイアログの開閉にかかわるstate
  const [dialogOpen, setDialogOpen] = React.useState(false);

  // ソフトウェアの名前入力にかかわるstate
  const [inputValue, setInputValue] = React.useState("");
  const handleInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const createSoftware = async () => {
    const token = parseCookies().token;

    const postConfig = {
      method: "post",
      url: "/softwares",
      data: {
        softwareName: inputValue,
      },
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
        setInputValue("");
        setDialogOpen(false);
      });
  };

  const handleCreateSoftware = useAsyncAndLoading(createSoftware, setState);

  return (
    <>
      <LoadingAndSnackbar {...snackbarProps} />
      <GlobalAddButton
        clickFunction={() => setDialogOpen(true)}
      ></GlobalAddButton>
      <SectionTitle title="Dashboard"></SectionTitle>
      {/* ------------Main--------------- */}
      <FunctionHandlersContext.Provider
        value={{ mutate: mutate, setCode: setCode, setState: setState }}
      >
        {data && data.softwares && data.softwares.length > 0 ? (
          <Paper>
            <List>
              {data.softwares.map(
                (
                  software: ListWithCollapseProps,
                  index: number,
                  arr: ListWithCollapseProps[]
                ) => (
                  <React.Fragment key={index}>
                    <ListWithCollapse {...software}></ListWithCollapse>
                    {index < arr.length - 1 && <Divider />}
                  </React.Fragment>
                )
              )}
            </List>
          </Paper>
        ) : data === undefined ? (
          <FullScreenLoading open={true}></FullScreenLoading>
        ) : (
          <>
            <Stack mt={5}>
              <Typography fontSize="1.2rem" component="div">
                ソフトウェアがまだ作成されていないようです。
              </Typography>
              <Typography fontSize="1.2rem" component="div">
                画面右下の+ボタンからソフトウェアを作成してください。
              </Typography>
            </Stack>
          </>
        )}
      </FunctionHandlersContext.Provider>
      {/* ------------Main--------------- */}
      {/* ------------------Dialog---------------- */}
      <Dialog
        onClose={() => setDialogOpen(false)}
        open={dialogOpen}
        maxWidth="sm"
        fullWidth
      >
        <DialogTitle>ソフトウェアの作成</DialogTitle>
        <DialogContent>
          <TextField
            sx={{ mt: 2 }}
            label="ソフトウェアの名前"
            fullWidth
            value={inputValue}
            onChange={handleInput}
          />
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
            type="submit"
            variant="contained"
            color="primary"
            size="large"
            disabled={inputValue.length === 0 || inputValue.length > 20}
            onClick={handleCreateSoftware}
          >
            作成
          </Button>
        </DialogActions>
      </Dialog>
      {/* ------------------Dialog---------------- */}
    </>
  );
};

export default Dashboard;
