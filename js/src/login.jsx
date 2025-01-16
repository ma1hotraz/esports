import axios from "axios";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { BASE_URL } from "./constants";

export default function LoginPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();
    const handleSubmit = async (event) => {
        event.preventDefault();

        try {
            const response = await axios.post(BASE_URL + "/api/login", {
                email,
                password,
            });
            if (response.status === 200) {
                localStorage.setItem("isLoggedIn", true);
                localStorage.setItem("isAdmin", response.data.is_admin);
                localStorage.setItem("email", response.data.email);
                localStorage.setItem("userId", response.data.userId);
                localStorage.setItem("user", response.data.user);
                
                if (!response.data.is_admin) {
                    window.location.href = window.location.origin + "/cod";
                } else {
                    window.location.href =
                        window.location.origin + "/admin/user-manage";
                }
            }
        } catch (error) {
            console.error("Login failed:", error);
            setError("Invalid email or password."); // Set error message

            if (error?.response?.data?.message === "user was expired") {
                navigate("/user/expired");
            }
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
            <div
                className="bg-white p-4 rounded shadow"
                style={{
                    width: "400px",
                }}
            >
                <form
                    onSubmit={handleSubmit}
                    className="d-flex flex-column align-items-center"
                >
                    <h1 className="mb-4">Login</h1>
                    <div className="mb-4">
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
                            style={{ width: "350px" }}
                        />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="password" className="form-label">
                            Password:
                        </label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="form-control"
                            style={{ width: "350px" }}
                        />
                    </div>
                    {error && (
                        <div className="alert alert-danger" role="alert">
                            {error}
                        </div>
                    )}
                    <button type="submit" className="btn btn-primary w-100">
                        Login
                    </button>
                    <p className="mt-3">
                        Don't have an account?{" "}
                        <Link to="/register">Sign up</Link>
                    </p>
                    <p className="mt-2 mb-4">
                        <a href="/user/forgot-password">
                            Forgot your password?
                        </a>
                    </p>
                </form>
            </div>
        </div>
    );
}
