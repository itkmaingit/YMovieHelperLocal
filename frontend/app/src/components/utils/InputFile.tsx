/** @jsxImportSource @emotion/react */

import CloudDoneIcon from "@mui/icons-material/CloudDone";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import { Backdrop, Card, Stack, TextField } from "@mui/material";
import Typography from "@mui/material/Typography";
import React from "react";
import { useDropzone } from "react-dropzone";

export type InputFileProps = {
  displayName?: string;
  inputText?: string;
  setInputText?: (updateText: string) => void;
  selectedFile: null | File;
  onDrop: (acceptedFiles: File[]) => void;
  children: React.ReactNode;
};

const InputFile: React.FC<InputFileProps> = ({
  displayName,
  inputText,
  setInputText,
  selectedFile,
  onDrop,
  children,
}) => {
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

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
      <Stack spacing={2} sx={{ width: "100%" }}>
        {selectedFile && (
          <Card variant="outlined" sx={{ mt: 3, p: 2 }}>
            <Typography variant="subtitle1">Selected File:</Typography>
            <Typography variant="h6" component="div">
              {selectedFile.name}
            </Typography>
          </Card>
        )}
        {inputText !== undefined && setInputText !== undefined && (
          <TextField
            label={displayName}
            value={inputText}
            onChange={(e) => setInputText(e.target.value)}
            fullWidth
            sx={{ mt: 2 }}
            error={inputText.length > 20}
            helperText={inputText.length > 20 && "文字数は20文字までです！"}
          />
        )}
        {children}
      </Stack>
    </>
  );
};

export default InputFile;
