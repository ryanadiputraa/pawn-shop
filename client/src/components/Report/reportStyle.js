import { makeStyles } from "@material-ui/core";

const reportStyle = makeStyles((theme) => ({
    root: {
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        marginTop: theme.spacing(5),
    },
    print: {
        alignSelf: "flex-start",
        position: "relative",
        left: 170,
        marginBottom: theme.spacing(2),
    },
    paper: {
        height: 800,
        width: 1200,
        padding: theme.spacing(4),
        display: "flex",
        flexDirection: "column",
    },
    center: {
        alignSelf: "center",
    },
    right: {
        alignSelf: "flex-end",
    },
    dataTable: {
        marginTop: theme.spacing(4),
        marginBottom: theme.spacing(4),
    },
    dataRow: {
        "&:nth-of-type(even)": {
            backgroundColor: theme.palette.action.hover,
        },
    },
    financial: {
        display: "flex",
        "& > *": {
            margin: theme.spacing(2),
        },
    },
    total: {
        borderBottom: "1px solid #000",
    },
    deficit: {
        marginRight: theme.spacing(2),
    },
}));

export default reportStyle;
