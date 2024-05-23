import PersonIcon from "@mui/icons-material/Person";
import { Button, Grid } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Link from "@mui/material/Link";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import {
  createUserWithEmailAndPassword,
  deleteUser,
  sendEmailVerification,
} from "firebase/auth";
import * as React from "react";

import { NextPageWithLayout } from "../_app";

import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import useFeedback from "@/hooks/use_feedback";
import NotDashboardLayout from "@/layouts/NotDashboardLayout";
import { clientAxios } from "@/libs/axios";
import { auth } from "@/libs/firebase/firebase";

const SignUp: NextPageWithLayout = () => {
  const [name, setName] = React.useState<string>("");
  const [email, setEmail] = React.useState<string>("");
  const [password, setPassword] = React.useState<string>("");
  const submitSignin = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    await createUserWithEmailAndPassword(auth, email, password)
      .then(async (userCredential) => {
        const postConfig = {
          method: "post",
          url: "/user",
          data: {
            uid: userCredential.user.uid,
            name: name,
          },
        };
        await clientAxios(postConfig)
          .then(async () => {
            if (!process.env.NEXT_PUBLIC_FIREBASE_ACTION_URL) {
              setCode({
                message:
                  "サーバーエラー！早急にTwitterのDMにまで報告していただけると助かります！",
                status: "warning",
              });
              return;
            }
            const actionCodeSettings = {
              url: `${process.env.NEXT_PUBLIC_FIREBASE_ACTION_URL}/login`,
              handleCodeInApp: false,
            };
            await sendEmailVerification(
              userCredential.user,
              actionCodeSettings
            ).then(async () => {
              alert(
                "認証メールを送信しました。メールに記載されているURLをクリックして、ユーザー登録を完了させてください。"
              );
            });
          })
          .catch(async () => {
            setState("completed");
            setCode({
              message: "サーバーエラー！時間をおいてから試してください！",
              status: "warning",
            });

            await deleteUser(userCredential.user);
          });
      })
      .catch((error) => {
        const errorCode: string = error.code;
        if (errorCode == "auth/email-already-in-use") {
          setCode({
            message: "既に使用されているメールアドレスです！",
            status: "warning",
          });
        } else if (errorCode == "auth/invalid-email") {
          setCode({ message: "不正なメールアドレスです！", status: "warning" });
        } else if (errorCode == "auth/weak-password") {
          setCode({
            message: "もう少し堅牢なパスワードを使用してください！(5文字以上)",
            status: "warning",
          });
        } else {
          setCode({
            message: "原因不明のエラーです!TwitterのDMまでお知らせください!",
            status: "warning",
          });
        }
      });
  };

  const [setCode, setState, snackbarProps] = useFeedback();
  const handleSubmitSignin = useAsyncAndLoading(submitSignin, setState);

  return (
    <>
      <LoadingAndSnackbar {...snackbarProps} />
      <NotDashboardLayout>
        <Container component="main" maxWidth="xs">
          <Box
            sx={{
              marginTop: 8,
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
            }}
          >
            <Avatar sx={{ m: 3, bgcolor: "primary.main" }}>
              <PersonIcon />
            </Avatar>
            <Typography component="h1" variant="h5">
              アカウント作成
            </Typography>
            <Box component="form" onSubmit={handleSubmitSignin} sx={{ mt: 1 }}>
              <TextField
                margin="normal"
                required
                fullWidth
                id="name"
                label="Your Name"
                name="name"
                autoFocus
                value={name}
                onChange={(e) => setName(e.target.value)}
                error={name.length > 20}
                helperText={
                  name.length > 20 && "名前は20文字以内でお願いします！"
                }
              />
              <TextField
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
              <TextField
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                disabled={!name || !email || !password || name.length > 20}
              >
                送信
              </Button>
              <Grid container>
                <Grid item xs>
                  <Link href="/" variant="body2">
                    トップに戻る
                  </Link>
                </Grid>
                <Grid item>
                  <Link href="/login" variant="body2">
                    既にアカウントを持っている場合はこちら
                  </Link>
                </Grid>
              </Grid>
            </Box>
          </Box>
        </Container>
      </NotDashboardLayout>
    </>
  );
};

SignUp.getLayout = function getLayout(page: React.ReactElement) {
  return <NotDashboardLayout>{page}</NotDashboardLayout>;
};
export default SignUp;
