import AddIcon from "@mui/icons-material/Add";
import { Fab } from "@mui/material";

type Props = {
  clickFunction: (value: any) => void;
};
const GlobalAddButton = ({ clickFunction }: Props) => {
  return (
    <Fab
      color="primary"
      aria-label="add"
      size="large"
      sx={{
        position: "fixed",
        right: 200,
        bottom: 200,
        width: 80, // Add this line
        height: 80, // Add this line
      }}
      onClick={clickFunction}
    >
      <AddIcon sx={{ fontSize: 50 }} />
    </Fab>
  );
};

export default GlobalAddButton;
