import EmailIcon from "@mui/icons-material/Email";
import { Grid } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Link from "@mui/material/Link";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { sendPasswordResetEmail } from "firebase/auth";
import * as React from "react";

import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import useFeedback from "@/hooks/use_feedback";
import NotDashboardLayout from "@/layouts/NotDashboardLayout";
import { auth } from "@/libs/firebase/firebase";
import { NextPageWithLayout } from "@/pages/_app";

const Login: NextPageWithLayout = () => {
  const [setCode, setState, snackbarProps] = useFeedback();
  const [email, setEmail] = React.useState<string>("");

  const resetPassword = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

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

    await sendPasswordResetEmail(auth, email, actionCodeSettings)
      .then(async () => {
        setCode({
          message:
            "パスワードの再設定メールを送信しました。メールを確認してください。",
          status: "success",
        });
        return;
      })
      .catch((error) => {
        const errorCode: string = error.code;
        if (errorCode == "auth/invalid-email") {
          setCode({
            message: "メールアドレス、あるいはパスワードが間違っています。",
            status: "warning",
          });
        } else if (errorCode == "auth/user-not-found") {
          setCode({
            message: "メールアドレスが存在しません！",
            status: "warning",
          });
        } else {
          setCode({
            message: "原因不明のエラーです! TwitterのDMまでお知らせください！",
            status: "warning",
          });
        }
      })
      .finally(() => setEmail(""));
  };

  const handleResetPassword = useAsyncAndLoading(resetPassword, setState);

  return (
    <>
      <NotDashboardLayout>
        <LoadingAndSnackbar {...snackbarProps} />
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
              <EmailIcon />
            </Avatar>
            <Typography component="h1" variant="h5">
              パスワードの再設定
            </Typography>
            <Box
              component="form"
              onSubmit={handleResetPassword}
              noValidate
              sx={{ mt: 1, width: "100%" }}
            >
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

              <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                disabled={!email}
              >
                送信
              </Button>
              <Grid container>
                <Grid item xs>
                  <Link href="/" variant="body2">
                    トップに戻る
                  </Link>
                </Grid>
                <Grid item xs>
                  <Link href="/login" variant="body2">
                    ログイン画面へ戻る
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

Login.getLayout = function getLayout(page: React.ReactElement) {
  return <>{page}</>;
};

export default Login;
