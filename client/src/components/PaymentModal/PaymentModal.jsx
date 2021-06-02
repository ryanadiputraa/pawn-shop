import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import { useState } from "react";
import paymentModalStyle from "./paymentModalStyle";

export default function PaymentModal({ openPayment, setOpenPayment }) {
    const classes = paymentModalStyle();

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
                    <h2 id="transition-modal-title">Pembayaran</h2>
                </div>
            </Fade>
        </Modal>
    );
}
