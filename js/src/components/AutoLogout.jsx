import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Button, Modal } from "rsuite";
import { fetchMyInfo } from "../store/userSlice";

import axios from "axios";
import { logOut } from "../store/usedSlice";

export const expireUser = async () => {
    try {
        const response = await axios.delete("/api/logout");
        if (response.status === 200) {
            localStorage.clear();
            window.location.href = window.location.origin + "/user/expired";
        }
    } catch (error) {
        console.error("failed:", error);
    }
};
const LogoutTimer = ({ expirationTime, logout }) => {
    const formatTime = (milliseconds) => {
        if (milliseconds < 0) {
            return "0 seconds";
        }
        const seconds = Math.floor((milliseconds / 1000) % 60);
        const minutes = Math.floor((milliseconds / (1000 * 60)) % 60);
        return `${minutes}: ${seconds < 10 ? "0" : ""}${seconds} `;
    };
    const [timeLeft, setTimeLeft] = useState(
        getDateDifferenceInSeconds(expirationTime, new Date())
    );

    const [cut, setCut] = useState(false);

    useEffect(() => {
        const timer = setTimeout(
            () => {
                logout();
            },
            timeLeft > 2_000_000 ? 2_000_000 : timeLeft
        );
        return () => clearTimeout(timer);
    }, [timeLeft]);

    useEffect(() => {
        const interval = setInterval(() => {
            setTimeLeft((prevTime) => {
                if (prevTime < 30000) {
                    setCut(true);
                }
                return prevTime - 5000;
            });
            console.log(
                `Your session will expire in ${timeLeft / 1000 / 60 / 60
                } hours.`
            );
        }, 5000);
        return () => clearInterval(interval);
    }, []);

    return (
        <>
            {cut ? (
                <ExitAlertModel
                    content={`Your account will expire in ${formatTime(
                        timeLeft
                    )}.`}
                />
            ) : null}
        </>
    );
};

const ExitAlertModel = (props) => {
    const [open, setOpen] = React.useState(true);
    const handleClose = () => setOpen(false);
    return (
        <>
            <Modal open={open} onClose={handleClose}>
                <Modal.Header>
                    <Modal.Title>Accout expires</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <div>{props.content}</div>
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={handleClose} appearance="primary">
                        Ok
                    </Button>
                    <Button onClick={handleClose} appearance="subtle">
                        Cancel
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
};
const AutoLogout = () => {
    const dispatch = useDispatch();

    const { myInfo, loading, error } = useSelector((state) => state.users);

    useEffect(() => {
        dispatch(fetchMyInfo());
    }, [dispatch]);

    const logout = async () => {
        localStorage.removeItem("isLoggedIn");
        localStorage.clear();
        dispatch(logOut());
        setTimeout(async () => {
            await expireUser();
        }, 1000);
    };

    if (loading || error || !myInfo) return <></>;
    if (myInfo.InviteCode?.is_infinite) {
        return <></>;
    }
    return (
        <div>
            {localStorage.getItem("isLoggedIn") && myInfo.InviteCode.expiration_time ? (
                <LogoutTimer
                    expirationTime={new Date(myInfo.InviteCode.expiration_time)}
                    logout={logout}
                />
            ) : null}
        </div>
    );
};

export default AutoLogout;

function getDateDifferenceInSeconds(date1, date2) {
    var date1Milliseconds = date1.getTime();
    var date2Milliseconds = date2.getTime();
    return date1Milliseconds - date2Milliseconds;
}
