import { useState } from "react";
import { useEffect } from "react";
import { Redirect } from "react-router";

export default function Dashboard() {
    const [redirect, setRedirect] = useState(false);

    useEffect(() => {
        const fetchCustomers = async () => {
            const res = await fetch("http://localhost:8000/api/customers", {
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            });
            const data = await res.json();

            // unauthorized user
            if (data.code === 401) setRedirect(true);
        };
        fetchCustomers();
    }, []);

    if (redirect) return <Redirect to="/" />;

    return <h1>Dashboard</h1>;
}
