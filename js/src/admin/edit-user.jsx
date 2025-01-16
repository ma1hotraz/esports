import axios from "axios";
import React, { useEffect, useState } from "react";
import { Alert, Button, Form } from "react-bootstrap";
import { useSelector } from "react-redux";
import { useNavigate, useParams } from "react-router-dom";
import { DatePicker } from "rsuite";
import { BASE_URL } from "../constants";

const EditUser = () => {
    const { userId } = useParams(); // Get userId from route parameters
    const {
        user,
        loading,
        error: fetchError,
    } = useSelector((state) => state.users);
    const navigate = useNavigate();
    useEffect(() => {
        if (user) {
            setFormData({
                name: user.name,
                email: user.email,
                password: "",
                is_admin: user.is_admin,
                is_changingPassword: false,
                is_infinite: user.InviteCode.is_infinite,
                expiration_time: new Date(user.InviteCode.expiration_time),
            });
        }
    }, [user]);
    const [error, setError] = useState("");

    const [formData, setFormData] = useState({
        name: "",
        email: "",
        password: "",
        is_admin: false,
        is_changingPassword: false,
        expiration_time: new Date(),
        is_infinite: false,
    });
    const handleChange = (e) => {
        const { name, value, type, checked } = e.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]: type === "checkbox" ? checked : value,
        }));
    };

    const handleSubmit = (e) => {
        e.preventDefault();

        if (error !== "") {
            return;
        }
        // return
        axios
            .post(BASE_URL + "/api/user/edit/" + userId, formData)
            .then(() => {
                navigate("/admin/user-manage");
            })
            .catch((error) => {
                setError(error.response.data.error);
            });
    };
    if (loading) return <div>Loading...</div>;
    if (fetchError) return <div>Error: {fetchError}</div>;
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
                        required
                        readOnly
                    />
                </Form.Group>
                <Form.Group className="mb-3" controlId="password">
                    <Form.Check
                        type="checkbox"
                        label="Change password?"
                        name="is_changingPassword"
                        checked={formData.is_changingPassword}
                        onChange={handleChange}
                    />
                    {formData.is_changingPassword ? (
                        <>
                            <Form.Label> Password</Form.Label>
                            <Form.Control
                                type="text"
                                placeholder="Password"
                                name="password"
                                value={formData.password}
                                onChange={handleChange}
                            />
                        </>
                    ) : null}
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
                <Form.Group className="mb-3" controlId="expiration_time">
                    <Form.Check
                        type="checkbox"
                        label="Infinite access?"
                        name="is_infinite"
                        checked={formData.is_infinite}
                        onChange={handleChange}
                    />
                    {!formData.is_infinite ? (
                        <>
                            <Form.Label> Expires at?</Form.Label>
                            <div>
                                <DatePicker
                                    format="yyyy-MM-dd HH:mm"
                                    value={formData.expiration_time}
                                    onChange={(newDate) => {
                                        if (!newDate) {
                                            setFormData((prevState) => ({
                                                ...prevState,
                                                expiration_time: new Date(),
                                            }));
                                            return;
                                        }
                                        if (
                                            newDate.getTime() <
                                            new Date().getTime()
                                        ) {
                                            setError(
                                                "The time was set in the past!"
                                            );
                                            return;
                                        }
                                        setError("");
                                        setFormData((prevState) => ({
                                            ...prevState,
                                            expiration_time: newDate,
                                        }));
                                    }}
                                />
                            </div>
                        </>
                    ) : null}
                </Form.Group>
               
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </div>
    );
};

export default EditUser;
