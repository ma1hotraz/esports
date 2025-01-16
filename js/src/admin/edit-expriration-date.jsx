import React, { useState } from "react";

const EditExpirationPage = ({
    codeId,
    code,
    currentExpiration,
    onUpdateExpiration,
}) => {
    // Initialize state for form fields
    const [expiration, setExpiration] = useState({
        year: currentExpiration.getFullYear(),
        month: currentExpiration.getMonth() + 1, // Months are zero-indexed, so add 1
        day: currentExpiration.getDate(),
        hours: currentExpiration.getHours(),
        minutes: currentExpiration.getMinutes(),
    });

    // Handle input change for expiration fields
    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setExpiration({
            ...expiration,
            [name]: parseInt(value, 10), // Parse the input value as integer
        });
    };

    // Handle form submission
    const handleSubmit = (event) => {
        event.preventDefault();
        // Construct new Date by adding the entered values to current expiration
        const newExpirationDate = new Date(
            currentExpiration.getFullYear() + expiration.year,
            currentExpiration.getMonth() + expiration.month,
            currentExpiration.getDate() + expiration.day,
            currentExpiration.getHours() + expiration.hours,
            currentExpiration.getMinutes() + expiration.minutes
        );
        // Call the parent function to update expiration
        onUpdateExpiration(codeId, code, newExpirationDate);
    };

    return (
        <div>
            <h2>Edit Expiration Date</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Year:</label>
                    <input
                        type="number"
                        name="year"
                        value={expiration.year}
                        onChange={handleInputChange}
                    />
                </div>
                <div>
                    <label>Month:</label>
                    <input
                        type="number"
                        name="month"
                        value={expiration.month}
                        onChange={handleInputChange}
                    />
                </div>
                <div>
                    <label>Day:</label>
                    <input
                        type="number"
                        name="day"
                        value={expiration.day}
                        onChange={handleInputChange}
                    />
                </div>
                <div>
                    <label>Hours:</label>
                    <input
                        type="number"
                        name="hours"
                        value={expiration.hours}
                        onChange={handleInputChange}
                    />
                </div>
                <div>
                    <label>Minutes:</label>
                    <input
                        type="number"
                        name="minutes"
                        value={expiration.minutes}
                        onChange={handleInputChange}
                    />
                </div>
                <button type="submit">Update Expiration</button>
            </form>
        </div>
    );
};

export default EditExpirationPage;
