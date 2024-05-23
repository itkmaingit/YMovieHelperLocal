/** @jsxImportSource @emotion/react */
import { Button, Stack, SxProps, Typography } from "@mui/material";
import { useRouter } from "next/router";

import handlePageTransition from "@/feature/page_transition";

const Footer = () => {
  const router = useRouter();

  const linkStyle: SxProps = {
    display: "inlineBlock",
    color: "white",
    fontSize: "1.3rem",
    textDecoration: "none",
    position: "relative",
    transition: ".3s",
    "&::after": {
      position: "absolute",
      bottom: 0,
      left: "50%",
      content: '""',
      width: 0,
      height: "1px",
      backgroundColor: "white",
      transition: "0.3s",
      transform: "translateX(-50%)",
    },
    "&:hover::after": {
      width: "100%",
    },
  };
  return (
    <Stack
      sx={{
        height: 500,
        backgroundColor: "black",
        marginTop: "300px",
        marginBottom: 0,
        minWidth: "100vw",
      }}
      alignItems="center"
      justifyContent="center"
    >
      <Button onClick={handlePageTransition(router, "/")}>
        <Typography sx={linkStyle} gutterBottom>
          ホーム
        </Typography>
      </Button>
      <Button onClick={handlePageTransition(router, "/how_to", true)}>
        <Typography sx={linkStyle} gutterBottom>
          使い方
        </Typography>
      </Button>
      <Button onClick={handlePageTransition(router, "/terms_of_use")}>
        <Typography sx={linkStyle} gutterBottom>
          利用規約
        </Typography>
      </Button>
      <Button
        onClick={handlePageTransition(
          undefined,
          "https://forms.gle/Gh3ZpW9DS64eb1qS7",
          true,
          true
        )}
      >
        <Typography sx={linkStyle} gutterBottom>
          お問い合わせ
        </Typography>
      </Button>
    </Stack>
  );
};

export default Footer;
