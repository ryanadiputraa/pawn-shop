import { makeStyles } from "@material-ui/core";

const dashboardStyle = makeStyles((theme) => ({
    container: {
        display: "flex",
        flexDirection: "column",
    },
    logout: {
        display: "flex",
        textDecoration: "none",
        alignSelf: "flex-end",
        marginTop: theme.spacing(3),
    },
    title: {
        marginTop: theme.spacing(4),
    },
    subTitle: {
        marginBottom: theme.spacing(2),
    },
    warning: {
        marginBottom: theme.spacing(2),
    },
    tableTitle: {
        marginBottom: theme.spacing(4),
        display: "flex",
        justifyContent: "space-between",
    },
    dataTitle: {
        position: "relative",
        top: "15px",
    },
    fabAdd: {
        position: "absolute",
        bottom: "8%",
        right: "3%",
    },
}));

export default dashboardStyle;
