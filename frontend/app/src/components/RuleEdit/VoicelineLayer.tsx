import TextField from "@mui/material/TextField";
import React from "react";

type VoicelineLayerProps = {
  voicelineLayer: number;
  setVoicelineLayer: (voicelineLayer: number) => void;
  validation: boolean;
};
const VoicelineLayer = ({
  voicelineLayer,
  setVoicelineLayer,
  validation,
}: VoicelineLayerProps) => {
  const handleVoicelineInput = (event: React.ChangeEvent<HTMLInputElement>) => {
    setVoicelineLayer(Number(event.target.value));
  };
  return (
    <>
      <TextField
        label="Layer"
        type="number"
        value={
          voicelineLayer !== undefined && voicelineLayer !== null
            ? voicelineLayer.toString()
            : ""
        }
        onChange={handleVoicelineInput}
        error={validation && voicelineLayer === undefined}
        sx={{ width: "30%" }}
      />
    </>
  );
};

export default VoicelineLayer;
