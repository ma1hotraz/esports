import axios from "axios";
import React, { useState } from "react";
import { Alert, Button } from "react-bootstrap";
import { Trash } from "react-bootstrap-icons";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../constants";
import { fetchUserById } from "../store/userSlice";

const UserDataTable = ({ data }) => {
    const dispatch = useDispatch();

    const [error, setError] = useState(null);
    const navigate = useNavigate();
    const handleAddUser = () => {
        navigate("/admin/user-manage/add-user");
    };

    const handleEditUser = (id) => {
        dispatch(fetchUserById(id));
        navigate(`/admin/user-manage/edit-user/${id}`);
    };

    const handleDeleteUser = (id) => {
        axios
            .delete(BASE_URL + "/api/user/" + id)
            .then(() => {
                window.location.reload();
            })
            .catch((error) => {
                setError("Error deleting user. Please try again.");
                console.error("Error deleting user:", error);
            });
    };

    return (
        <div className="container mt-3">
            <div className="row align-items-center">
                <div className="col">
                    <h2>User Data</h2>
                </div>
                <div className="col-auto">
                    <Button variant="primary" onClick={handleAddUser}>
                        Add User
                    </Button>
                </div>
            </div>
            {error && <Alert variant="danger">{error}</Alert>}
            <table className="table">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Email</th>
                        <th>Invite Code</th>
                        <th>Expires at?</th>
                        <th>Admin</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {data?.map((user) => (
                        <tr key={user.id}>
                            <td>{user.id}</td>
                            <td>{user.name}</td>
                            <td>{user.email}</td>
                            <td>{user?.InviteCode?.invite_code}</td>
                            <td>
                                {user?.InviteCode?.expiration_time &&
                                !user?.InviteCode?.is_infinite
                                    ? new Date(
                                          user?.InviteCode?.expiration_time
                                      ).toLocaleString()
                                    : "infinite access"}
                            </td>
                            <td>{user?.is_admin ? "Yes" : "No"}</td>
                            <td>
                                <Button
                                    variant="danger"
                                    onClick={() => handleDeleteUser(user.id)}
                                >
                                    <Trash />
                                </Button>
                                <Button
                                    variant="info"
                                    onClick={() => handleEditUser(user.id)}
                                >
                                    🔗
                                </Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default UserDataTable;
