import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import HighlightOffIcon from "@mui/icons-material/HighlightOff";
import {
  FormControl,
  Grid,
  IconButton,
  InputLabel,
  MenuItem,
  Paper,
  Select,
  SelectChangeEvent,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import * as React from "react";

import { SelectedMultipleItem } from "./MultipleItemSelection";
import { SelectedSingleItem } from "./SingleItemSelection";

import { isEmptyOrWhitespace } from "@/libs/utils/utils_function";

export type CharacterItem = {
  ID: number;
  Name: string;
  IsEmpty: boolean;
  Sentence: string;
};

type CharacterItemSelectionProps = {
  characterItems: CharacterItem[];
  selectedCharacterItems: SelectedCharacterItem[];
  setSelectedCharacterItems: (items: SelectedCharacterItem[]) => void;
  selectedSingleItems: SelectedSingleItem[];
  setSelectedSingleItems: (items: SelectedSingleItem[]) => void;
  selectedMultipleItems: SelectedMultipleItem[];
  setSelectedMultipleItems: (items: SelectedMultipleItem[]) => void;
  validation: boolean;
};

export type SelectedCharacterItem = {
  id?: number;
  sentence?: string;
  name?: string;
  isEmpty?: boolean;
};

const CharacterItemSelection = ({
  characterItems,
  selectedCharacterItems,
  setSelectedCharacterItems,
  selectedSingleItems,
  setSelectedSingleItems,
  selectedMultipleItems,
  setSelectedMultipleItems,
  validation,
}: CharacterItemSelectionProps) => {
  // イベントハンドラは全て、新しく入力された値をselectedCharacterItemsに入れているか、要素を追加、削除しているだけ
  const handleSentenceInput =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedCharacterItems = [...selectedCharacterItems];
      newSelectedCharacterItems[index].sentence = event.target.value;
      setSelectedCharacterItems(newSelectedCharacterItems);
    };

  //SingleItemsとMultipleItemsが、選択前のCharacterItemを参照していなかったかを確認し、参照されていればundefinedに戻す
  const handleCharacterSelect =
    (index: number) => (event: SelectChangeEvent<number>) => {
      const newSelectedCharacterItems = [...selectedCharacterItems];
      const selectedItem = characterItems.find(
        (item) => item.ID === Number(event.target.value)
      );
      const selectedItemName = selectedItem?.Name;
      const selectedItemIsEmpty = selectedItem?.IsEmpty;
      const selectedItemSentence = selectedItem?.Sentence;
      const selectedCharacterItem: SelectedCharacterItem = {
        id: Number(event.target.value),
        name: selectedItemName,
        isEmpty: selectedItemIsEmpty,
        sentence: selectedItemSentence,
      };
      newSelectedCharacterItems[index] = selectedCharacterItem;
      setSelectedCharacterItems(newSelectedCharacterItems);
    };

  const handleAddClick = () => {
    const newSelectedCharacterItems = [...selectedCharacterItems];
    newSelectedCharacterItems.push({ sentence: "空白" });
    setSelectedCharacterItems(newSelectedCharacterItems);
  };

  const handleDeleteClick = (index: number) => () => {
    const newSelectedCharacterItems = [...selectedCharacterItems];
    newSelectedCharacterItems.splice(index, 1);
    setSelectedCharacterItems(newSelectedCharacterItems);
  };

  React.useEffect(() => {
    if (selectedSingleItems != null) {
      selectedSingleItems.forEach((singleItem, index) => {
        const includes = selectedCharacterItems.some((characterItem) => {
          return characterItem.id === singleItem.start?.characterID;
        });
        if (!includes) {
          const newSelectedSingleItems = [...selectedSingleItems];
          if (newSelectedSingleItems[index].start?.characterID) {
            newSelectedSingleItems[index].start.characterID = undefined;
            newSelectedSingleItems[index].start.characterName = undefined;
          }
          setSelectedSingleItems(newSelectedSingleItems);
        }
      });
    }

    if (selectedMultipleItems != null) {
      selectedMultipleItems.forEach((singleItem, index) => {
        const includes = selectedCharacterItems.some((characterItem) => {
          return characterItem.id === singleItem.start?.characterID;
        });
        if (!includes) {
          const newSelectedMultipleItems = [...selectedMultipleItems];
          if (newSelectedMultipleItems[index].start?.characterID) {
            newSelectedMultipleItems[index].start.characterID = undefined;
            newSelectedMultipleItems[index].start.characterName = undefined;
          }
          setSelectedMultipleItems(newSelectedMultipleItems);
        }
      });
    }
  }, [selectedCharacterItems]);

  return (
    <>
      <Typography variant="h4" component="h1" sx={{ my: 7 }}>
        Character Rule
      </Typography>
      <Stack spacing={2} alignItems="center">
        <Grid container>
          {selectedCharacterItems &&
            selectedCharacterItems.map(
              (item: SelectedCharacterItem, index: number) => (
                <React.Fragment key={index}>
                  <Paper
                    sx={{
                      padding: 5,
                      marginBottom: 6,
                      position: "relative",
                      width: "100%",
                    }}
                    elevation={5}
                  >
                    <IconButton
                      onClick={handleDeleteClick(index)}
                      disabled={selectedCharacterItems.length === 1}
                      sx={{ position: "absolute", right: 20, top: 20 }}
                    >
                      <HighlightOffIcon fontSize="large" />
                    </IconButton>
                    <Grid
                      item
                      sx={{
                        margin: 2,
                        marginBottom: 5,
                      }}
                    >
                      <Stack direction="row" spacing={20}>
                        <FormControl sx={{ minWidth: 200 }}>
                          <InputLabel id="character-item">Character</InputLabel>
                          <Select
                            value={selectedCharacterItems[index]?.id || ""}
                            onChange={handleCharacterSelect(index)}
                            label="CharacterItem"
                            labelId="character-item"
                            error={
                              validation &&
                              selectedCharacterItems[index]?.id === undefined
                            }
                          >
                            {characterItems &&
                              characterItems.map((item) => (
                                // やってることは高度だが、結局選択されたアイテムをドロップダウンで選択できなくしただけ
                                <MenuItem
                                  value={item.ID}
                                  key={item.ID}
                                  disabled={
                                    item.ID !==
                                      selectedCharacterItems[index]?.id &&
                                    selectedCharacterItems
                                      .map((i) => i.id)
                                      .includes(item.ID)
                                  }
                                >
                                  {item.Name}
                                </MenuItem>
                              ))}
                          </Select>
                        </FormControl>
                        {/* 選択されたアイテムのisEmptyがtrueならSentenceの入力フィールドを表示 */}
                        {characterItems &&
                          characterItems.find(
                            (item) =>
                              item.ID === selectedCharacterItems[index]?.id
                          )?.IsEmpty && (
                            <TextField
                              label="空読み文"
                              value={
                                selectedCharacterItems[index]?.sentence || ""
                              }
                              onChange={handleSentenceInput(index)}
                              error={
                                validation &&
                                (selectedCharacterItems[index]?.sentence ===
                                  undefined ||
                                  (
                                    selectedCharacterItems[index]
                                      ?.sentence as string
                                  )?.length > 255 ||
                                  isEmptyOrWhitespace(
                                    selectedCharacterItems[index]
                                      ?.sentence as string
                                  ))
                              }
                            />
                          )}
                      </Stack>
                    </Grid>
                  </Paper>
                </React.Fragment>
              )
            )}
        </Grid>
        <div style={{ width: "min-content" }}>
          <IconButton
            onClick={handleAddClick}
            disabled={
              characterItems &&
              selectedCharacterItems &&
              selectedCharacterItems.length === characterItems.length
            }
            sx={{ marginTop: "0px !important" }}
          >
            <AddCircleOutlineIcon fontSize="large" />
          </IconButton>
        </div>
      </Stack>
    </>
  );
};

export default CharacterItemSelection;
