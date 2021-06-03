import { makeStyles } from "@material-ui/core/styles";

const paymentModalStyle = makeStyles((theme) => ({
    modal: {
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
    },
    paper: {
        display: "flex",
        flexDirection: "column",
        backgroundColor: theme.palette.background.paper,
        boxShadow: theme.shadows[5],
        padding: theme.spacing(2, 4, 3),
        width: 500,
        borderRadius: 10,
    },
    title: {
        alignSelf: "center",
    },
    image: {
        alignSelf: "center",
        maxHeight: 300,
    },
    totalPayment: {
        display: "inline-block",
    },
    modalAction: {
        display: "flex",
        flexDirection: "row",
        alignSelf: "flex-end",
        marginTop: theme.spacing(1),
    },
    actionButton: {
        display: "flex",
        alignItems: "center",
        margin: theme.spacing(1),
        width: 120,
    },
}));

export default paymentModalStyle;
