import {
    Button,
    Card,
    CardContent,
    Grid,
    InputAdornment,
    makeStyles,
    TextField,
    Typography,
} from "@material-ui/core";
import Container from "@material-ui/core/Container";
import AccountCircle from "@material-ui/icons/AccountCircle";
import LockIcon from "@material-ui/icons/Lock";
import { useState } from "react";

const useStyles = makeStyles({
    card: {
        minWidth: 275,
        marginTop: 80,
    },
    title: {
        textAlign: "center",
    },
    loginForm: {
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        padding: 40,
    },
    input: {
        margin: (0, 10),
        width: "100%",
        minWidth: 260,
    },
    loginButton: {
        marginTop: 20,
        alignSelf: "flex-end",
    },
});

export default function Login() {
    const classes = useStyles();
    const [employeeId, setEmployeeId] = useState("");
    const [password, setPassword] = useState("");

    const handleLogin = async (event) => {
        event.preventDefault();
        try {
            await fetch("http://localhost:8000/api/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
                body: JSON.stringify({
                    id: parseInt(employeeId),
                    password: password,
                }),
            });
        } catch (err) {
            console.log(err);
        }
    };

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
