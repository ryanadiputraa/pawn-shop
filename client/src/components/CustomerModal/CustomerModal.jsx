import {
    Button,
    Collapse,
    TextField,
    FormControl,
    IconButton,
    FormLabel,
    RadioGroup,
    FormControlLabel,
    Radio,
} from "@material-ui/core";
import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import Alert from "@material-ui/lab/Alert";
import CloseIcon from "@material-ui/icons/Close";
import customerModalStyle from "./customerModalStyle";
import { useState } from "react";

export default function CustomerModal({
    openModal,
    setOpenModal,
    setIsSuccess,
}) {
    const classes = customerModalStyle();
    const [firstname, setFirstname] = useState("");
    const [lastname, setLastname] = useState("");
    const [gender, setGender] = useState("pria");
    const [contact, setContact] = useState("");
    const [loan, setLoan] = useState("");
    const [insuranceItem, setInsuranceItem] = useState("");
    const [image, setImage] = useState("");
    const [isAlert, setIsAlert] = useState(false);

    const handleSubmit = async (event) => {
        event.preventDefault();
        const formData = new FormData();

        formData.append("firstname", firstname);
        formData.append("lastname", lastname);
        formData.append("gender", gender);
        formData.append("contact", contact);
        formData.append("loan", loan);
        formData.append("insuranceItem", insuranceItem);
        formData.append("upload", image);

        const res = await fetch("http://localhost:8000/api/customers", {
            method: "POST",
            credentials: "include",
            body: formData,
        });
        const data = await res.json();

        if (data.code === 201) {
            window.location.reload();
            setIsAlert(false);
            setOpenModal(false);
            setIsSuccess(true);
        } else {
            setIsAlert(true);
            setIsSuccess(false);
        }
    };

    return (
        <Modal
            aria-labelledby="transition-modal-title"
            aria-describedby="transition-modal-description"
            className={classes.modal}
            open={openModal}
            onClose={() => setOpenModal(false)}
            closeAfterTransition
            BackdropComponent={Backdrop}
            BackdropProps={{
                timeout: 500,
            }}
        >
            <Fade in={openModal}>
                <div className={classes.paper}>
                    <h2 id="transition-modal-title">Form Gadai Barang</h2>
                    <p>Masukan Data Nasabah</p>
                    <Collapse in={isAlert} className={classes.warning}>
                        <Alert
                            severity="error"
                            action={
                                <IconButton
                                    aria-label="close"
                                    color="inherit"
                                    size="small"
                                    onClick={() => setIsAlert(false)}
                                >
                                    <CloseIcon fontSize="inherit" />
                                </IconButton>
                            }
                        >
                            Mohon Periksa Kembali data yang anda masukan!
                        </Alert>
                    </Collapse>
                    <form
                        encType="multipart/form-data"
                        onSubmit={handleSubmit}
                        noValidate
                        autoComplete="off"
                        className={classes.inputForm}
                    >
                        <TextField
                            type="text"
                            label="Nama Depan"
                            onChange={(event) =>
                                setFirstname(event.target.value)
                            }
                        />
                        <TextField
                            type="text"
                            label="Nama Belakang"
                            onChange={(event) =>
                                setLastname(event.target.value)
                            }
                        />
                        <FormControl
                            component="fieldset"
                            className={classes.gender}
                        >
                            <FormLabel component="legend">
                                Jenis Kelamin
                            </FormLabel>
                            <RadioGroup
                                className={classes.roleSelect}
                                aria-label="role"
                                name="role"
                                value={gender}
                                onChange={(event) =>
                                    setGender(event.target.value)
                                }
                            >
                                <FormControlLabel
                                    value="pria"
                                    control={<Radio />}
                                    label="Pria"
                                />
                                <FormControlLabel
                                    value="wanita"
                                    control={<Radio />}
                                    label="Wanita"
                                />
                            </RadioGroup>
                        </FormControl>
                        <TextField
                            type="text"
                            label="Kontak"
                            onChange={(event) => setContact(event.target.value)}
                        />
                        <TextField
                            type="text"
                            label="Pinjaman"
                            onChange={(event) => setLoan(event.target.value)}
                        />
                        <TextField
                            type="text"
                            label="Barang Gadai"
                            onChange={(event) =>
                                setInsuranceItem(event.target.value)
                            }
                        />
                        <Button variant="contained" component="label">
                            Foto Barang
                            <input
                                type="file"
                                hidden
                                onChange={(event) =>
                                    setImage(event.target.files[0])
                                }
                            />
                        </Button>
                        <Button
                            className={classes.loginButton}
                            variant="contained"
                            color="primary"
                            type="submit"
                        >
                            Simpan
                        </Button>
                    </form>
                </div>
            </Fade>
        </Modal>
    );
}
