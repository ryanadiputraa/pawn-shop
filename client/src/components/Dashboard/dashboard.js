import { makeStyles } from "@material-ui/core";

const dashboardStyle = makeStyles((theme) => ({
    title: {
        marginTop: theme.spacing(4),
    },
    subTitle: {
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
}));

export default dashboardStyle;
