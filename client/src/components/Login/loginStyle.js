import { makeStyles } from "@material-ui/core";

const loginStyle = makeStyles((theme) => ({
    card: {
        minWidth: 275,
        marginTop: 80,
    },
    title: {
        textAlign: "center",
        marginTop: theme.spacing(2),
        marginBottom: theme.spacing(2),
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
    role: {
        alignSelf: "flex-start",
        marginTop: theme.spacing(2),
    },
    roleSelect: {
        flexDirection: "row",
    },
    loginButton: {
        marginTop: 20,
        alignSelf: "flex-end",
    },
}));

export default loginStyle;
