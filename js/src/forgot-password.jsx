import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function ForgotPasswordPage() {
    const [email, setEmail] = useState("");
    const [error, setError] = useState("");
    const [successMessage, setSuccessMessage] = useState("");
    const navigate = useNavigate();
    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const response = await axios.post("/api/user/forgot-password", {
                email,
            });
            if (response.status === 200) {
                localStorage.setItem("forgotEmail", email);
                setSuccessMessage(
                    "Password reset instructions sent to your email."
                );
                setTimeout(() => {
                    window.location.href =
                        window.location.origin + "/user/change-password";
                }, 1000);
            }
        } catch (error) {
            setError(error.response.data.error);
        }
    };

    return (
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
                <h3>Forgot Password</h3>
                {error && (
                    <div className="alert alert-danger" role="alert">
                        {error}
                    </div>
                )}
                {successMessage && (
                    <div className="alert alert-success" role="alert">
                        {successMessage}
                    </div>
                )}
                <div className="mb-3">
                    <label htmlFor="email" className="form-label">
                        Email:
                    </label>
                    <input
                        type="email"
                        id="email"
                        name="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        className="form-control"
                        style={{ width: "300px" }}
                        required
                    />
                </div>
                <button
                    type="submit"
                    className="btn btn-primary"
                    style={{ width: "300px" }}
                >
                    Reset Password
                </button>
                <p style={{ marginTop: "10px" }}>
                    Remembered your password? <a href="/login">Login</a>
                </p>
            </form>
        </div>
    );
}
