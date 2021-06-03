import {
    CircularProgress,
    Container,
    InputAdornment,
    TextField,
    Typography,
    Collapse,
    IconButton,
} from "@material-ui/core";
import { useState } from "react";
import { useEffect } from "react";
import { Redirect } from "react-router";
import { Link } from "react-router-dom";
import DataTable from "../Table/DataTable";
import dashboardStyle from "./dashboard";
import SearchIcon from "@material-ui/icons/Search";
import EmployeeTable from "../EmployeeTable/EmployeeTable";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import Alert from "@material-ui/lab/Alert";
import CloseIcon from "@material-ui/icons/Close";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import CustomerModal from "../CustomerModal/CustomerModal";
import EmployeeModal from "../EmployeeModal/EmployeeModal";

export default function Dashboard(props) {
    const classes = dashboardStyle();
    const [redirect, setRedirect] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [tableData, setTableData] = useState([]);
    const [isNotFound, setIsNotFound] = useState(false);
    const [isSuccess, setIsSuccess] = useState(false);
    const [openModal, setOpenModal] = useState(false);
    const [isPaymentSuccess, setIsPaymentSuccess] = useState(false);
    const [openEmployeeModal, setOpenEmployeeModal] = useState(false);

    const params = new URLSearchParams(props.location.search);
    const role = params.get("role");

    const fetchTableData = async (role) => {
        setIsLoading(true);

        if (role === "manager") {
            const res = await fetch(`http://localhost:8000/api/employees`, {
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            });
            const data = await res.json();
            setTableData(data);

            if (data.code === 401) setRedirect(true);
        } else if (role === "employee") {
            const res = await fetch(`http://localhost:8000/api/customers`, {
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            });
            const data = await res.json();

            setTableData(data);
            if (data.code === 401) setRedirect(true);
        }

        setIsLoading(false);
        setIsNotFound(false);
    };

    const handleSearch = async (event) => {
        event.preventDefault();
        const name = event.target.value;
        if (name === "") {
            fetchTableData(role);
            return;
        }

        setIsLoading(true);
        let endpoint;
        if (role === "employee") {
            endpoint = "customers";
        } else if (role === "manager") {
            endpoint = "employees";
        }
        const res = await fetch(
            `http://localhost:8000/api/${endpoint}?name=${name}`,
            {
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            }
        );
        const data = await res.json();
        setIsLoading(false);

        if (!data.code) {
            setTableData(data);
            setIsNotFound(false);
            return;
        }
        setIsNotFound(true);
    };

    useEffect(() => {
        const abortCont = new AbortController();
        fetchTableData(role);

        return () => abortCont.abort();
    }, [role]);

    if (redirect) return <Redirect to="/" />;

    const handleLogout = async () => {
        await fetch("http://localhost:8000/api/logout", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
        });
    };

    const renderDashboard = (tableData) => {
        if (role === "manager")
            return (
                <EmployeeTable data={tableData} setIsSuccess={setIsSuccess} />
            );
        else if (role === "employee")
            return (
                <DataTable
                    data={tableData}
                    setIsPaymentSuccess={setIsPaymentSuccess}
                />
            );
        else return <Redirect to="/" />;
    };

    return (
        <Container className={classes.container}>
            <Link className={classes.logout} to="/" onClick={handleLogout}>
                Logout <ExitToAppIcon />
            </Link>
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
                {role === "employee" ? (
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
                    onChange={(event) => handleSearch(event)}
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
            <Collapse in={isNotFound} className={classes.warning}>
                <Alert
                    severity="error"
                    action={
                        <IconButton
                            aria-label="close"
                            color="inherit"
                            size="small"
                            onClick={() => setIsNotFound(false)}
                        >
                            <CloseIcon fontSize="inherit" />
                        </IconButton>
                    }
                >
                    Tidak ada nama yang sesuai ditemukan!
                </Alert>
            </Collapse>
            <Collapse in={isSuccess} className={classes.warning}>
                <Alert
                    severity="success"
                    action={
                        <IconButton
                            aria-label="close"
                            color="inherit"
                            size="small"
                            onClick={() => setIsSuccess(false)}
                        >
                            <CloseIcon fontSize="inherit" />
                        </IconButton>
                    }
                >
                    Data berhasil disimpan!
                </Alert>
            </Collapse>
            <Collapse in={isPaymentSuccess} className={classes.warning}>
                <Alert
                    severity="success"
                    action={
                        <IconButton
                            aria-label="close"
                            color="inherit"
                            size="small"
                            onClick={() => setIsPaymentSuccess(false)}
                        >
                            <CloseIcon fontSize="inherit" />
                        </IconButton>
                    }
                >
                    Proses Pembayaran Berhasil!
                </Alert>
            </Collapse>
            {isLoading ? <CircularProgress /> : renderDashboard(tableData)}
            {role === "employee" ? (
                <Fab
                    className={classes.fabAdd}
                    onClick={() => setOpenModal(true)}
                    color="primary"
                    aria-label="add"
                >
                    <AddIcon />
                </Fab>
            ) : (
                <Fab
                    className={classes.fabAdd}
                    onClick={() => setOpenEmployeeModal(true)}
                    color="primary"
                    aria-label="add"
                >
                    <AddIcon />
                </Fab>
            )}
            <CustomerModal
                openModal={openModal}
                setOpenModal={setOpenModal}
                setIsSuccess={setIsSuccess}
            />
            <EmployeeModal
                openModal={openEmployeeModal}
                setOpenModal={setOpenEmployeeModal}
                setIsSuccess={setIsSuccess}
            />
        </Container>
    );
}
