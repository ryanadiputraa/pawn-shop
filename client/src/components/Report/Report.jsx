import { useEffect, useState } from "react";
import {
    Button,
    Paper,
    Table,
    TableContainer,
    TableHead,
    TableCell,
    withStyles,
    Typography,
    TableRow,
    TableBody,
} from "@material-ui/core";
import { useRef } from "react";
import reportStyle from "./reportStyle";
import ReactToPdf from "react-to-pdf";

export default function Report() {
    const classes = reportStyle();
    const [currentDate, setCurrentDate] = useState("");
    const [tableData, setTableData] = useState([]);
    const [totalLoans, setTotalLoans] = useState(0);
    const [payedLoans, setPayedLoans] = useState(0);
    const [unpayedLoans, setUnpayedLoans] = useState(0);
    const printOutput = useRef(null);
    const options = {
        orientation: "landscape",
        unit: "in",
        format: [10.5, 10],
    };

    useEffect(() => {
        const abortCont = new AbortController();

        getCurrentDate();

        const fetchData = async () => {
            const res = await fetch(`http://localhost:8000/api/customers`, {
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            });
            const data = await res.json();

            setTableData(data);
        };

        const getFinancialStatements = async () => {
            const res = await fetch(
                `http://localhost:8000/api/customers/financial`,
                {
                    headers: { "Content-Type": "application/json" },
                    credentials: "include",
                }
            );
            const data = await res.json();
            setTotalLoans(data.totalLoans);
            setPayedLoans(data.loanStatus[0].total);
            setUnpayedLoans(data.loanStatus[1].total);
        };

        fetchData();
        getFinancialStatements();

        return () => abortCont.abort();
    }, []);

    const getCurrentDate = () => {
        const today = new Date();
        const date = `${today.getDate()}/${today.getMonth()}/${today.getFullYear()}`;
        const time = `${today.getHours()}:${today.getMinutes()}`;
        setCurrentDate(date + " " + time);
    };

    const StyledTableCell = withStyles((theme) => ({
        head: {
            backgroundColor: theme.palette.common.black,
            color: theme.palette.common.white,
        },
        body: {
            fontSize: 14,
        },
    }))(TableCell);

    const formatCurrency = (currency) => {
        currency = currency.split("").reverse();

        let digit = 0;
        let countFormat = 1;
        currency.map((e, idx) => {
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

    const headers = [
        "Nama",
        "Jenis Kelamin",
        "Barang Gadai",
        "Status Barang",
        "Pinjaman",
        "Total Pelunasan",
        "Kontak",
    ];

    return (
        <div className={classes.root}>
            <ReactToPdf
                targetRef={printOutput}
                filename="test.pdf"
                options={options}
                scale={0.8}
            >
                {({ toPdf }) => (
                    <Button
                        className={classes.print}
                        variant="contained"
                        color="primary"
                        onClick={toPdf}
                    >
                        Cetak
                    </Button>
                )}
            </ReactToPdf>
            <Paper elevation={0} className={classes.paper} ref={printOutput}>
                <Typography
                    variant="h3"
                    color="primary"
                    className={classes.center}
                >
                    PEGADAIAN
                </Typography>
                <Typography variant="h5" className={classes.center}>
                    Laporan Data Nasabah Pegadaian
                </Typography>
                <Typography variant="h6" className={classes.center}>
                    {currentDate}
                </Typography>
                <TableContainer className={classes.dataTable} component={Paper}>
                    <Table>
                        <TableHead>
                            <TableRow>
                                {headers.map((header) =>
                                    header === "Nama" ? (
                                        <StyledTableCell key={header}>
                                            {header}
                                        </StyledTableCell>
                                    ) : (
                                        <StyledTableCell
                                            key={header}
                                            align="center"
                                        >
                                            {header}
                                        </StyledTableCell>
                                    )
                                )}
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {tableData.map((data) => (
                                <TableRow
                                    className={classes.dataRow}
                                    key={data.customerId}
                                >
                                    <TableCell>{`${data.firstname} ${data.lastname}`}</TableCell>
                                    <TableCell align="center">
                                        {data.gender}
                                    </TableCell>
                                    <TableCell align="center">
                                        {data.insuranceItem}
                                    </TableCell>
                                    <TableCell align="center">
                                        {data.itemStatus}
                                    </TableCell>
                                    <TableCell align="center">
                                        {`Rp.${formatCurrency(
                                            data.loan.toString()
                                        )}`}
                                    </TableCell>
                                    <TableCell align="center">
                                        {`Rp.${formatCurrency(
                                            (
                                                parseInt(data.loan) +
                                                parseInt(data.interest)
                                            ).toString()
                                        )}`}
                                    </TableCell>
                                    <TableCell align="center">
                                        {data.contact}
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                <div className={`${classes.right} ${classes.financial}`}>
                    <div className="status">
                        <Typography variant="h6">Status</Typography>
                        <Typography variant="h6">Total Pinjaman</Typography>
                        <Typography variant="h6">
                            Belum Dibayar (+bunga10%)
                        </Typography>
                        <Typography variant="h6">Lunas (+bunga10%)</Typography>
                    </div>
                    <div className={classes.total}>
                        <Typography variant="h6">Total</Typography>
                        <Typography variant="h6">
                            Rp.{formatCurrency(totalLoans.toString())}
                        </Typography>
                        <Typography variant="h6" color="secondary">
                            Rp.{formatCurrency(unpayedLoans.toString())}
                        </Typography>
                        <Typography variant="h6" color="primary">
                            Rp.{formatCurrency(payedLoans.toString())}
                        </Typography>
                    </div>
                </div>
                {payedLoans - totalLoans < 0 ? (
                    <Typography
                        className={`${classes.right} ${classes.deficit}`}
                        variant="h6"
                        color="secondary"
                    >
                        - Rp.
                        {formatCurrency(
                            (payedLoans - totalLoans).toString().substring(1)
                        )}
                    </Typography>
                ) : (
                    <Typography
                        className={`${classes.right} ${classes.deficit}`}
                        variant="h6"
                        color="primary"
                    >
                        Rp.
                        {formatCurrency((payedLoans - totalLoans).toString())}
                    </Typography>
                )}
            </Paper>
        </div>
    );
}
