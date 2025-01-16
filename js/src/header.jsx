import axios from "axios";
import { useDispatch, useSelector } from "react-redux";
import { useLocation } from "react-router-dom";
import { Button, ButtonGroup, Header } from "rsuite";
import { BASE_URL } from "./constants";
import { clearProjections, selectProjections } from "./store/prizepickProjectionSlice";
import { removeAllUsed } from "./store/usedSlice";
import { fetchMyInfo } from "./store/userSlice";
import Subnav from "./subnav";
import { getProjectionUrl } from "./utils";
import { fetchLines } from "./store/lineSlice";
function getTitle(location) {
    const parts = (location || "").split("/");
    return parts.length < 2 ? "" : parts[1].toUpperCase();
}
import React, { useState } from 'react';

export const CopyLinkButton = ({ ppCheckedList, showNumber }) => {
    const [isCopied, setIsCopied] = useState(false);

    const copyToClipboard = async () => {
        const link = getProjectionUrl(ppCheckedList);
        try {
            await navigator.clipboard.writeText(link);
            setIsCopied(true);
            setTimeout(() => setIsCopied(false), 1500); // Reset after 2 seconds
        } catch (err) {
            console.error('Failed to copy text: ', err);
        }
    };
    return (
        <Button size="sm" appearance="primary" onClick={copyToClipboard}>
            {isCopied ? 'Link Copied!' : `Copy Link`}
        </Button>
    );
};

export const handleLogout = async () => {
    try {
        const response = await axios.delete(BASE_URL + "/api/logout");
        if (response.status === 200) {
            localStorage.removeItem("isLoggedIn");
            localStorage.clear();
            window.location.href = window.location.origin + "/login";
        }
    } catch (error) {
        console.error("failed:", error);
    }
};
export default function MainHeader({ pairs, slips, usedList, userId, isLoading }) {

    const location = useLocation();
    const dispatch = useDispatch();

    const isAdminUrl = location.pathname.includes("admin");
    const ppCheckedList = useSelector(state => selectProjections(state));
    async function handleClearUsed() {
        dispatch(removeAllUsed(userId));
    }
    const handleGenClick = () => {
        open(getProjectionUrl(ppCheckedList), "_blank");
        dispatch(clearProjections());
    }
    const handleGenClear = () => {
        dispatch(clearProjections());
    }
    return (
        <Header className="main-header">
            <h2 className="page-title">{getTitle(location.pathname)}</h2>
            <Subnav
                pathname={location.pathname}
                slips={slips}
                pairs={pairs}
            />

            <div className="header-controls">
                {!isAdminUrl && ppCheckedList.length !== 0 ? (
                    <>
                        <Button size="sm" appearance="primary" onClick={handleGenClick}>
                            ({ppCheckedList.length}) Bet on PrizePicks
                        </Button>
                        <CopyLinkButton ppCheckedList={ppCheckedList} />
                        <Button size="sm" appearance="default" onClick={handleGenClear}>
                            Clear PrizePicks
                        </Button>
                    </>
                ) : (
                    null
                )}
                {!isAdminUrl && usedList.length !== 0 ? (
                    <Button size="sm" appearance="primary" onClick={handleClearUsed}>
                        Clear
                    </Button>
                ) : (
                    null
                )}
                {!isAdminUrl ? (
                    <Button
                        size="sm"
                        appearance="primary"
                        color="blue"
                        className="refresh-btn"
                        loading={isLoading}
                        onClick={() => {
                            dispatch(fetchMyInfo());
                            dispatch(fetchLines());
                        }}
                    >
                        Refresh
                    </Button>
                ) : null}
                <Button
                    size="sm"
                    appearance="primary"
                    color="blue"
                    className="refresh-btn"
                    onClick={() => handleLogout()}
                >
                    Logout
                </Button>
            </div>
        </Header>
    );
}


