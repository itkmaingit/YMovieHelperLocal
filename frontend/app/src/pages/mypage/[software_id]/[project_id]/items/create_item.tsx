import {
  Box,
  FormControl,
  FormControlLabel,
  FormLabel,
  Radio,
  RadioGroup,
  Stack,
  Step,
  StepLabel,
  Stepper,
  Typography,
} from "@mui/material";
import * as React from "react";

import DynamicItem from "@/components/ItemEdit/DynamicItem";
import MultipleItem from "@/components/ItemEdit/MultipleItem";
import SingleItem from "@/components/ItemEdit/SingleItem";
import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import SectionTitle from "@/components/utils/SectionTitle";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useFeedback from "@/hooks/use_feedback";

const CreateItem = () => {
  const [selectedValue, setSelectedValue] = React.useState<string | null>(null);
  const [activeStep, setActiveStep] = React.useState(0);

  const steps = ["アイテムの属性設定", "詳細設定"];

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSelectedValue(event.target.value);
    setActiveStep(1);
  };

  const [setCode, setState, snackbarProps] = useFeedback();

  return (
    <>
      <LoadingAndSnackbar {...snackbarProps}></LoadingAndSnackbar>
      <SectionTitle title="Create Item"></SectionTitle>
      <Box>
        <Stepper
          activeStep={activeStep}
          alternativeLabel
          sx={{ marginBottom: 10 }}
        >
          {steps.map((label) => (
            <Step key={label}>
              <StepLabel
                StepIconProps={{
                  sx: {
                    "&.MuiStepIcon-root": {
                      width: "36px",
                      height: "36px",
                    },
                    "& .MuiStepIcon-text": {
                      fontSize: "1.2rem",
                    },
                  },
                }}
              >
                <Typography>{label}</Typography>
              </StepLabel>
            </Step>
          ))}
        </Stepper>
        <Box>
          <Stack spacing={5}>
            <FormControl component="fieldset">
              <FormLabel component="legend" sx={{ fontSize: "1.6rem", my: 2 }}>
                Choose Item Category
              </FormLabel>
              <RadioGroup
                aria-label="modes"
                value={selectedValue}
                onChange={handleChange}
              >
                <Stack spacing={2} direction="row">
                  <FormControlLabel
                    value="dynamic"
                    control={<Radio />}
                    label="Dynamic Item"
                    sx={{
                      "& .MuiFormControlLabel-label": { fontSize: "1.4rem" },
                    }}
                  />
                  <FormControlLabel
                    value="single"
                    control={<Radio />}
                    label="Single Item"
                    sx={{
                      "& .MuiFormControlLabel-label": { fontSize: "1.4rem" },
                    }}
                  />
                  <FormControlLabel
                    value="multiple"
                    control={<Radio />}
                    label="Multiple Item"
                    sx={{
                      "& .MuiFormControlLabel-label": { fontSize: "1.4rem" },
                    }}
                  />
                </Stack>
              </RadioGroup>
            </FormControl>
            <FunctionHandlersContext.Provider
              value={{ setCode: setCode, setState: setState }}
            >
              {activeStep === 1 && (
                <>
                  {selectedValue === "multiple" && (
                    <Box>
                      <MultipleItem></MultipleItem>
                    </Box>
                  )}
                  {selectedValue === "single" && (
                    <Stack spacing={3}>
                      <SingleItem></SingleItem>
                    </Stack>
                  )}
                  {selectedValue === "dynamic" && (
                    <Box>
                      <DynamicItem></DynamicItem>
                    </Box>
                  )}
                </>
              )}
            </FunctionHandlersContext.Provider>
          </Stack>
        </Box>
      </Box>
    </>
  );
};

export default CreateItem;
