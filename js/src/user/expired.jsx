import axios from "axios";
import React, { useState } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

const ExpiredPage = () => {
    const [inviteCode, setInviteCode] = useState("");
    const [email, setEmail] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();
    const handleInviteCodeChange = (event) => {
        setInviteCode(event.target.value);
    };

    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        const payload = {
            email,
            invite_code: inviteCode,
        };

        try {
            await axios.post("/api/user/extend-time", payload);
            toast.success("Success");
            setTimeout(() => {
                navigate("/login");
            }, 2000);
        } catch (e) {
            setError("Invite code is not valid.");
            console.error(e);
        }
    };

    return (
        <div>
            <div
                className="d-flex justify-content-center align-items-center"
                style={{
                    minHeight: "100vh",
                    backgroundColor: "#1675e0",
                }}
            >
                <form
                    onSubmit={handleSubmit}
                    style={{
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                        padding: "20px",
                        border: "1px solid #ddd",
                        borderRadius: "4px",
                        margin: "0 auto",
                        width: "400px",
                        backgroundColor: "#fff",
                        boxShadow: "0px 0px 10px rgba(0, 0, 0, 0.1)",
                    }}
                >
                    <h2 style={{ margin: "10px" }}> Account expired</h2>

                    {error && (
                        <div className="alert alert-danger" role="alert">
                            {error}
                        </div>
                    )}
                    <div className="mb-3">
                        <label htmlFor="email" className="form-label">
                            Enter your email:
                        </label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            value={email}
                            onChange={handleEmailChange}
                            className="form-control"
                            style={{ width: "300px" }}
                        />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="name" className="form-label">
                            Enter invite code to extend expiration time:
                        </label>
                        <input
                            type="text"
                            id="name"
                            name="name"
                            value={inviteCode}
                            onChange={handleInviteCodeChange}
                            className="form-control"
                            style={{ width: "300px" }}
                        />
                    </div>
                    <button
                        type="submit"
                        className="btn btn-primary"
                        style={{ width: "300px" }}
                    >
                        Extend
                    </button>
                    <div style={{ margin: "10px" }}>
                        <p>
                            If you do not have a valid invite code, please
                            contact admin via{" "}
                            <a href={`mailto:${email}`}>
                                admin@esportsdifference.com{" "}
                            </a>{" "}
                            or discord: Railcats to extend your account
                            lifetime.
                        </p>
                    </div>

                    <p style={{ marginTop: "10px" }}>
                        Already extended? <a href="/login">Login</a>
                    </p>
                </form>
            </div>
        </div>
    );
};

export default ExpiredPage;
