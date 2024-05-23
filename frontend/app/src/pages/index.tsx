import {
  Box,
  Button,
  Divider,
  Grid,
  Link,
  Stack,
  Typography,
  useMediaQuery,
} from "@mui/material";
import Image from "next/image";
import { useRouter } from "next/router";
import { ReactElement } from "react";

import Item from "@/components/IndexPage/Item";
import handlePageTransition from "@/feature/page_transition";
import Footer from "@/layouts/Footer";
import NotDashboardLayout from "@/layouts/NotDashboardLayout";
import theme from "@/theme";

import { NextPageWithLayout } from "./_app";

const Home: NextPageWithLayout = () => {
  const router = useRouter();
  const matches = useMediaQuery(theme.breakpoints.down("md"));
  return matches ? (
    <>
      <Box sx={{ backgroundColor: "white", padding: "8 0 6" }}>
        <Box sx={{ width: "100%", height: "500px", position: "relative" }}>
          <Image
            src="/demo.gif"
            style={{ objectFit: "cover" }}
            alt="demo.gif"
            fill
          ></Image>
          <Typography
            component="h1"
            variant="h2"
            align="center"
            color="textPrimary"
            gutterBottom
            sx={{
              position: "absolute",
              top: "25%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              color: "white",
              fontSize: "3em",
              textShadow: "2px 2px 4px rgba(0, 0, 0, 0.5)",
            }}
          >
            YMovieHelper
          </Typography>
          <Typography
            component="h2"
            align="center"
            color="textPrimary"
            gutterBottom
            sx={{
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              color: "white",
              fontSize: "1.3em",
              textShadow: "2px 2px 4px rgba(0, 0, 0, 0.5)",
            }}
          >
            動画編集者も<br></br>ゆっくりできる<br></br>時代に。
          </Typography>
          <Stack
            sx={{
              position: "absolute",
              top: "80%",
              left: "50%",
              transform: "translate(-50%, -50%)",
            }}
            direction="row"
            spacing={10}
          >
            <Button
              size="large"
              variant="contained"
              onClick={handlePageTransition(router, "/how_to")}
            >
              <Typography fontSize="1.0rem">How to</Typography>
            </Button>
            <Button
              size="large"
              variant="contained"
              onClick={handlePageTransition(router, "/mypage/dashboard")}
            >
              <Typography fontSize="1.0rem">My Page</Typography>
            </Button>
          </Stack>
        </Box>
        <Grid container spacing={0}>
          <Grid item xs={12} sx={{ height: "300px" }}>
            <Box
              sx={{
                display: "flex",
                height: "100%", // 100% of the viewport height
                justifyContent: "center",
                alignItems: "center",
                backgroundColor: "primary.main",
              }}
            >
              <Typography variant="h4" color="white">
                YMovieHelperとは
              </Typography>
            </Box>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "700px" }}>
            <Box
              sx={{
                display: "flex",
                height: "100%",
                justifyContent: "center",
                alignItems: "center",
                px: 5,
                flexDirection: "column",
              }}
            >
              <Box
                sx={{ width: "100%", height: "400px", position: "relative" }}
              >
                <Image
                  src="/flow_chart.gif"
                  style={{ objectFit: "contain" }}
                  alt="flow_chart"
                  fill
                ></Image>
              </Box>
              <Typography sx={{ fontSize: "clamp(1.2rem, 3vw, 3rem)" }}>
                指定したCSV形式の台本をアップロードするだけで、作成したい動画のymmpファイルを作成できるWebアプリケーションです。
              </Typography>
            </Box>
          </Grid>
          <Divider></Divider>
          <Grid item xs={12} sx={{ height: "100px", mt: "100px" }}>
            <Stack alignItems="center">
              <Typography
                variant="h3"
                sx={{
                  textDecoration: "underline solid #4682b4",
                  textUnderlinePosition: "under",
                }}
              >
                特徴
              </Typography>
            </Stack>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "400px", my: "50px", mx: 0 }}>
            <Item
              src="/item.svg"
              alt="アイテムの画像"
              title="アイテム"
              description="YMM4の全てのアイテムに対応しています。"
            ></Item>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "400px", my: "50px" }}>
            <Item
              src="/rule.svg"
              alt="ルールの画像"
              title="ルール"
              description="動画の構成を自由に決めることができます。"
            ></Item>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "400px", my: "50px" }}>
            <Item
              src="/character.svg"
              alt="キャラクターの画像"
              title="ボイスエンジン"
              description="全てのボイスエンジンに対応しています。上手く音声が生成されないなどの不具合を確認した際には、お問い合わせまでご連絡いただけると助かります。"
            ></Item>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "400px", my: "50px" }}>
            <Item
              src="/money.svg"
              alt="お金の画像"
              title="完全無料"
              description="アカウント登録をするだけでYMovieHelperは無料で使えます。"
            ></Item>
          </Grid>
          <Divider></Divider>
          <Grid item xs={12} sx={{ height: "100px", mt: "100px" }}>
            <Stack alignItems="center">
              <Typography
                sx={{
                  textDecoration: "underline solid #4682b4",
                  textUnderlinePosition: "under",
                  fontSize: "clamp(2.3rem, 5vw, 3rem)",
                }}
              >
                サポート
              </Typography>
            </Stack>
          </Grid>
          <Grid item xs={12} sx={{ mt: " 100px ", mx: 5 }}>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                textAlign: "center",
              }}
            >
              <Typography
                sx={{
                  fontSize: "clamp(1.3rem, 3vw, 2rem)",
                }}
              >
                何かわからないことがあれば、気軽に
                <Link
                  sx={{ cursor: "pointer" }}
                  href="https://forms.gle/Gh3ZpW9DS64eb1qS7"
                  target="_blank"
                >
                  お問い合わせフォーム
                </Link>
                でご連絡ください。
                些細なことでもご連絡いただければ対応いたします。
              </Typography>
            </Box>
          </Grid>
          <Footer />
        </Grid>
      </Box>
    </>
  ) : (
    <>
      <Box sx={{ backgroundColor: "white", padding: "8 0 6" }}>
        <Box sx={{ width: "100%", height: "500px", position: "relative" }}>
          <Image
            src="/demo.gif"
            style={{ objectFit: "cover" }}
            alt="demo.gif"
            fill
          ></Image>
          <Typography
            component="h1"
            variant="h2"
            align="center"
            color="textPrimary"
            gutterBottom
            sx={{
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              color: "white",
              fontSize: "3em",
              textShadow: "2px 2px 4px rgba(0, 0, 0, 0.5)",
            }}
          >
            YMovieHelper
          </Typography>
          <Typography
            component="h2"
            variant="h3"
            align="center"
            color="textPrimary"
            gutterBottom
            sx={{
              position: "absolute",
              top: "60%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              color: "white",
              fontSize: "1.5em",
              textShadow: "2px 2px 4px rgba(0, 0, 0, 0.5)",
            }}
          >
            動画編集者もゆっくりできる時代に。
          </Typography>
          <Stack
            sx={{
              position: "absolute",
              top: "75%",
              left: "50%",
              transform: "translate(-50%, -50%)",
            }}
            direction="row"
            justifyContent="space-between"
            alignItems="space-between"
            spacing={3}
          >
            <Button
              size="large"
              variant="contained"
              onClick={handlePageTransition(router, "/how_to")}
            >
              <Typography fontSize="1.3rem">How to</Typography>
            </Button>
            <Button
              size="large"
              variant="contained"
              onClick={handlePageTransition(router, "/mypage/dashboard")}
            >
              <Typography fontSize="1.3rem">My Page</Typography>
            </Button>
          </Stack>
        </Box>
        <Grid container spacing={0}>
          <Grid item xs={12} lg={6} sx={{ height: "700px" }}>
            <Box
              sx={{
                display: "flex",
                height: "100%", // 100% of the viewport height
                justifyContent: "center",
                alignItems: "center",
                backgroundColor: "primary.main",
              }}
            >
              <Typography variant="h3" color="white">
                YMovieHelperとは
              </Typography>
            </Box>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "700px" }}>
            <Box
              sx={{
                display: "flex",
                height: "100%",
                justifyContent: "center",
                alignItems: "center",
                px: 20,
                flexDirection: "column",
              }}
            >
              <Box
                sx={{ width: "100%", height: "400px", position: "relative" }}
              >
                <Image
                  src="/flow_chart.gif"
                  style={{ objectFit: "contain" }}
                  alt="flow_chart"
                  fill
                ></Image>
              </Box>
              <Typography variant="h5">
                指定したCSV形式の台本をアップロードするだけで、作成したい動画のymmpファイルを作成できるWebアプリケーションです。
              </Typography>
            </Box>
          </Grid>
          <Divider></Divider>
          <Grid item xs={12} sx={{ height: "100px", mt: "100px" }}>
            <Stack alignItems="center">
              <Typography
                variant="h2"
                sx={{
                  textDecoration: "underline solid #4682b4",
                  textUnderlinePosition: "under",
                  fontSize: "clamp(2.3rem, 5vw, 3rem)",
                }}
              >
                特徴
              </Typography>
            </Stack>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "200px", my: "50px" }}>
            <Item
              src="/item.svg"
              alt="アイテムの画像"
              title="アイテム"
              description="YMM4の全てのアイテムに対応しています。"
            ></Item>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "200px", my: "50px" }}>
            <Item
              src="/rule.svg"
              alt="ルールの画像"
              title="ルール"
              description="動画の構成を自由に決めることができます。"
            ></Item>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "200px", my: "50px" }}>
            <Item
              src="/character.svg"
              alt="キャラクターの画像"
              title="ボイスエンジン"
              description="全てのボイスエンジンに対応しています。上手く音声が生成されないなどの不具合を確認した際には、お問い合わせまでご連絡いただけると助かります。"
            ></Item>
          </Grid>
          <Grid item xs={12} lg={6} sx={{ height: "200px", my: "50px" }}>
            <Item
              src="/money.svg"
              alt="お金の画像"
              title="完全無料"
              description="アカウント登録をするだけでYMovieHelperは無料で使えます。"
            ></Item>
          </Grid>
          <Divider></Divider>
          <Grid item xs={12} sx={{ height: "100px", mt: "100px" }}>
            <Stack alignItems="center">
              <Typography
                variant="h2"
                sx={{
                  textDecoration: "underline solid #4682b4",
                  textUnderlinePosition: "under",
                }}
              >
                サポート
              </Typography>
            </Stack>
          </Grid>
          <Grid item xs={12} sx={{ mt: " 100px " }}>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                textAlign: "center",
              }}
            >
              <Typography fontSize="1.6rem">
                何かわからないことがあれば、気軽に
                <Link
                  sx={{ cursor: "pointer" }}
                  href="https://forms.gle/Gh3ZpW9DS64eb1qS7"
                  target="_blank"
                >
                  お問い合わせフォーム
                </Link>
                でご連絡ください。
                些細なことでもご連絡いただければ対応いたします。
              </Typography>
            </Box>
          </Grid>
          <Footer />
        </Grid>
      </Box>
    </>
  );
};

Home.getLayout = function getLayout(page: ReactElement) {
  return <NotDashboardLayout>{page}</NotDashboardLayout>;
};

export default Home;
