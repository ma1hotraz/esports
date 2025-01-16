import PlusIcon from '@rsuite/icons/Plus';
import axios from "axios";
import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from 'react-redux';
import { IconButton, Input, Table } from 'rsuite';
import { fetchAdminMessages, selectAdminMessages, deleteAdminMessage, selectAdminMessagesLoading } from '../store/adminMessageSlice';

const { Column, HeaderCell, Cell } = Table;

const ActionCell = ({ rowData, onRemove, ...props }) => {
    return (
        <Cell {...props} style={{ padding: '6px', display: 'flex', gap: '4px' }}>
            <button
                onClick={() => {
                    onRemove(rowData.notification_id);
                }}
            >Remove</button>
        </Cell>
    );
};

const AdminMessageForm = () => {
    const [message, setMessage] = useState("");
    const [status, setStatus] = useState(null);
    const dispatch = useDispatch();
    const adminMessages = useSelector(selectAdminMessages);
    const loading = useSelector(selectAdminMessagesLoading);
    
    const handleSubmit = async (e) => {
        e.preventDefault();
        setStatus("Sending message...");
        if (!message) {
            setStatus("❌ Please enter a message.");
            return;
        }
        try {
            const response = await axios.post("/api/notification", {
                message
            });
            if (response.status === 200) {
                setStatus("✅ Message sent successfully!");
                setMessage("");
                dispatch(fetchAdminMessages());
            } else {
                setStatus("❌ Failed to send message.");
            }
        } catch (error) {
            setStatus("❌ An error occurred.");
        }
    };

    useEffect(() => {
        if (!!adminMessages || adminMessages.length === 0) {
            dispatch(fetchAdminMessages());
        }
    }, []);

    const handleRemove = notification_id => {
        dispatch(deleteAdminMessage(notification_id));
    };


    return (
        <div>
            <div className="admin-message-form">
                <h2>Send a Message to Users</h2>
                <form onSubmit={handleSubmit}>
                    <div style={{ marginTop: "1rem" }}>
                        <Input onChange={(value) => setMessage(value)} required value={message}
                            as="textarea" rows={3} placeholder="Enter your message here..." />
                    </div>

                    <div style={{ marginTop: "1rem" }}>
                        <IconButton onClick={handleSubmit} icon={<PlusIcon />}>Send Message</IconButton>
                    </div>
                </form>
                {status && <strong>{status}</strong>}
            </div>
            {loading && <div>Processing...</div>}
            <Table style={{ marginTop: "1rem" }}
                height={700}
                data={adminMessages}
            >
                <Column width={60} align="center" fixed>
                    <HeaderCell>Id</HeaderCell>
                    <Cell dataKey="notification_id" />
                </Column>

                <Column width={200}>
                    <HeaderCell>Created At</HeaderCell>
                    <Cell style={{ padding: '6px' }}>
                        {rowData => (
                            new Date(rowData.CreatedAt).toLocaleString()
                        )}
                    </Cell>
                </Column>

                <Column width={400}>
                    <HeaderCell>Content</HeaderCell>
                    <Cell dataKey="message" />
                </Column>

                <Column width={100}>
                    <HeaderCell>Action</HeaderCell>
                    <ActionCell dataKey="notification_id" onRemove={handleRemove} />
                </Column>
            </Table>
        </div>

    );
};

export default AdminMessageForm;

