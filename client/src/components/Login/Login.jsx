import {
    Button,
    Card,
    CardContent,
    Grid,
    InputAdornment,
    TextField,
    Typography,
} from "@material-ui/core";
import Container from "@material-ui/core/Container";
import AccountCircle from "@material-ui/icons/AccountCircle";
import LockIcon from "@material-ui/icons/Lock";
import { useState } from "react";
import { Redirect } from "react-router";
import loginStyle from "./loginStyle";

export default function Login() {
    const classes = loginStyle();
    const [employeeId, setEmployeeId] = useState("");
    const [password, setPassword] = useState("");
    const [redirect, setRedirect] = useState(false);

    const handleLogin = async (event) => {
        event.preventDefault();

        await fetch("http://localhost:8000/api/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
            body: JSON.stringify({
                id: parseInt(employeeId),
                password: password,
            }),
        });

        setRedirect(true);
    };

    if (redirect) return <Redirect to="/dashboard" />;

    return (
        <Container>
            <Grid container>
                <Grid item xs={3}></Grid>
                <Grid item xs={6}>
                    <Card className={classes.card}>
                        <CardContent>
                            <Typography
                                className={classes.title}
                                color="textPrimary"
                                variant="h4"
                                gutterBottom
                            >
                                Login
                            </Typography>
                            <form
                                onSubmit={handleLogin}
                                noValidate
                                autoComplete="off"
                                className={classes.loginForm}
                            >
                                <TextField
                                    className={classes.input}
                                    type="number"
                                    label="ID Karyawan"
                                    InputProps={{
                                        startAdornment: (
                                            <InputAdornment position="start">
                                                <AccountCircle />
                                            </InputAdornment>
                                        ),
                                    }}
                                    onChange={(event) =>
                                        setEmployeeId(event.target.value)
                                    }
                                />
                                <TextField
                                    className={classes.input}
                                    label="Password"
                                    type="password"
                                    InputProps={{
                                        startAdornment: (
                                            <InputAdornment position="start">
                                                <LockIcon />
                                            </InputAdornment>
                                        ),
                                    }}
                                    onChange={(event) =>
                                        setPassword(event.target.value)
                                    }
                                />
                                <Button
                                    className={classes.loginButton}
                                    variant="contained"
                                    color="primary"
                                    type="submit"
                                >
                                    Login
                                </Button>
                            </form>
                        </CardContent>
                    </Card>
                </Grid>
                <Grid item xs={3}></Grid>
            </Grid>
        </Container>
    );
}
