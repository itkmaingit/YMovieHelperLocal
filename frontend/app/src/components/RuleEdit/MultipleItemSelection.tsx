import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import HighlightOffIcon from "@mui/icons-material/HighlightOff";
import {
  FormControl,
  FormControlLabel,
  FormLabel,
  Grid,
  IconButton,
  InputLabel,
  MenuItem,
  Paper,
  Radio,
  RadioGroup,
  Select,
  SelectChangeEvent,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import * as React from "react";

import { SelectedCharacterItem } from "./CharacterItemSelection";

export type MultipleItem = {
  id: number;
  name: string;
};

export type MultipleItemSelectionProps = {
  multipleItems: MultipleItem[];
  selectedCharacterItems: SelectedCharacterItem[];
  selectedMultipleItems: SelectedMultipleItem[];
  setSelectedMultipleItems: (items: SelectedMultipleItem[]) => void;
  validation: boolean;
};

export type SelectedMultipleItem = {
  id?: number;
  layer?: number;
  start: Start;
  isFixedStart?: boolean;
};

export type Start = {
  insertPlace?: string;
  characterID?: number;
  characterName?: string;
  adjustmentValue?: number;
};

const MultipleItemSelection = ({
  multipleItems,
  selectedCharacterItems,
  selectedMultipleItems,
  setSelectedMultipleItems,
  validation,
}: MultipleItemSelectionProps) => {
  //やばいことをやっているが、結局すべて入力値に対してselectedMultipleItemsを更新する

  const fixedStartList: string[] = ["最初", "最後"];

  const handleItemSelect =
    (index: number) => (event: SelectChangeEvent<number>) => {
      const newSelectedMultipleItems = [...selectedMultipleItems];
      const selectedMultipleItem = {
        id: Number(event.target.value),
        start: {},
        end: {},
      };
      newSelectedMultipleItems[index] = selectedMultipleItem;
      setSelectedMultipleItems(newSelectedMultipleItems);
    };

  const handleAddClick = () => {
    const newSelectedMultipleItems = selectedMultipleItems
      ? [...selectedMultipleItems]
      : [];
    newSelectedMultipleItems.push({ start: {} });
    setSelectedMultipleItems(newSelectedMultipleItems);
  };

  const handleDeleteClick = (index: number) => () => {
    const newSelectedMultipleItems = [...selectedMultipleItems];
    newSelectedMultipleItems.splice(index, 1);
    setSelectedMultipleItems(newSelectedMultipleItems);
  };

  const handleLayerChange =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedMultipleItems = [...selectedMultipleItems];
      newSelectedMultipleItems[index].layer = Number(event.target.value);
      setSelectedMultipleItems(newSelectedMultipleItems);
    };

  const handleSelectStart =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedMultipleItems = [...selectedMultipleItems];
      const isFixed = event.target.value === "true";
      newSelectedMultipleItems[index].isFixedStart = isFixed;
      if (isFixed) {
        newSelectedMultipleItems[index].start.insertPlace = undefined;
      } else if (!isFixed) {
        newSelectedMultipleItems[index].start.characterID = undefined;
        newSelectedMultipleItems[index].start.adjustmentValue = undefined;
      }
      setSelectedMultipleItems(newSelectedMultipleItems);
    };

  const handleCharacterSelect =
    (index: number) => (event: SelectChangeEvent<number>) => {
      const newSelectedMultipleItems = [...selectedMultipleItems];
      newSelectedMultipleItems[index].start.characterID = Number(
        event.target.value
      );
      const character = selectedCharacterItems.find(
        (item) => item.id === Number(event.target.value)
      );
      newSelectedMultipleItems[index].start.characterName = character?.name;
      setSelectedMultipleItems(newSelectedMultipleItems);
    };

  const handleInputAdjustmentValueInStart =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedMultipleItems = [...selectedMultipleItems];
      newSelectedMultipleItems[index].start.adjustmentValue = Number(
        event.target.value
      );
      setSelectedMultipleItems(newSelectedMultipleItems);
    };

  const handleSelectFixedStart =
    (index: number) => (event: SelectChangeEvent<string>) => {
      const newSelectedMultipleItems = [...selectedMultipleItems];
      newSelectedMultipleItems[index].start.insertPlace = event.target.value;
      setSelectedMultipleItems(newSelectedMultipleItems);
    };

  return (
    <>
      <Typography variant="h4" component="h1" sx={{ my: 7 }}>
        Multiple Item Rule
      </Typography>
      <Stack spacing={2} alignItems="center">
        <Grid container>
          {selectedMultipleItems &&
            selectedMultipleItems.map(
              (item: SelectedMultipleItem, index: number) => (
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
                          <InputLabel id="multiple-item">
                            Multiple Item
                          </InputLabel>
                          <Select
                            value={selectedMultipleItems[index]?.id || ""}
                            onChange={handleItemSelect(index)}
                            label="MultipleItem"
                            labelId="multiple-item"
                            error={
                              validation &&
                              selectedMultipleItems[index]?.id === undefined
                            }
                          >
                            {multipleItems &&
                              multipleItems.map((item) => (
                                <MenuItem value={item.id} key={item.id}>
                                  {item.name}
                                </MenuItem>
                              ))}
                          </Select>
                        </FormControl>
                        {selectedMultipleItems[index]?.id && (
                          <TextField
                            label="Layer"
                            type="number"
                            value={
                              selectedMultipleItems[index]?.layer !==
                                undefined &&
                              selectedMultipleItems[index]?.layer !== null
                                ? (
                                    selectedMultipleItems[index]
                                      ?.layer as number
                                  ).toString()
                                : ""
                            }
                            onChange={handleLayerChange(index)}
                            error={
                              validation &&
                              selectedMultipleItems[index]?.layer === undefined
                            }
                          />
                        )}
                      </Stack>
                    </Grid>

                    {selectedMultipleItems[index]?.id && (
                      <>
                        <Grid item sx={{ margin: 2 }}>
                          <Stack direction="row" spacing={20}>
                            <Stack direction="column" spacing={3}>
                              <FormControl
                                sx={{ minWidth: 200 }}
                                error={
                                  validation &&
                                  selectedMultipleItems[index].isFixedStart ===
                                    undefined
                                }
                              >
                                <FormLabel id="start-label">始点</FormLabel>
                                <RadioGroup
                                  row
                                  aria-labelledby="start-label"
                                  name="start-radio-group"
                                  onChange={handleSelectStart(index)}
                                  value={
                                    selectedMultipleItems[index]
                                      .isFixedStart === undefined
                                      ? ""
                                      : selectedMultipleItems[index]
                                          .isFixedStart
                                  }
                                >
                                  <FormControlLabel
                                    value={true}
                                    control={<Radio />}
                                    label="固定始点"
                                  />
                                  <FormControlLabel
                                    value={false}
                                    control={<Radio />}
                                    label="自由始点"
                                  />
                                </RadioGroup>
                              </FormControl>
                              {selectedMultipleItems[index].isFixedStart ===
                                true && (
                                <FormControl sx={{ width: 210 }}>
                                  <InputLabel id="fixed-start">
                                    挿入位置
                                  </InputLabel>
                                  <Select
                                    label="挿入位置"
                                    labelId="fixed-start"
                                    onChange={handleSelectFixedStart(index)}
                                    value={
                                      selectedMultipleItems[index].start
                                        ?.insertPlace || ""
                                    }
                                    error={
                                      validation &&
                                      selectedMultipleItems[index].start
                                        ?.insertPlace === undefined
                                    }
                                  >
                                    {fixedStartList.map((item) => (
                                      <MenuItem value={item} key={item}>
                                        {item}
                                      </MenuItem>
                                    ))}
                                  </Select>
                                </FormControl>
                              )}

                              {selectedMultipleItems[index].isFixedStart ===
                                false && (
                                <>
                                  <FormControl sx={{ minWidth: 200 }}>
                                    <InputLabel id="character-item">
                                      Character Item
                                    </InputLabel>
                                    <Select
                                      value={
                                        selectedMultipleItems[index].start
                                          .characterID || ""
                                      }
                                      onChange={handleCharacterSelect(index)}
                                      label="Character Item"
                                      labelId="character-item"
                                      error={
                                        validation &&
                                        selectedMultipleItems[index].start
                                          ?.characterID === undefined
                                      }
                                    >
                                      {selectedCharacterItems &&
                                        selectedCharacterItems.map((item) => (
                                          // やってることは高度だが、結局選択されたアイテムをドロップダウンで選択できなくしただけ
                                          <MenuItem
                                            value={item.id}
                                            key={`multiple-character-${item.id}`}
                                          >
                                            {item.name}
                                          </MenuItem>
                                        ))}
                                    </Select>
                                  </FormControl>
                                  <TextField
                                    label="調整値(フレーム単位)"
                                    type="number"
                                    value={
                                      selectedMultipleItems[index].start
                                        ?.adjustmentValue !== undefined &&
                                      selectedMultipleItems[index].start
                                        ?.adjustmentValue !== null
                                        ? (
                                            selectedMultipleItems[index].start
                                              ?.adjustmentValue as number
                                          ).toString()
                                        : ""
                                    }
                                    onChange={handleInputAdjustmentValueInStart(
                                      index
                                    )}
                                    sx={{ minWidth: 200 }}
                                    error={
                                      validation &&
                                      selectedMultipleItems[index].start
                                        ?.adjustmentValue === undefined
                                    }
                                  />
                                </>
                              )}
                            </Stack>
                          </Stack>
                        </Grid>
                      </>
                    )}
                  </Paper>
                </React.Fragment>
              )
            )}
        </Grid>
        <div style={{ width: "min-content" }}>
          <IconButton
            onClick={handleAddClick}
            sx={{ marginTop: "0px !important" }}
          >
            <AddCircleOutlineIcon fontSize="large" />
          </IconButton>
        </div>
      </Stack>
    </>
  );
};

export default MultipleItemSelection;
