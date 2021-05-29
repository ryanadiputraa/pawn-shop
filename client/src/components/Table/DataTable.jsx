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

const formatCurrency = (currency) => {
    currency = currency.split("").reverse();

    let digit = 0;
    let countFormat = 1;
    currency.map((_, idx) => {
        digit++;
        if (digit === 3) {
            // add period after 3th digit
            currency.splice(idx + countFormat, 0, ".");
            digit = 0;
            countFormat++;
        }
        return null;
    });
    return currency.reverse().join("");
};

export default function DataTable({ data }) {
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
                    {data.map((d) => (
                        <TableRow key={d.customerId}>
                            <TableCell component="th" scope="row">
                                {`${d.firstname} ${d.lastname}`}
                            </TableCell>
                            <TableCell align="center">{d.gender}</TableCell>
                            <TableCell align="center">
                                {d.insuranceItem}
                            </TableCell>
                            <TableCell align="center">{d.itemStatus}</TableCell>
                            <TableCell align="center">
                                {`Rp.${formatCurrency(d.loan.toString())}`}
                            </TableCell>
                            <TableCell align="center">{`Rp.${formatCurrency(
                                (
                                    parseInt(d.loan) + parseInt(d.interest)
                                ).toString()
                            )}`}</TableCell>
                            <TableCell align="center">{d.contact}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
}
