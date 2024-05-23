import { Box, Stack, Typography, useMediaQuery } from "@mui/material";
import Image from "next/image";

import theme from "@/theme";

type ItemProps = {
  src: string;
  alt: string;
  title: string;
  description: string;
};

const Item = ({ src, alt, title, description }: ItemProps) => {
  const matches = useMediaQuery(theme.breakpoints.down("md"));

  return matches ? (
    <Stack
      direction="column"
      sx={{
        alignItems: "center",
        justifyContent: "center",
        textAlign: "center",
        my: "100px",
        // width: "70%",
        mx: 2,
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          mb: 8,
        }}
      >
        <Box
          sx={{
            width: "100px ",
            minWidth: "100px ",
            height: "100px ",
            borderRadius: "50%",
            position: "relative",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            backgroundColor: "primary.main",
            mr: 4,
          }}
        >
          <Box
            sx={{
              width: 60,
              height: 60,
              position: "relative",
            }}
          >
            <Image src={src} alt={alt} fill style={{ objectFit: "contain" }} />
          </Box>
        </Box>
        <Typography
          color="primary.main"
          mb={0.5}
          sx={{ fontSize: "clamp(1.2rem, 4vw, 3rem)" }}
        >
          {title}
        </Typography>
      </Box>
      <Box>
        <Typography sx={{ fontSize: "clamp(1.2rem, 3vw, 3rem)" }}>
          {description}
        </Typography>
      </Box>
    </Stack>
  ) : (
    <Stack
      direction="row"
      sx={{
        alignItems: "center",
        my: "100px",
        ml: 20,
      }}
    >
      <Box
        sx={{
          width: "100px ",
          minWidth: "100px ",
          height: "100px ",
          borderRadius: "50%",
          position: "relative",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          backgroundColor: "primary.main",
          mr: 8,
        }}
      >
        <Box
          sx={{
            width: 60,
            height: 60,
            position: "relative",
          }}
        >
          <Image src={src} alt={alt} fill style={{ objectFit: "contain" }} />
        </Box>
      </Box>
      <Box pr={5}>
        <Typography variant="h4" color="primary.main" mb={0.5}>
          {title}
        </Typography>
        <Typography variant="h5">{description}</Typography>
      </Box>
    </Stack>
  );
};

export default Item;
