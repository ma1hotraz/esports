import axios from "axios";
import React, { useState } from "react";
import toast from "react-hot-toast";
import { BASE_URL } from "./constants";

export default function ResetPasswordPage() {
    const [code, setCode] = useState("");
    const [newPassword, setNewPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [error, setError] = useState("");
    const email = localStorage.getItem("forgotEmail");
    const handleSubmit = async (event) => {
        event.preventDefault();
        if (newPassword !== confirmPassword) {
            setError("Passwords do not match");
            return;
        }
        try {
            const response = await axios.post(
                BASE_URL + "/api/user/change-password",
                {
                    code,
                    newPassword,
                    confirmPassword,
                    email,
                }
            );
            if (response.status === 200) {
                toast.success("Password reset successfully");
                setTimeout(
                    () =>
                        (window.location.href =
                            window.location.origin + "/login"),
                    1000
                );
            }
        } catch (error) {
            setError(error.response.data.error);
        }
    };

    return email ? (
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
                <h1>Reset Password</h1>
                {error && (
                    <div className="alert alert-danger" role="alert">
                        {error}
                    </div>
                )}
                <div className="mb-3">
                    <label htmlFor="code" className="form-label">
                        Reset Code:
                    </label>
                    <input
                        type="text"
                        id="code"
                        name="code"
                        value={code}
                        onChange={(e) => setCode(e.target.value)}
                        className="form-control"
                        style={{ width: "300px" }}
                        required
                    />
                </div>
                <div className="mb-3">
                    <label htmlFor="newPassword" className="form-label">
                        New Password:
                    </label>
                    <input
                        type="password"
                        id="newPassword"
                        name="newPassword"
                        value={newPassword}
                        onChange={(e) => setNewPassword(e.target.value)}
                        className="form-control"
                        style={{ width: "300px" }}
                        required
                    />
                </div>
                <div className="mb-3">
                    <label htmlFor="confirmPassword" className="form-label">
                        Confirm Password:
                    </label>
                    <input
                        type="password"
                        id="confirmPassword"
                        name="confirmPassword"
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
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
                <p className="mt-2 mb-4">
                    <a href="/user/forgot-password">
                        Resend your forgot password code?
                    </a>
                </p>
            </form>
        </div>
    ) : null;
}
