import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";

import SingleItemRecord, {
  SingleItem,
} from "@/components/ShowItems/SingleItemRecord";

type SingleItemTableProps = {
  items: SingleItem[];
};

const SingleItemTable = ({ items }: SingleItemTableProps) => {
  return (
    <TableContainer
      component={Paper}
      sx={{ marginBottom: "20px", borderRadius: "15px" }}
    >
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell width="20%" align="center">
              <Typography>名前</Typography>
            </TableCell>
            <TableCell width="40%" align="center">
              <Typography>アイテムの属性</Typography>
            </TableCell>
            <TableCell width="10%" align="center">
              <Typography>長さ</Typography>
            </TableCell>
            <TableCell width="30%" align="center">
              <Typography>Actions</Typography>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {items &&
            items.map((item) => (
              <SingleItemRecord {...item} key={item.id}></SingleItemRecord>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default SingleItemTable;
