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

import MultipleItemRecord, {
  MultipleItem,
} from "@/components/ShowItems/MultipleItemRecord";

type MultipleItemTableProps = {
  items: MultipleItem[];
};

const MultipleItemTable = ({ items }: MultipleItemTableProps) => {
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
              <Typography>アイテム数</Typography>
            </TableCell>
            <TableCell width="10%" align="center">
              <Typography></Typography>
            </TableCell>
            <TableCell width="30%" align="center">
              <Typography>Actions</Typography>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {items &&
            items.map((item) => (
              <MultipleItemRecord {...item} key={item.id}></MultipleItemRecord>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default MultipleItemTable;
