import {
  InputBaseComponentProps,
  SxProps,
  TextField,
  Typography,
} from "@mui/material";
import React, { useRef } from "react";

type EditableTextProps = {
  text: string;
  setText: React.Dispatch<React.SetStateAction<string>>;
  isEditing: boolean;
  handleBlur: () => void;
  handleTextClick?: () => void;
  textFieldStyle: InputBaseComponentProps;
  typographyStyle: SxProps;
  type?: "number" | "text";
};

const EditableText = ({
  text,
  setText,
  handleBlur,
  isEditing,
  handleTextClick,
  textFieldStyle,
  typographyStyle,
  type = "text",
}: EditableTextProps) => {
  const handleTextChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value);
  };

  const textFieldRef = useRef<HTMLInputElement>(null);

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      textFieldRef.current?.blur();
    }
  };

  return isEditing ? (
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
    />
  ) : (
    <Typography
      variant="body1"
      sx={typographyStyle}
      onClick={handleTextClick && handleTextClick}
    >
      {text}
    </Typography>
  );
};

export default EditableText;
