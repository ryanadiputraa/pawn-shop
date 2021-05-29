import { CircularProgress, Container, Typography } from "@material-ui/core";
import { useState } from "react";
import { useEffect } from "react";
import { Redirect } from "react-router";
import DataTable from "../Table/DataTable";

export default function Dashboard() {
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
            <Typography variant="h4" component="h1" align="center" gutterBottom>
                PEGADAIAN
            </Typography>
            {isLoading ? <CircularProgress /> : <DataTable data={tableData} />}
        </Container>
    );
}
