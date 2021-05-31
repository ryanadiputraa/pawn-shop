import { makeStyles } from "@material-ui/core";

const loginStyle = makeStyles({
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

export default loginStyle;
