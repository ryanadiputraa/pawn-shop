import {
    Button,
    Card,
    CardContent,
    FormControl,
    FormLabel,
    FormControlLabel,
    Grid,
    InputAdornment,
    RadioGroup,
    TextField,
    Typography,
} from "@material-ui/core";
import Radio from "@material-ui/core/Radio";
import Container from "@material-ui/core/Container";
import AccountCircle from "@material-ui/icons/AccountCircle";
import LockIcon from "@material-ui/icons/Lock";
import { useState } from "react";
import { Redirect } from "react-router";
import loginStyle from "./loginStyle";

export default function Login() {
    const classes = loginStyle();
    const [userId, setUserId] = useState("");
    const [password, setPassword] = useState("");
    const [redirect, setRedirect] = useState(false);
    const [role, setRole] = useState("employee");

    const handleLogin = async (event) => {
        event.preventDefault();

        const res = await fetch("http://localhost:8000/api/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
            body: JSON.stringify({
                id: parseInt(userId),
                password: password,
            }),
        });
        const data = await res.json();

        if (data.code === 202) setRedirect(true);
        else setRedirect(false);
    };

    if (redirect) return <Redirect to={"/dashboard?role=" + role} />;

    return (
        <Container>
            <Grid container>
                <Grid item xs={3}></Grid>
                <Grid item xs={6}>
                    <Card className={classes.card} elevation={3}>
                        <CardContent>
                            <Typography
                                className={classes.title}
                                color="textPrimary"
                                variant="h4"
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
                                    label={
                                        role === "employee"
                                            ? "ID Karyawan"
                                            : "ID Manager"
                                    }
                                    InputProps={{
                                        startAdornment: (
                                            <InputAdornment position="start">
                                                <AccountCircle />
                                            </InputAdornment>
                                        ),
                                    }}
                                    onChange={(event) =>
                                        setUserId(event.target.value)
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
                                <FormControl
                                    component="fieldset"
                                    className={classes.role}
                                >
                                    <FormLabel component="legend">
                                        Masuk Sebagai
                                    </FormLabel>
                                    <RadioGroup
                                        className={classes.roleSelect}
                                        aria-label="role"
                                        name="role"
                                        value={role}
                                        onChange={(event) =>
                                            setRole(event.target.value)
                                        }
                                    >
                                        <FormControlLabel
                                            value="manager"
                                            control={<Radio />}
                                            label="Manager"
                                        />
                                        <FormControlLabel
                                            value="employee"
                                            control={<Radio />}
                                            label="Karyawan"
                                        />
                                    </RadioGroup>
                                </FormControl>
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
