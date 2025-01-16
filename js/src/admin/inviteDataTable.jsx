import axios from "axios";
import React, { useState } from "react";
import { Alert, Button } from "react-bootstrap";
import { Trash } from "react-bootstrap-icons";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../constants";
const InviteDataTable = ({ data }) => {
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const handleAddCode = () => {
        navigate("/admin/invite-code/add-code");
    };

    const handleDeleteCode = (id) => {
        axios
            .delete(BASE_URL + "/api/invite-code/" + id)
            .then(() => {
                window.location.reload();
            })
            .catch((error) => {
                setError("Error deleting invite code. Please try again.");
                console.error("Error deleting invite code:", error);
            });
    };

    return (
        <div className="container mt-3">
            <div className="row align-items-center">
                <div className="col">
                    <h2>Invite Code</h2>
                </div>
                <div className="col-auto">
                    <button className="btn btn-primary" onClick={handleAddCode}>
                        Add Invite Code
                    </button>
                </div>
            </div>
            {error && <Alert variant="danger">{error}</Alert>}
            <table className="table">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Invite Code</th>
                        <th>Code Used?</th>
                        <th>Assigned to user</th>
                        <th>Expires at?</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((code) => (
                        <tr key={code.id}>
                            <td>{code.id}</td>
                            <td>{code?.invite_code}</td>
                            <td>{code?.used ? "Yes" : "No"}</td>
                            <td>
                                {code.users && code.users.length > 0
                                    ? code.users[0].email
                                    : ""}
                            </td>
                            <td>
                                {code?.is_infinite
                                    ? "infinite access	"
                                    : new Date(
                                          code?.expiration_time
                                      ).toLocaleString()}
                            </td>
                            <td>
                                {code.users && code.users.length > 0 ? null : (
                                    <Button
                                        too
                                        variant="danger"
                                        onClick={() =>
                                            handleDeleteCode(code.id)
                                        }
                                    >
                                        <Trash />
                                    </Button>
                                )}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default InviteDataTable;
