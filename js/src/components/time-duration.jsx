import React, { useState } from "react";

const TimeDurationSelector = ({
    years,
    months,
    days,
    hours,
    minutes,
    handleYearsChange,
    handleMonthsChange,
    handleDaysChange,
    handleHoursChange,
    handleMinutesChange,
}) => {
    const getTotalMinutes = () => {
        return (
            years * 525600 + months * 43800 + days * 1440 + hours * 60 + minutes
        );
    };

    return (
        <div>
            <label>
                Years:
                <input
                    type="number"
                    value={years}
                    onChange={handleYearsChange}
                />
            </label>
            <br />
            <label>
                Months:
                <input
                    type="number"
                    value={months}
                    onChange={handleMonthsChange}
                />
            </label>
            <br />
            <label>
                Days:
                <input type="number" value={days} onChange={handleDaysChange} />
            </label>
            <br />
            <label>
                Hours:
                <input
                    type="number"
                    value={hours}
                    onChange={handleHoursChange}
                />
            </label>
            <br />
            <label>
                Minutes:
                <input
                    type="number"
                    value={minutes}
                    onChange={handleMinutesChange}
                />
            </label>
            <br />
            <p>Total Duration: {getTotalMinutes()} minutes</p>
        </div>
    );
};

export default TimeDurationSelector;
