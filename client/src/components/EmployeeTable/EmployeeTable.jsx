import {
    withStyles,
    Paper,
    Table,
    TableContainer,
    TableCell,
    TableRow,
    TableHead,
    TableBody,
} from "@material-ui/core";

const StyledTableCell = withStyles((theme) => ({
    head: {
        backgroundColor: theme.palette.common.black,
        color: theme.palette.common.white,
    },
    body: {
        fontSize: 14,
    },
}))(TableCell);

const headers = [
    "Nama",
    "Jenis Kelamin",
    "Barang Gadai",
    "Status Barang",
    "Pinjaman",
    "Total Pelunasan",
    "Kontak",
];

export default function EmployeeTable({ data }) {
    console.log(data);
    return (
        <TableContainer component={Paper}>
            <Table aria-label="data-table">
                <TableHead>
                    <TableRow>
                        {headers.map((header) =>
                            header === "Nama" ? (
                                <StyledTableCell key={header}>
                                    {header}
                                </StyledTableCell>
                            ) : (
                                <StyledTableCell key={header} align="center">
                                    {header}
                                </StyledTableCell>
                            )
                        )}
                    </TableRow>
                </TableHead>
                <TableBody>
                    {/* {data.map((d) => (
                        <TableRow key={d.customerId}>
                            <TableCell component="th" scope="row">
                                {`${d.firstname} ${d.lastname}`}
                            </TableCell>
                        </TableRow>
                    ))} */}
                </TableBody>
            </Table>
        </TableContainer>
    );
}
