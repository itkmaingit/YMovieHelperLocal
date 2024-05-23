import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import { Stack } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import Link from "@mui/material/Link";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import {
  sendEmailVerification,
  signInWithEmailAndPassword,
} from "firebase/auth";
import { useRouter } from "next/router";
import { setCookie } from "nookies";
import * as React from "react";

import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import handlePageTransition from "@/feature/page_transition";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import useFeedback from "@/hooks/use_feedback";
import NotDashboardLayout from "@/layouts/NotDashboardLayout";
import { clientAxios } from "@/libs/axios";
import { auth } from "@/libs/firebase/firebase";
import { NextPageWithLayout } from "@/pages/_app";

const Login: NextPageWithLayout = () => {
  const [setCode, setState, snackbarProps] = useFeedback();
  const router = useRouter();
  const [email, setEmail] = React.useState<string>("");
  const [password, setPassword] = React.useState<string>("");
  // 1. When the login information is correct and the email address is verified, log in
  // 2. When the login information is correct but the email address is not verified, send a verification email again
  // 3. When the login information is incorrect, display an error

  // Diagram
  // +--------------------------------------+
  // | Check login information              |
  // +--------------------------------------+
  //         |                 |
  //   correct        incorrect
  //         |                 |
  //         v                 v
  // +-------------+    +-----------------------+
  // | Email       |    | Display error         |
  // | verification|    +-----------------------+
  // +-------------+
  //         |
  // verified+--------------------+ not verified
  //         |                    |
  //         v                    v
  // +--------------------+  +------------------+
  // | Log in and redirect|  | Send verification|
  // | dashboard page     |  | email again      |
  // +--------------------+  +------------------+

  const loginAccount = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const requestData = {
      email: email,
      password: password,
    };

    await signInWithEmailAndPassword(
      auth,
      requestData.email,
      requestData.password
    )
      .then(async (userCredential) => {
        if (userCredential.user) {
          if (userCredential.user.emailVerified) {
            const config = {
              method: "post",
              url: "/login",
              data: {
                uid: userCredential.user.uid,
              },
            };
            await clientAxios(config)
              .then(async (res) => {
                setCookie(null, "token", res.data.token, {
                  path: "/",
                  maxAge: 60 * 60 * 24,
                  secure: true,
                  httpOnly: false,
                  sameSite: "strict",
                });
                await router.push("/mypage/dashboard");
              })
              .catch(() => {
                setCode({
                  message: "サーバーエラー！時間をおいてから試してください！",
                  status: "warning",
                });
              });
          } else {
            setCode({
              message:
                "メールアドレスが認証されていません。認証メールを確認してください。",
              status: "warning",
            });
            if (auth.currentUser != null) {
              await sendEmailVerification(auth.currentUser).then(() => {
                alert(
                  "認証メールを再度送信しました。メールに記載されているURLをクリックして、ユーザー登録を完了させてください。"
                );
              });
            }
          }
        }
      })
      .catch((error) => {
        const errorCode: string = error.code;
        if (
          errorCode == "auth/wrong-password" ||
          errorCode == "auth/user-not-found" ||
          errorCode == "auth/invalid-email"
        ) {
          setCode({
            message: "メールアドレス、あるいはパスワードが間違っています。",
            status: "warning",
          });
        } else {
          setCode({
            message: "原因不明のエラーです! TwitterのDMまでお知らせください！",
            status: "warning",
          });
        }
      });
  };

  const handleLoginAccount = useAsyncAndLoading(loginAccount, setState);

  return (
    <>
      <NotDashboardLayout>
        <LoadingAndSnackbar {...snackbarProps}></LoadingAndSnackbar>
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
              <LockOutlinedIcon />
            </Avatar>
            <Typography component="h1" variant="h5">
              ログイン
            </Typography>
            <Box component="form" onSubmit={handleLoginAccount} sx={{ mt: 1 }}>
              <TextField
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                autoFocus
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
                disabled={!email || !password}
              >
                ログイン
              </Button>
              <Grid container>
                <Grid item xs>
                  <Link href="/" variant="body2">
                    トップに戻る
                  </Link>
                </Grid>
                <Grid item xs>
                  <Link href="/login/password_forgatten" variant="body2">
                    パスワードを忘れてしまった場合はこちら
                  </Link>
                </Grid>
              </Grid>
              <Stack alignItems="center" spacing={3} mt={8}>
                <Grid item xs>
                  <Link href="/mypage/dashboard" variant="body2">
                    自動でページが遷移しない時はここをクリックしてください。
                  </Link>
                </Grid>
                <Typography
                  variant="body1"
                  sx={{
                    background: "linear-gradient(transparent 60%, #fff3c8 60%)",
                  }}
                >
                  アカウントをお持ちでない方はこちら
                </Typography>
                <Button
                  variant="contained"
                  color="secondary"
                  sx={{ width: "100%", color: "white", fontSize: "1.0rem" }}
                  onClick={handlePageTransition(router, "/login/signup")}
                >
                  新規登録
                </Button>
              </Stack>
            </Box>
          </Box>
        </Container>
      </NotDashboardLayout>
    </>
  );
};

Login.getLayout = function getLayout(page: React.ReactElement) {
  return <>{page}</>;
};

export default Login;
