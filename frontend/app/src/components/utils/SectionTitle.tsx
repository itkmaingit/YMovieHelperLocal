import { Stack, Typography } from "@mui/material";

type SectionTitleProps = {
  title: string;
  variant?: string;
};
const SectionTitle = ({ title }: SectionTitleProps) => {
  return (
    <Stack spacing={1} sx={{ my: "80px" }}>
      <Typography variant="h2">{title}</Typography>
    </Stack>
  );
};

export default SectionTitle;
