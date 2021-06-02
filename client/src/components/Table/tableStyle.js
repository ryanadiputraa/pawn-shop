import { makeStyles } from "@material-ui/core";

const tableStyle = makeStyles((theme) => ({
    dataRow: {
        "&:hover": {
            backgroundColor: theme.palette.action.hover,
            cursor: "pointer",
        },
    },
}));

export default tableStyle;
