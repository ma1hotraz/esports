import React, { useState } from "react";
import { Form, Button, Alert } from "react-bootstrap";
import axios from "axios";
import { BASE_URL } from "../constants";

const AddUser = () => {
    const [formData, setFormData] = useState({
        name: "",
        email: "",
        password: "",
        invite_code: "",
        is_admin: false,
    });

    const handleChange = (e) => {
        const { name, value, type, checked } = e.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]: type === "checkbox" ? checked : value,
        }));
    };

    const [error, setError] = useState(""); // State for error message

    const handleSubmit = (e) => {
        e.preventDefault();
        const formDataCopy = {
            ...formData,
            is_admin: formData.is_admin ? "true" : "false",
        };

        axios
            .post(BASE_URL + "/api/register", formDataCopy)
            .then(() => {
                window.location.href =
                    window.location.origin + "/admin/user-manage";
            })
            .catch((error) => {
                setError(error.response.data.error);
            });
    };

    return (
        <div>
            {error && <Alert variant="danger">{error}</Alert>}
            <Form onSubmit={handleSubmit}>
                <Form.Group className="mb-3" controlId="name">
                    <Form.Label>Name</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter your name"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        required
                    />
                </Form.Group>

                <Form.Group className="mb-3" controlId="email">
                    <Form.Label>Email address</Form.Label>
                    <Form.Control
                        type="email"
                        placeholder="Enter email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                </Form.Group>

                <Form.Group className="mb-3" controlId="password">
                    <Form.Label>Password</Form.Label>
                    <Form.Control
                        type="password"
                        placeholder="Password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                </Form.Group>

                <Form.Group className="mb-3" controlId="invite_code">
                    <Form.Label>Invite Code</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter invite code"
                        name="invite_code"
                        value={formData.invite_code}
                        onChange={handleChange}
                    />
                </Form.Group>

                <Form.Group className="mb-3" controlId="is_admin">
                    <Form.Check
                        type="checkbox"
                        label="Admin"
                        name="is_admin"
                        checked={formData.is_admin}
                        onChange={handleChange}
                    />
                </Form.Group>

                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </div>
    );
};

export default AddUser;
