import {
  Card,
  CardActions,
  CardContent,
  Grid,
  Stack,
  TextField,
  Typography,
} from "@mui/material";

export type SingleItemDescriptionProps = {
  ID: string;
  ItemType: string;
  Descriptions: string[];
  Name: string;
  Length: number;
};

type RequiredData = {
  singleItems: SingleItemDescriptionProps[];
  setSingleItems: React.Dispatch<
    React.SetStateAction<SingleItemDescriptionProps[]>
  >;
  index: number;
};

type Props = SingleItemDescriptionProps & RequiredData;

const SingleItemDescription = ({
  index,
  ItemType,
  Descriptions,
  singleItems,
  setSingleItems,
}: Props) => {
  const handleChangeName =
    (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => {
      const newSingleItems = [...singleItems];
      newSingleItems[index].Name = e.target.value;
      setSingleItems(newSingleItems);
    };
  return (
    <>
      <Grid item xs={12} md={6}>
        <Card
          sx={{
            boxShadow: "1px 1px 10px rgba(0,0,0,0.15)",
            transition: "0.3s",
            borderRadius: "15px",
          }}
        >
          <CardContent>
            <Typography variant="h6" component="div" gutterBottom>
              {ItemType}
            </Typography>
            <Stack>
              {Descriptions.map((description, index) => (
                <Typography key={`${description}-${index}`}>
                  {description}
                </Typography>
              ))}
            </Stack>
          </CardContent>
          <CardActions>
            <TextField
              error={singleItems[index].Name?.length > 20}
              fullWidth
              label={`${ItemType}の名称`}
              value={singleItems[index].Name || ""}
              onChange={handleChangeName(index)}
              variant="outlined"
              helperText={
                singleItems[index].Name?.length > 20 && "名前は20文字以内です！"
              }
            />
          </CardActions>
        </Card>
      </Grid>
    </>
  );
};

export default SingleItemDescription;
