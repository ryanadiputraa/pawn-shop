import { makeStyles } from "@material-ui/core/styles";

const customerModalStyle = makeStyles((theme) => ({
    modal: {
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
    },
    paper: {
        backgroundColor: theme.palette.background.paper,
        border: "2px solid #000",
        boxShadow: theme.shadows[5],
        padding: theme.spacing(2, 4, 3),
        width: 500,
    },
    inputForm: {
        height: 500,
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
    },
    gender: {
        marginTop: theme.spacing(3),
    },
}));

export default customerModalStyle;
