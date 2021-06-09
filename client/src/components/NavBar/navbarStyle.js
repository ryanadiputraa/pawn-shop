import { makeStyles } from "@material-ui/core";

const navbarStyle = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    menu: {
        flexGrow: 1,
    },
    menuItem: {
        fontSize: 24,
    },
    logout: {
        display: "flex",
        textDecoration: "none",
        alignSelf: "flex-end",
        // marginTop: theme.spacing(3),
        color: "#fff",
    },
}));

export default navbarStyle;
