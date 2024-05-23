import CloudDoneIcon from "@mui/icons-material/CloudDone";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import {
  Backdrop,
  Card,
  Dialog,
  DialogContent,
  DialogTitle,
  Stack,
  Typography,
} from "@mui/material";
import { DropzoneInputProps, DropzoneRootProps } from "react-dropzone";

type InputFileAndUploadProps = {
  dialogOpen: boolean;
  setDialogOpen: (open: boolean) => void;
  selectedFile: File | null;
  getRootProps: () => DropzoneRootProps;
  getInputProps: () => DropzoneInputProps;
  isDragActive: boolean;
  title: string;
};

const InputFileAndUpload = ({
  dialogOpen,
  setDialogOpen,
  selectedFile,
  getRootProps,
  getInputProps,
  isDragActive,
  title,
}: InputFileAndUploadProps) => {
  return (
    <>
      <Backdrop
        open={isDragActive}
        sx={{
          backgroundColor: "rgba(255, 255, 255, 0.5)",
          zIndex: 1500,
          pointerEvents: "none",
        }}
      ></Backdrop>
      <Dialog
        onClose={() => setDialogOpen(false)}
        open={dialogOpen}
        maxWidth="sm"
        fullWidth
      >
        <DialogTitle>{title}</DialogTitle>
        <DialogContent>
          <Stack
            {...getRootProps()}
            sx={{
              border: "2px dashed #eeeeee",
              p: 2,
              mt: 2,
              mb: 2,
              textAlign: "center",
              alignItems: "center",
              ":hover": {
                cursor: "pointer",
              },
            }}
          >
            <input {...getInputProps()} />
            {selectedFile ? (
              <CloudDoneIcon
                sx={{ fontSize: 60, mb: 2 }}
                color="primary"
              ></CloudDoneIcon>
            ) : (
              <CloudUploadIcon sx={{ fontSize: 60, mb: 2 }} color="disabled" />
            )}

            <Typography variant="subtitle1">
              {selectedFile
                ? "新しいファイルをアップロードしたいときは、もう一度ここにファイルをドラッグしてください。"
                : isDragActive
                ? "ファイルをここにドラッグしてください。"
                : "ファイルをここにドラッグしてください。クリックで選択することもできます。"}
            </Typography>
          </Stack>

          <Stack spacing={2}>
            {selectedFile && (
              <Card variant="outlined" sx={{ mt: 3, p: 2, width: "100%" }}>
                <Typography variant="subtitle1">Selected File:</Typography>
                <Typography variant="h6" component="div">
                  {selectedFile.name}
                </Typography>
              </Card>
            )}
          </Stack>
        </DialogContent>
      </Dialog>
    </>
  );
};

export default InputFileAndUpload;
