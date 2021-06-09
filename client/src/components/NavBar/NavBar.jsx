import React from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Button from "@material-ui/core/Button";
import navbarStyle from "./navbarStyle";
import { Container } from "@material-ui/core";
import { Link } from "react-router-dom";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";

export default function NavBar({ role }) {
    const classes = navbarStyle();

    const handleLogout = async () => {
        await fetch("http://localhost:8000/api/logout", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
        });
    };

    return (
        <div className={classes.root}>
            <AppBar position="static">
                <Container>
                    <Toolbar>
                        <div className={classes.menu}>
                            <Button
                                className={classes.menuItem}
                                color="inherit"
                            >
                                <Link
                                    className={classes.logout}
                                    to={`/dashboard?role=${role}`}
                                >
                                    Dashboard
                                </Link>
                            </Button>
                            <Button
                                className={classes.menuItem}
                                color="inherit"
                            >
                                <Link className={classes.logout} to="report">
                                    LAPORAN
                                </Link>
                            </Button>
                        </div>
                        <Button color="inherit">
                            <Link
                                className={classes.logout}
                                to="/"
                                onClick={handleLogout}
                            >
                                Logout <ExitToAppIcon />
                            </Link>
                        </Button>
                    </Toolbar>
                </Container>
            </AppBar>
        </div>
    );
}
