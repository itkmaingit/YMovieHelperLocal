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

export type DynamicItem = {
  id: number;
  name: string;
};

type DynamicItemSelectionProps = {
  dynamicItems: DynamicItem[];
  selectedDynamicItems: SelectedDynamicItem[];
  setSelectedDynamicItems: (items: SelectedDynamicItem[]) => void;
  validation: boolean;
};

export type SelectedDynamicItem = {
  id?: number;
  layer?: number;
};

const DynamicItemSelection = ({
  dynamicItems,
  selectedDynamicItems,
  setSelectedDynamicItems,
  validation,
}: DynamicItemSelectionProps) => {
  // イベントハンドラは全て、新しく入力された値をselectedDynamicItemsに入れているか、要素を追加、削除しているだけ
  const handleItemSelect =
    (index: number) => (event: SelectChangeEvent<number>) => {
      const newSelectedDynamicItems = [...selectedDynamicItems];
      const selectedDynamicItem = {
        id: Number(event.target.value),
        layer: undefined,
      };
      newSelectedDynamicItems[index] = selectedDynamicItem;
      setSelectedDynamicItems(newSelectedDynamicItems);
    };

  const handleLayerChange =
    (index: number) => (event: React.ChangeEvent<HTMLInputElement>) => {
      const newSelectedDynamicItems = [...selectedDynamicItems];
      const changedDynamicItem = newSelectedDynamicItems[index];
      changedDynamicItem.layer = Number(event.target.value);
      newSelectedDynamicItems.splice(index, 1, changedDynamicItem);
      setSelectedDynamicItems(newSelectedDynamicItems);
    };

  const handleDeleteClick = (index: number) => () => {
    const newSelectedDynamicItems = [...selectedDynamicItems];
    newSelectedDynamicItems.splice(index, 1);
    setSelectedDynamicItems(newSelectedDynamicItems);
  };

  const handleAddClick = () => {
    const newSelectedDynamicItems = selectedDynamicItems
      ? [...selectedDynamicItems]
      : [];
    newSelectedDynamicItems.push({});
    setSelectedDynamicItems(newSelectedDynamicItems);
  };

  return (
    <>
      <Typography variant="h4" component="h1" sx={{ my: 7 }}>
        Dynamic Item Rule
      </Typography>
      <Stack spacing={2} alignItems="center">
        <Grid container>
          {selectedDynamicItems &&
            selectedDynamicItems.map(
              (item: SelectedDynamicItem, index: number) => (
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
                          <InputLabel id="dynamic-item">
                            Dynamic Item
                          </InputLabel>
                          <Select
                            value={selectedDynamicItems[index]?.id || ""}
                            onChange={handleItemSelect(index)}
                            label="DynamicItem"
                            labelId="dynamic-item"
                            error={
                              validation &&
                              selectedDynamicItems[index]?.id === undefined
                            }
                          >
                            {dynamicItems &&
                              dynamicItems.map((item) => (
                                // やってることは高度だが、結局選択されたアイテムをドロップダウンで選択できなくしただけ
                                <MenuItem
                                  value={item.id}
                                  key={item.id}
                                  disabled={
                                    item.id !==
                                      selectedDynamicItems[index]?.id &&
                                    selectedDynamicItems
                                      .map((i) => i.id)
                                      .includes(item.id)
                                  }
                                >
                                  {item.name}
                                </MenuItem>
                              ))}
                          </Select>
                        </FormControl>
                        {selectedDynamicItems[index].id && (
                          <TextField
                            label="Layer"
                            type="number"
                            value={
                              selectedDynamicItems[index]?.layer !==
                                undefined &&
                              selectedDynamicItems[index]?.layer !== null
                                ? (
                                    selectedDynamicItems[index]?.layer as number
                                  ).toString()
                                : ""
                            }
                            error={
                              validation &&
                              selectedDynamicItems[index]?.layer == undefined
                            }
                            onChange={handleLayerChange(index)}
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
              dynamicItems &&
              selectedDynamicItems.length === dynamicItems.length
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

export default DynamicItemSelection;
