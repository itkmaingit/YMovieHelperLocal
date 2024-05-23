import { Stack, Step, StepLabel, Stepper } from "@mui/material";
import { GetServerSideProps } from "next";
import { parseCookies } from "nookies";
import * as React from "react";

import DownloadCompleteYMMP from "@/components/MakeYMMP/DownloadCompleteYMMP";
import DownloadScenarioTxt from "@/components/MakeYMMP/DownloadScenarioTxt";
import DownloadTemplateCSV from "@/components/MakeYMMP/DownloadTemplateCSV";
import InputCSV from "@/components/MakeYMMP/InputCSV";
import InputYMMP from "@/components/MakeYMMP/InputYMMP";
import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import SectionTitle from "@/components/utils/SectionTitle";
import FunctionHandlersContext from "@/contexts/FunctionHandlersContext";
import useFeedback from "@/hooks/use_feedback";
import { serverAxios } from "@/libs/axios";

export const getServerSideProps: GetServerSideProps = async (context) => {
  const token = parseCookies(context).token;
  const softwareID = context.params?.software_id;
  const projectID = context.params?.project_id;
  const getConfig = {
    method: "get",
    url: `/softwares/${softwareID}/projects/${projectID}/make_ymmp/can_make_ymmp`,
    headers: { Authorization: `Bearer ${token}` },
  };

  try {
    const response = await serverAxios(getConfig);

    if (response.data.success == "false") {
      return {
        redirect: {
          destination: `/mypage/${softwareID}/${projectID}/rules`,
          permanent: false,
        },
      };
    }
  } catch (error) {
    return {
      redirect: {
        destination: `/mypage/${softwareID}/${projectID}/rules`,
        permanent: false,
      },
    };
  }

  return { props: {} };
};

const steps = [
  "テンプレートCSVのダウンロード",
  "台本CSVのアップロード",
  "scenario.txtのダウンロード",
  "ymmpファイルのアップロード",
  "complete.ymmpのダウンロード",
];

const MakeYMMP = () => {
  const [step, setStep] = React.useState<number>(0);

  const [setCode, setState, snackbarProps] = useFeedback();

  return (
    <>
      <LoadingAndSnackbar {...snackbarProps}></LoadingAndSnackbar>
      <SectionTitle title="Create YMMP"></SectionTitle>

      <Stack>
        <Stepper activeStep={step}>
          {steps.map((label) => {
            const stepProps: { completed?: boolean } = {};
            const labelProps: {
              optional?: React.ReactNode;
            } = {};
            return (
              <Step key={label} {...stepProps} sx={{ mb: 10 }}>
                <StepLabel
                  {...labelProps}
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
                  {label}
                </StepLabel>
              </Step>
            );
          })}
        </Stepper>
        <FunctionHandlersContext.Provider
          value={{ setCode: setCode, setState: setState }}
        >
          {step === 0 && (
            <DownloadTemplateCSV setStep={setStep}></DownloadTemplateCSV>
          )}

          {step === 1 && <InputCSV setStep={setStep}></InputCSV>}
          {step === 2 && (
            <DownloadScenarioTxt setStep={setStep}></DownloadScenarioTxt>
          )}
          {step === 3 && <InputYMMP setStep={setStep}></InputYMMP>}
          {step === 4 && (
            <DownloadCompleteYMMP setStep={setStep}></DownloadCompleteYMMP>
          )}
        </FunctionHandlersContext.Provider>
      </Stack>
    </>
  );
};

export default MakeYMMP;
