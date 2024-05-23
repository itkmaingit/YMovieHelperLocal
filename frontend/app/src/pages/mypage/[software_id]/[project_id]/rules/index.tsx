import SaveIcon from "@mui/icons-material/Save";
import { Box, Fab, Stack, Typography } from "@mui/material";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { parseCookies } from "nookies";
import { useState } from "react";

import CharacterItemSelection, {
  CharacterItem,
  SelectedCharacterItem,
} from "@/components/RuleEdit/CharacterItemSelection";
import DynamicItemSelection, {
  DynamicItem,
  SelectedDynamicItem,
} from "@/components/RuleEdit/DynamicItemSelection";
import MultipleItemSelection, {
  MultipleItem,
  SelectedMultipleItem,
} from "@/components/RuleEdit/MultipleItemSelection";
import SingleItemSelection, {
  SelectedSingleItem,
  SingleItem,
} from "@/components/RuleEdit/SingleItemSelection";
import VoicelineLayer from "@/components/RuleEdit/VoicelineLayer";
import LoadingAndSnackbar from "@/components/utils/LoadingAndSnackbar";
import SectionTitle from "@/components/utils/SectionTitle";
import useAsyncAndLoading from "@/hooks/use_async_and_loading";
import useFeedback from "@/hooks/use_feedback";
import { clientAxios, serverAxios } from "@/libs/axios";
import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

//TODO: TextFieldを全てExtendTextFieldにいつかは変更する...
//TODO: selectedItemsを全てcontext化する
export const getServerSideProps: GetServerSideProps = async (context) => {
  const token = parseCookies(context).token;
  const softwareID = context.params?.software_id;
  const projectID = context.params?.project_id;
  const getConfig = {
    method: "get",
    url: `/softwares/${softwareID}/projects/${projectID}/rules`,
    headers: { Authorization: `Bearer ${token}` },
  };
  const response = await serverAxios(getConfig).catch(() => {
    return null;
  });
  let data;
  if (response?.data.initialSelectedCharacterItems === null) {
    response.data.initialSelectedCharacterItems = [{}];
  }
  if (response?.data.initialSelectedDynamicItems === null) {
    response.data.initialSelectedDynamicItems = [];
  }
  if (response?.data.initialSelectedSingleItems === null) {
    response.data.initialSelectedSingleItems = [];
  }
  if (response?.data.initialSelectedMultipleItems === null) {
    response.data.initialSelectedMultipleItems = [];
  }
  if (response != null) {
    data = response.data;
  } else {
    data = null;
  }
  return { props: { ...data } };
};

type InitialProps = {
  initialVoicelineLayer: number;
  charaItems: CharacterItem[];
  dynamicItems: DynamicItem[];
  singleItems: SingleItem[];
  multipleItems: MultipleItem[];
  initialSelectedCharacterItems: SelectedCharacterItem[];
  initialSelectedDynamicItems: SelectedDynamicItem[];
  initialSelectedSingleItems: SelectedSingleItem[];
  initialSelectedMultipleItems: SelectedMultipleItem[];
};

const RulePage = ({
  initialVoicelineLayer,
  charaItems,
  dynamicItems,
  singleItems,
  multipleItems,
  initialSelectedCharacterItems,
  initialSelectedDynamicItems,
  initialSelectedSingleItems,
  initialSelectedMultipleItems,
}: InitialProps) => {
  const router = useRouter();
  const { software_id, project_id } = router.query;
  // -------------------初期値---------------------
  const [voicelineLayer, setVoicelineLayer] = useState<number>(
    initialVoicelineLayer
  );
  const [selectedCharacterItems, setSelectedCharacterItems] = useState<
    SelectedCharacterItem[]
  >(initialSelectedCharacterItems);

  const [selectedDynamicItems, setSelectedDynamicItems] = useState<
    SelectedDynamicItem[]
  >(initialSelectedDynamicItems);

  const [selectedSingleItems, setSelectedSingleItems] = useState<
    SelectedSingleItem[]
  >(initialSelectedSingleItems);

  const [selectedMultipleItems, setSelectedMultipleItems] = useState<
    SelectedMultipleItem[]
  >(initialSelectedMultipleItems);

  // ---------------------------バリデーション-----------------------
  const [validation, setValidation] = useState<boolean>(false);

  const voicelineLayerOK = () => {
    if (voicelineLayer === undefined || voicelineLayer < 0) {
      return false;
    }
    return true;
  };

  const characterItemsOK = () => {
    const allFieldsFilled = selectedCharacterItems
      ? selectedCharacterItems.every((item) => {
          if (item.id === undefined) {
            return false;
          }
          if (item.isEmpty) {
            if (
              item.sentence === undefined ||
              item.sentence.length > 255 ||
              isEmptyOrWhitespace(item.sentence)
            ) {
              return false;
            }
          }
          return true;
        })
      : true;

    return allFieldsFilled;
  };

  const dynamicItemsOK = () => {
    const allFieldsFilled = selectedDynamicItems
      ? selectedDynamicItems.every((item) => {
          if (item.id === undefined && item.layer === undefined) {
            return false;
          }
          return true;
        })
      : true;

    return allFieldsFilled;
  };

  const singleItemsOK = () => {
    const allFieldsFilled = selectedSingleItems
      ? selectedSingleItems.every((item) => {
          if (
            item.id === undefined ||
            item.layer === undefined ||
            item.isFixedStart === undefined ||
            item.isFixedEnd === undefined
          ) {
            return false;
          }

          if (item.isFixedStart === true) {
            if (item.start.insertPlace === undefined) {
              return false;
            }
          }

          if (item.isFixedStart === false) {
            if (
              item.start.characterID === undefined ||
              item.start.adjustmentValue === undefined
            ) {
              return false;
            }
          }

          if (item.isFixedEnd === true) {
            if (item.end.isUnique === undefined) {
              return false;
            }
            if (item.end.isUnique === false) {
              if (item.end.length === undefined) {
                return false;
              }
            }
          }

          if (item.isFixedEnd === false) {
            if (
              item.end.adjustmentValue === undefined ||
              item.end.howManyAheads === undefined ||
              item.end.howManyAheads < 0
            ) {
              return false;
            }
          }

          return true;
        })
      : true;
    return allFieldsFilled;
  };

  const multipleItemsOK = () => {
    const allFieldsFilled = selectedMultipleItems
      ? selectedMultipleItems.every((item) => {
          if (
            item.id === undefined ||
            item.layer === undefined ||
            item.isFixedStart === undefined
          ) {
            return false;
          }

          if (item.isFixedStart === true) {
            if (item.start.insertPlace === undefined) {
              return false;
            }
          }

          if (item.isFixedStart === false) {
            if (
              item.start.characterID === undefined ||
              item.start.adjustmentValue === undefined
            ) {
              return false;
            }
          }

          return true;
        })
      : true;

    return allFieldsFilled;
  };

  const saveClick = async () => {
    if (
      voicelineLayerOK() &&
      characterItemsOK() &&
      dynamicItemsOK() &&
      singleItemsOK() &&
      multipleItemsOK()
    ) {
      //APIを叩く
      setValidation(false);

      const token = parseCookies().token;
      const uploadConfig = {
        method: "post",
        url: `/softwares/${software_id}/projects/${project_id}/rules`,
        data: {
          voicelineLayer: voicelineLayer,
          charaItems: selectedCharacterItems,
          dynamicItems: selectedDynamicItems,
          singleItems: selectedSingleItems,
          multipleItems: selectedMultipleItems,
        },
        headers: { Authorization: `Bearer ${token}` },
      };

      await clientAxios(uploadConfig)
        .then(() => {
          setCode({ message: "ルール保存完了", status: "success" });
        })
        .catch(() => {
          setCode({
            message: "サーバーエラー！時間をおいてから試してください！",
            status: "warning",
          });
        });
    } else {
      setValidation(true);
      setCode({
        message:
          "未入力、あるいは不正な値があります！ 赤くなっている入力欄を確認してください！",
        status: "warning",
      });
    }
  };

  const [setCode, setState, snackbarProps] = useFeedback();

  const handleSaveClick = useAsyncAndLoading(saveClick, setState);

  //なぜかバグが発生するのでコメントアウト
  // useEffect(() => {
  //   const handleKeyDown = async (event: KeyboardEvent) => {
  //     if (event.ctrlKey && event.key === "s") {
  //       // Prevent browser's default save dialog
  //       event.preventDefault();

  //       // Call your save function
  //       await handleSaveClick();
  //     }
  //   };

  //   window.addEventListener("keydown", handleKeyDown);

  //   // Cleanup after the effect:
  //   return () => {
  //     window.removeEventListener("keydown", handleKeyDown);
  //   };
  // }, []);

  return (
    <>
      {/* --------------------------Hidden or sticky----------------------- */}
      <Fab
        variant="extended"
        color="primary"
        aria-label="add"
        sx={{
          position: "fixed",
          right: 120,
          bottom: 160,
          width: "160px", // ボタンの幅を調整
          height: "60px", // ボタンの高さを調整
        }}
        onClick={handleSaveClick}
      >
        <SaveIcon sx={{ fontSize: 40, marginRight: 1 }} />
        <Typography variant="h5"> Save</Typography>
      </Fab>
      <LoadingAndSnackbar {...snackbarProps}></LoadingAndSnackbar>

      {/* ----------------------------Hidden or sticky----------------------- */}
      {/* ------------------------------Main---------------------------------- */}
      <Stack spacing={2}>
        <SectionTitle title="Rule Definition"></SectionTitle>
        <Typography variant="h4" component="h1">
          Voiceline
        </Typography>
        <VoicelineLayer
          voicelineLayer={voicelineLayer}
          setVoicelineLayer={setVoicelineLayer}
          validation={validation}
        ></VoicelineLayer>
        <Box>
          <CharacterItemSelection
            characterItems={charaItems}
            selectedCharacterItems={selectedCharacterItems}
            setSelectedCharacterItems={setSelectedCharacterItems}
            selectedSingleItems={selectedSingleItems}
            setSelectedSingleItems={setSelectedSingleItems}
            selectedMultipleItems={selectedMultipleItems}
            setSelectedMultipleItems={setSelectedMultipleItems}
            validation={validation}
          />

          <DynamicItemSelection
            dynamicItems={dynamicItems}
            selectedDynamicItems={selectedDynamicItems}
            setSelectedDynamicItems={setSelectedDynamicItems}
            validation={validation}
          />

          <SingleItemSelection
            singleItems={singleItems}
            selectedSingleItems={selectedSingleItems}
            setSelectedSingleItems={setSelectedSingleItems}
            selectedCharacterItems={selectedCharacterItems}
            validation={validation}
          />

          <MultipleItemSelection
            multipleItems={multipleItems}
            selectedMultipleItems={selectedMultipleItems}
            setSelectedMultipleItems={setSelectedMultipleItems}
            selectedCharacterItems={selectedCharacterItems}
            validation={validation}
          />
        </Box>
      </Stack>
    </>
  );
};

export default RulePage;
