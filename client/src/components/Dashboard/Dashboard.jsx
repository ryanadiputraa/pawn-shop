import {
    CircularProgress,
    Container,
    InputAdornment,
    TextField,
    Typography,
} from "@material-ui/core";
import { useState } from "react";
import { useEffect } from "react";
import { Redirect } from "react-router";
import DataTable from "../Table/DataTable";
import dashboardStyle from "./dashboard";
import SearchIcon from "@material-ui/icons/Search";

export default function Dashboard() {
    const classes = dashboardStyle();
    const [redirect, setRedirect] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [tableData, setTableData] = useState([]);

    useEffect(() => {
        const fetchCustomers = async () => {
            setIsLoading(true);
            const res = await fetch("http://localhost:8000/api/customers", {
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            });
            const data = await res.json();
            setTableData(data);

            // unauthorized user
            if (data.code === 401) setRedirect(true);

            setIsLoading(false);
        };
        fetchCustomers();
    }, []);

    if (redirect) return <Redirect to="/" />;

    return (
        <Container>
            <Typography
                className={classes.title}
                variant="h3"
                component="h1"
                align="center"
                color="primary"
                gutterBottom
            >
                PEGADAIAN
            </Typography>
            <Typography
                className={classes.subTitle}
                variant="h5"
                component="h2"
                align="center"
                gutterBottom
            >
                Sistem Informasi Pegadaian
            </Typography>
            <div className={classes.tableTitle}>
                {tableData.length && tableData[0].customerId ? (
                    <Typography
                        className={classes.dataTitle}
                        variant="h6"
                        component="h5"
                        align="left"
                    >
                        Data Nasabah
                    </Typography>
                ) : (
                    <Typography
                        className={classes.dataTitle}
                        variant="h6"
                        component="h5"
                        align="left"
                    >
                        Data Karyawan
                    </Typography>
                )}
                <TextField
                    className={classes.search}
                    label="Cari..."
                    type="text"
                    InputProps={{
                        endAdornment: (
                            <InputAdornment position="end">
                                <SearchIcon />
                            </InputAdornment>
                        ),
                    }}
                />
            </div>
            {isLoading ? <CircularProgress /> : <DataTable data={tableData} />}
        </Container>
    );
}
