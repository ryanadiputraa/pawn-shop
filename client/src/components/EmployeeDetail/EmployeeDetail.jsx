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
    Typography,
} from "@material-ui/core";
import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import Alert from "@material-ui/lab/Alert";
import CloseIcon from "@material-ui/icons/Close";
import { useState, useEffect } from "react";
import customerModalStyle from "../CustomerModal/customerModalStyle";

export default function EmployeeDetail({
    data,
    openModal,
    setOpenModal,
    setIsSuccess,
}) {
    const classes = customerModalStyle();
    const [firstname, setFirstname] = useState();
    const [lastname, setLastname] = useState("");
    const [gender, setGender] = useState("");
    const [birthdate, setBirthdate] = useState("");
    const [address, setAddress] = useState("");
    const [password, setPassword] = useState("");
    const [isAlert, setIsAlert] = useState(false);

    useEffect(() => {
        const abortCont = new AbortController();

        setFirstname(data.firstname);
        setLastname(data.lastname);
        setGender(data.gender);
        setBirthdate(data.birthdate);
        setAddress(data.address);
        setPassword(data.password);

        return () => abortCont.abort();
    }, [data]);

    const handleDelete = async (event) => {
        event.preventDefault();

        const res = await fetch(
            `http://localhost:8000/api/employees/${data.id}`,
            {
                method: "DELETE",
                credentials: "include",
            }
        );
        const resData = await res.json();

        if (resData.code === 200) {
            setIsAlert(false);
            setOpenModal(false);
            setIsSuccess(true);
        } else {
            setIsAlert(true);
            setIsSuccess(false);
        }
    };

    const handleUpdate = async (event) => {
        event.preventDefault();

        const res = await fetch(
            `http://localhost:8000/api/employees/${data.id}`,
            {
                method: "PUT",
                credentials: "include",
                body: JSON.stringify({
                    firstname: firstname,
                    lastname: lastname,
                    gender: gender,
                    birthdate: birthdate,
                    address: address,
                    password: password,
                }),
            }
        );
        const resData = await res.json();

        if (resData.code === 200) {
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
                    <Typography
                        className={classes.title}
                        variant="h4"
                        color="primary"
                        gutterBottom
                    >
                        Form Karyawan
                    </Typography>
                    <p>Masukan Data Karyawan</p>
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
                        noValidate
                        autoComplete="off"
                        className={classes.inputForm}
                    >
                        <TextField
                            type="text"
                            label="Nama Depan"
                            InputLabelProps={{ shrink: true }}
                            value={firstname}
                            onChange={(event) =>
                                setFirstname(event.target.value)
                            }
                        />
                        <TextField
                            type="text"
                            label="Nama Belakang"
                            InputLabelProps={{ shrink: true }}
                            value={lastname}
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
                        <FormLabel component="legend">Tanggal Lahir</FormLabel>
                        <TextField
                            type="date"
                            value={birthdate}
                            onChange={(event) =>
                                setBirthdate(event.target.value)
                            }
                        />
                        <TextField
                            type="text"
                            label="Alamat"
                            InputLabelProps={{ shrink: true }}
                            value={address}
                            onChange={(event) => setAddress(event.target.value)}
                        />
                        <TextField
                            type="text"
                            label="Password"
                            InputLabelProps={{ shrink: true }}
                            value={password}
                            onChange={(event) =>
                                setPassword(event.target.value)
                            }
                        />
                        <Button
                            className={classes.loginButton}
                            variant="contained"
                            color="secondary"
                            onClick={handleDelete}
                        >
                            Hapus
                        </Button>
                        <Button
                            className={classes.loginButton}
                            variant="contained"
                            color="primary"
                            onClick={handleUpdate}
                        >
                            Simpan
                        </Button>
                    </form>
                </div>
            </Fade>
        </Modal>
    );
}
