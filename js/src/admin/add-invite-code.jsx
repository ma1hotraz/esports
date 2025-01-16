import React, { useState } from "react";
import { Form, Button, Alert } from "react-bootstrap";
import axios from "axios";
import { BASE_URL } from "../constants";
import { DatePicker } from "rsuite";
import { generateRandomString } from "../utils";
import TimeDurationSelector from "../components/time-duration";
import { useNavigate } from "react-router-dom";

const AddInviteCode = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        invite_code: "",
        expiration_time: new Date(),
        is_infinite: false,
    });

    // const [years, setYears] = useState(0);
    // const [months, setMonths] = useState(0);
    // const [days, setDays] = useState(0);
    // const [hours, setHours] = useState(0);
    // const [minutes, setMinutes] = useState(0);

    // const handleYearsChange = (e) => {
    //     setYears(parseInt(e.target.value));
    // };

    // const handleMonthsChange = (e) => {
    //     setMonths(parseInt(e.target.value));
    // };

    // const handleDaysChange = (e) => {
    //     setDays(parseInt(e.target.value));
    // };

    // const handleHoursChange = (e) => {
    //     setHours(parseInt(e.target.value));
    // };

    // const handleMinutesChange = (e) => {
    //     setMinutes(parseInt(e.target.value));
    // };
    const [error, setError] = useState("");

    const handleChange = (e) => {
        const { name, value, type, checked } = e.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]: type === "checkbox" ? checked : value,
        }));
    };
    const generateCode = () => {
        setFormData((prevState) => ({
            ...prevState,
            invite_code: generateRandomString(15),
        }));
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (error !== "") {
            return;
        }
        axios
            .post("/api/createInviteCode", {
                ...formData,
            })
            .then(() => {
                navigate("/admin/invite-code");
            })
            .catch((error) => {
                // Set error state
                setError(error.response.data.error);
                console.error("Error creating invite code:", error);
            });
    };

    return (
        <>
            {error && <Alert variant="danger">{error}</Alert>}
            <Form onSubmit={handleSubmit}>
                <Form.Group className="mb-3" controlId="invite_code">
                    <Form.Label>Invite Code</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter the code"
                        name="invite_code"
                        value={formData.invite_code}
                        onChange={handleChange}
                        required
                    />
                    <Button
                        onClick={generateCode}
                        variant="secondary"
                        type="button"
                    >
                        Generate
                    </Button>
                </Form.Group>
                <Form.Group className="mb-3" controlId="invite_code">
                    <Form.Label>
                        Expiration Time after enter invite code?
                    </Form.Label>
                    <Form.Check
                        type="checkbox"
                        label="Infinite access?"
                        name="is_infinite"
                        checked={formData.is_infinite}
                        onChange={handleChange}
                    />
                    {!formData.is_infinite ? (
                        <>
                            {/* <TimeDurationSelector
                                years={years}
                                days={days}
                                months={months}
                                hours={hours}
                                minutes={minutes}
                                handleDaysChange={handleDaysChange}
                                handleYearsChange={handleYearsChange}
                                handleMonthsChange={handleMonthsChange}
                                handleHoursChange={handleHoursChange}
                                handleMinutesChange={handleMinutesChange}
                            /> */}
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
        </>
    );
};

export default AddInviteCode;
