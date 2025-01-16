import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useLocation } from "react-router-dom";
import { Button, List, Modal } from "rsuite";
import { clearMessages, dismissAllMyMessages, fetchMyMessages, selectMessages } from "../store/userMessageSlice";

const NotificationBar = () => {
    const dispatch = useDispatch();
    const messages = useSelector((selectMessages));
    const location = useLocation();

    const handleClose = () => {
        dispatch(clearMessages());
    };
    useEffect(() => {
        if (!!messages || messages.length === 0) {
            dispatch(fetchMyMessages());
        }
    }, []);

    const handleDismiss = () => {
        dispatch(dismissAllMyMessages());
    }

    return (messages && messages.length > 0 && !location.pathname.includes("/admin/")) ? <div>
        <Modal
            open={true}
            onClose={handleClose}
            size="md"
        >
            <Modal.Header>
                <Modal.Title><strong>Messages from Esport Differences</strong></Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <List>
                    {
                        messages.map((message) => {
                            return <List.Item key={message.notification_id}>
                                🌟 <strong style={{ whiteSpace: "pre" }}>[{new Date(message.CreatedAt).toLocaleString()}] {"\n"} {message.message}</strong>
                            </List.Item>
                        })
                    }
                </List>
            </Modal.Body>
            <Modal.Footer>
                <Button onClick={handleDismiss} appearance="primary">
                    Got that!
                </Button>
            </Modal.Footer>
        </Modal>
    </div> : null;
};

export default NotificationBar;