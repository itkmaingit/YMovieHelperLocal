import { Backdrop, CircularProgress } from "@mui/material";

type Props = {
  open: boolean;
};

function FullScreenLoading({ open }: Props) {
  return (
    <Backdrop open={open} style={{ color: "#fff", zIndex: 1500 }}>
      <CircularProgress color="inherit" />
    </Backdrop>
  );
}

export default FullScreenLoading;
