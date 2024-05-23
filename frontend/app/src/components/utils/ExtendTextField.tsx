import { InputBaseComponentProps, SxProps, TextField } from "@mui/material";
import React, { useRef } from "react";

type ExtendTextFieldProps = {
  text: string;
  setText: React.Dispatch<React.SetStateAction<string>>;
  handleBlur: () => void;
  textFieldStyle: InputBaseComponentProps;
  typographyStyle: SxProps;
  error?: boolean;
  type?: "number" | "text";
};

const ExtendTextField = ({
  text,
  setText,
  handleBlur,
  textFieldStyle,
  type = "text",
  error = false,
}: ExtendTextFieldProps) => {
  const handleTextChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value);
  };

  const textFieldRef = useRef<HTMLInputElement>(null);

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      textFieldRef.current?.blur();
    }
  };

  return (
    <TextField
      value={text}
      onChange={handleTextChange}
      onBlur={handleBlur}
      autoFocus
      inputProps={textFieldStyle}
      margin="dense"
      onKeyDown={handleKeyDown}
      inputRef={textFieldRef}
      type={type}
      error={error}
    />
  );
};

export default ExtendTextField;
