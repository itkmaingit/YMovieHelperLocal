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

export type SingleItem = {
  id: number;
  name: string;
  length: number;
  ItemType: string;
};

export type SingleItemSelectionProps = {
  singleItems: SingleItem[];
  selectedCharacterItems: SelectedCharacterItem[];
  selectedSingleItems: SelectedSingleItem[];
  setSelectedSingleItems: (items: SelectedSingleItem[]) => void;
  validation: boolean;
};

export type SelectedSingleItem = {
  id?: number;
  layer?: number;
  isFixedStart?: boolean;
  isFixedEnd?: boolean;
  start: Start;
  end: End;
};

export type Start = {
  insertPlace?: string;
  characterID?: number;
  characterName?: string;
  adjustmentValue?: number;
};

export type End = {
  isUnique?: boolean;
  length?: number;
  adjustmentValue?: number;
  howManyAheads?: number;
};

const SingleItemSelection = ({
  singleItems,
  selectedCharacterItems,
  selectedSingleItems,
  setSelectedSingleItems,
  validation,
}: SingleItemSelectionProps) => {
  //やばいことをやっているが、結局すべて入力値に対してselectedSingleItemsを更新する
  const fixedStartList: string[] = ["最初", "最後"];

  const handleSingleItemSelect =
    (index: number) => (event: SelectChangeEvent<number>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      const selectedSingleItem = {
        id: Number(event.target.value),
        start: {},
        end: {},
      };
      newSelectedSingleItems[index] = selectedSingleItem;
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleAddClick = () => {
    const newSelectedSingleItems = selectedSingleItems
      ? [...selectedSingleItems]
      : [];
    newSelectedSingleItems.push({ start: {}, end: {} });
    setSelectedSingleItems(newSelectedSingleItems);
  };

  const handleDeleteClick = (index: number) => () => {
    const newSelectedSingleItems = [...selectedSingleItems];
    newSelectedSingleItems.splice(index, 1);
    setSelectedSingleItems(newSelectedSingleItems);
  };

  const handleLayerChange =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      newSelectedSingleItems[index].layer = Number(event.target.value);
      setSelectedSingleItems(newSelectedSingleItems);
    };

  //ラジオボタンから飛んでくるのは、あくまで"true"や"false"なので、それらの比較演算子を加える
  const handleSelectStart =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      const isFixed = event.target.value === "true";
      newSelectedSingleItems[index].isFixedStart = isFixed;
      if (isFixed) {
        newSelectedSingleItems[index].start.characterID = undefined;
        newSelectedSingleItems[index].start.adjustmentValue = undefined;
      } else if (!isFixed) {
        newSelectedSingleItems[index].start.insertPlace = undefined;
      }
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleSelectEnd =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      const isFixed = event.target.value === "true";
      newSelectedSingleItems[index].isFixedEnd = isFixed;
      if (isFixed) {
        newSelectedSingleItems[index].end.isUnique = undefined;
        newSelectedSingleItems[index].end.length = undefined;
      } else if (!isFixed) {
        newSelectedSingleItems[index].end.howManyAheads = undefined;
        newSelectedSingleItems[index].end.adjustmentValue = undefined;
      }
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleSelectLength =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      const value = event.target.value === "true";
      newSelectedSingleItems[index].end.isUnique = value;
      if (event.target.value === "true") {
        const singleItem = singleItems.find(
          (item) => item.id === newSelectedSingleItems[index].id
        );
        if (singleItem !== undefined) {
          newSelectedSingleItems[index].end.length = singleItem.length;
        }
      }
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleInputLength =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      const length = Number(event.target.value);
      newSelectedSingleItems[index].end.length = length;
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleCharacterSelect =
    (index: number) => (event: SelectChangeEvent<number>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      newSelectedSingleItems[index].start.characterID = Number(
        event.target.value
      );
      const character = selectedCharacterItems.find(
        (item) => item.id === Number(event.target.value)
      );
      newSelectedSingleItems[index].start.characterName = character?.name;
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleInputAdjustmentValueInStart =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      newSelectedSingleItems[index].start.adjustmentValue = Number(
        event.target.value
      );
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleInputAdjustmentValueInEnd =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      newSelectedSingleItems[index].end.adjustmentValue = Number(
        event.target.value
      );
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleInputHowManyAheads =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      newSelectedSingleItems[index].end.howManyAheads = Number(
        event.target.value
      );
      setSelectedSingleItems(newSelectedSingleItems);
    };

  const handleSelectFixedStart =
    (index: number) => (event: SelectChangeEvent<string>) => {
      const newSelectedSingleItems = [...selectedSingleItems];
      newSelectedSingleItems[index].start.insertPlace = event.target.value;
      if (newSelectedSingleItems[index].start.insertPlace === "最後") {
        newSelectedSingleItems[index].isFixedEnd = true;
        newSelectedSingleItems[index].end.adjustmentValue = undefined;
        newSelectedSingleItems[index].end.howManyAheads = undefined;
      }
      setSelectedSingleItems(newSelectedSingleItems);
    };

  return (
    <>
      <Typography variant="h4" component="h1" sx={{ my: 7 }}>
        Single Item Rule
      </Typography>
      <Stack spacing={2} alignItems="center">
        <Grid container>
          {selectedSingleItems &&
            selectedSingleItems.map(
              (item: SelectedSingleItem, index: number) => (
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
                          <InputLabel id="single-item">Single Item</InputLabel>
                          <Select
                            value={selectedSingleItems[index]?.id || ""}
                            onChange={handleSingleItemSelect(index)}
                            label="SingleItem"
                            labelId="single-item"
                            error={
                              validation &&
                              selectedSingleItems[index]?.id === undefined
                            }
                          >
                            {singleItems &&
                              singleItems.map((item) => (
                                <MenuItem
                                  value={item.id}
                                  key={`single-item-menu${item.id}`}
                                >
                                  {item.name}
                                </MenuItem>
                              ))}
                          </Select>
                        </FormControl>
                        {selectedSingleItems[index]?.id && (
                          <TextField
                            label="Layer"
                            type="number"
                            value={
                              selectedSingleItems[index]?.layer !== undefined &&
                              selectedSingleItems[index]?.layer !== null
                                ? (
                                    selectedSingleItems[index]?.layer as number
                                  ).toString()
                                : ""
                            }
                            onChange={handleLayerChange(index)}
                            error={
                              validation &&
                              selectedSingleItems[index]?.layer === undefined
                            }
                          />
                        )}
                      </Stack>
                    </Grid>

                    {selectedSingleItems[index]?.id && (
                      <>
                        <Grid item sx={{ margin: 2 }}>
                          <Stack direction="row" spacing={20}>
                            <Stack direction="column" spacing={3}>
                              <FormControl
                                sx={{ minWidth: 200 }}
                                error={
                                  validation &&
                                  selectedSingleItems[index].isFixedStart ===
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
                                    selectedSingleItems[index].isFixedStart ===
                                    undefined
                                      ? ""
                                      : selectedSingleItems[index].isFixedStart
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
                              {selectedSingleItems[index].isFixedStart ===
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
                                      selectedSingleItems[index].start
                                        ?.insertPlace || ""
                                    }
                                    error={
                                      validation &&
                                      selectedSingleItems[index].start
                                        ?.insertPlace === undefined
                                    }
                                  >
                                    {fixedStartList.map((item) => (
                                      <MenuItem
                                        value={item}
                                        key={`single-item-fixed-start-list-${item}`}
                                      >
                                        {item}
                                      </MenuItem>
                                    ))}
                                  </Select>
                                </FormControl>
                              )}

                              {selectedSingleItems[index].isFixedStart ===
                                false && (
                                <>
                                  <FormControl sx={{ minWidth: 200 }}>
                                    <InputLabel id="character-item">
                                      Character Item
                                    </InputLabel>
                                    <Select
                                      value={
                                        selectedSingleItems[index].start
                                          .characterID || ""
                                      }
                                      onChange={handleCharacterSelect(index)}
                                      label="Character Item"
                                      labelId="character-item"
                                      error={
                                        validation &&
                                        selectedSingleItems[index].start
                                          ?.characterID === undefined
                                      }
                                    >
                                      {selectedCharacterItems.map((item) => (
                                        // やってることは高度だが、結局選択されたアイテムをドロップダウンで選択できなくしただけ
                                        <MenuItem
                                          value={item.id}
                                          key={`single-item-character-item-list-${item.id}`}
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
                                      selectedSingleItems[index].start
                                        ?.adjustmentValue !== undefined &&
                                      selectedSingleItems[index].start
                                        ?.adjustmentValue !== null
                                        ? (
                                            selectedSingleItems[index].start
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
                                      selectedSingleItems[index].start
                                        ?.adjustmentValue === undefined
                                    }
                                  />
                                </>
                              )}
                            </Stack>
                            <Stack direction="column" spacing={3}>
                              <FormControl
                                error={
                                  validation &&
                                  selectedSingleItems[index].isFixedEnd ===
                                    undefined
                                }
                              >
                                <FormLabel id="end-label">長さ</FormLabel>
                                <RadioGroup
                                  row
                                  aria-labelledby="end-label"
                                  name="end-radio-group"
                                  onChange={handleSelectEnd(index)}
                                  value={
                                    selectedSingleItems[index].isFixedEnd ===
                                    undefined
                                      ? ""
                                      : selectedSingleItems[index].isFixedEnd
                                  }
                                >
                                  <FormControlLabel
                                    value={true}
                                    control={<Radio />}
                                    label="自由長"
                                  />
                                  <FormControlLabel
                                    value={false}
                                    control={<Radio />}
                                    label="可変長"
                                    disabled={
                                      selectedSingleItems[index].start
                                        .insertPlace === "最後"
                                    }
                                  />
                                </RadioGroup>
                              </FormControl>
                              {selectedSingleItems[index].isFixedEnd ===
                                true && (
                                <>
                                  <FormControl
                                    error={
                                      validation &&
                                      selectedSingleItems[index].end
                                        ?.isUnique === undefined
                                    }
                                  >
                                    <FormLabel id="end-label">数値</FormLabel>
                                    <RadioGroup
                                      row
                                      aria-labelledby="end-label"
                                      name="end-radio-group"
                                      onChange={handleSelectLength(index)}
                                      value={
                                        selectedSingleItems[index].end
                                          ?.isUnique === undefined
                                          ? ""
                                          : selectedSingleItems[index].end
                                              .isUnique
                                      }
                                    >
                                      <FormControlLabel
                                        value={true}
                                        control={<Radio />}
                                        label="固有値"
                                      />
                                      <FormControlLabel
                                        value={false}
                                        control={<Radio />}
                                        label="自由値"
                                      />
                                    </RadioGroup>
                                  </FormControl>
                                  {!(
                                    selectedSingleItems[index].end.isUnique ===
                                    undefined
                                  ) && (
                                    <TextField
                                      label="Length"
                                      type="number"
                                      value={
                                        selectedSingleItems[index]?.end
                                          .length !== undefined &&
                                        selectedSingleItems[index]?.end
                                          .length !== null
                                          ? (
                                              selectedSingleItems[index]?.end
                                                .length as number
                                            ).toString()
                                          : ""
                                      }
                                      disabled={
                                        selectedSingleItems[index].end.isUnique
                                      }
                                      onChange={handleInputLength(index)}
                                      error={
                                        validation &&
                                        selectedSingleItems[index].end
                                          ?.length === undefined
                                      }
                                    />
                                  )}
                                </>
                              )}

                              {selectedSingleItems[index].isFixedEnd ===
                                false && (
                                <>
                                  <FormControl sx={{ width: 210 }}>
                                    {selectedSingleItems[index]?.id && (
                                      <TextField
                                        label="先読み数"
                                        type="number"
                                        value={
                                          selectedSingleItems[index].end
                                            ?.howManyAheads !== undefined &&
                                          selectedSingleItems[index].end
                                            ?.howManyAheads !== null
                                            ? (
                                                selectedSingleItems[index].end
                                                  ?.howManyAheads as number
                                              ).toString()
                                            : ""
                                        }
                                        onChange={handleInputHowManyAheads(
                                          index
                                        )}
                                        sx={{ minWidth: 200 }}
                                        error={
                                          validation &&
                                          (selectedSingleItems[index]?.end
                                            ?.howManyAheads === undefined ||
                                            (selectedSingleItems[index]?.end
                                              ?.howManyAheads as number) < 0)
                                        }
                                      />
                                    )}
                                  </FormControl>
                                  <FormControl sx={{ width: 210 }}>
                                    {selectedSingleItems[index]?.id && (
                                      <TextField
                                        label="調整値(フレーム単位)"
                                        type="number"
                                        value={
                                          selectedSingleItems[index].end
                                            ?.adjustmentValue !== undefined &&
                                          selectedSingleItems[index].end
                                            ?.adjustmentValue !== null
                                            ? (
                                                selectedSingleItems[index].end
                                                  ?.adjustmentValue as number
                                              ).toString()
                                            : ""
                                        }
                                        onChange={handleInputAdjustmentValueInEnd(
                                          index
                                        )}
                                        sx={{ minWidth: 200 }}
                                        error={
                                          validation &&
                                          selectedSingleItems[index].end
                                            ?.adjustmentValue === undefined
                                        }
                                      />
                                    )}
                                  </FormControl>
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
            sx={{ marginTop: "0px !important", flexGrow: 0 }}
          >
            <AddCircleOutlineIcon fontSize="large" />
          </IconButton>
        </div>
      </Stack>
    </>
  );
};

export default SingleItemSelection;
