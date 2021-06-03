import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import paymentModalStyle from "./paymentModalStyle";
import { Button, Typography } from "@material-ui/core";
import AttachMoneyIcon from "@material-ui/icons/AttachMoney";
import CancelIcon from "@material-ui/icons/Cancel";

export default function PaymentModal({
    openPayment,
    setOpenPayment,
    data,
    setIsPaymentSuccess,
}) {
    const classes = paymentModalStyle();

    const handlePayment = async () => {
        const res = await fetch(
            `http://localhost:8000/api/customers/${data.customerId}`,
            {
                method: "PUT",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            }
        );
        const resData = await res.json();

        if (resData.code === 200) {
            setIsPaymentSuccess(true);
            setOpenPayment(false);
        }
    };

    const formatCurrency = (currency) => {
        currency = currency.split("").reverse();

        let digit = 0;
        let countFormat = 1;
        currency.map((_, idx) => {
            digit++;
            if (digit === 3) {
                // add period after 3th digit
                currency.splice(idx + countFormat, 0, ".");
                digit = 0;
                countFormat++;
            }
            return null;
        });
        return currency.reverse().join("");
    };

    return (
        <Modal
            aria-labelledby="transition-modal-title"
            aria-describedby="transition-modal-description"
            className={classes.modal}
            open={openPayment}
            onClose={() => setOpenPayment(false)}
            closeAfterTransition
            BackdropComponent={Backdrop}
            BackdropProps={{
                timeout: 500,
            }}
        >
            <Fade in={openPayment}>
                <div className={classes.paper}>
                    <Typography
                        className={classes.title}
                        variant="h4"
                        color="primary"
                        gutterBottom
                    >
                        Pembayaran Pinjaman
                    </Typography>
                    <a
                        href={data.image}
                        className={classes.image}
                        target="_blank"
                        rel="noreferrer"
                    >
                        <img src={data.image} alt="customer insurance item" />
                    </a>
                    <Typography variant="h6" component="p">
                        Barang Jaminan : {data.insuranceItem}
                    </Typography>
                    <Typography variant="h6" component="p">
                        Nama Nasabah : {`${data.firstname} ${data.lastname}`}
                    </Typography>
                    <Typography variant="h6" component="p">
                        Jenis Kelamin : {data.gender}
                    </Typography>
                    <Typography variant="h6" component="p">
                        Kontak : {data.contact}
                    </Typography>
                    <Typography variant="h6" component="p">
                        Total Pelunasan Pinjaman :
                        <Typography
                            className={classes.totalPayment}
                            variant="h6"
                            component="p"
                            color="error"
                        >
                            Rp.
                            {formatCurrency(
                                (data.loan + data.interest).toString()
                            )}
                        </Typography>
                    </Typography>
                    <div className={classes.modalAction}>
                        <Button
                            variant="contained"
                            color="secondary"
                            className={classes.actionButton}
                            endIcon={<CancelIcon />}
                            onClick={() => setOpenPayment(false)}
                        >
                            Batal
                        </Button>
                        <Button
                            variant="contained"
                            color="primary"
                            className={classes.actionButton}
                            endIcon={<AttachMoneyIcon />}
                            onClick={handlePayment}
                        >
                            Proses
                        </Button>
                    </div>
                </div>
            </Fade>
        </Modal>
    );
}
