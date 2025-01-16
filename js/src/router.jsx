import React from "react";
import {
    Route,
    Navigate,
    createRoutesFromElements,
    createBrowserRouter,
    useRouteError,
} from "react-router-dom";
import Layout from "./layout";
import { COD, CSGO, LOL, VAL, DOTA, HALO } from "./lines";
import { CODSlips, CSGOSlips, HALOSlips, VALSlips, LOLSlips } from "./slips";
import { CODPairs, CSGOPairs, HALOPairs, VALPairs, LOLPairs } from "./pairs";
import LoginPage from "./login";
import SignUpPage from "./signup";
import UserManage from "./admin/user-management";
import InviteCode from "./admin/inviteCode";
import AddUser from "./admin/addUser";
import EditUser from "./admin/edit-user";

import AddInviteCode from "./admin/add-invite-code";
import ExpiredPage from "./user/expired";
import ForgotPasswordPage from "./forgot-password";
import ResetPasswordPage from "./reset-pw-page";
import AdminMessageForm from "./admin/admin-message-form";

const isAuthenticated = () => {
    const isLoggedIn = localStorage.getItem("isLoggedIn");
    return isLoggedIn === "true";
};

const isAdmin = () => {
    const isAdmin = localStorage.getItem("isAdmin");
    return isAdmin === "true";
};

function ErrorBoundary() {
    let error = useRouteError();
    console.error(error);
    // Uncaught ReferenceError: path is not defined
    return <div>Dang!</div>;
}

const routes = createRoutesFromElements(
    isAuthenticated() ? (
        <Route element={<Layout />} errorElement={<ErrorBoundary />}>
            <Route path="/" element={<Navigate to="/cod" replace />} />
            <Route path="login" element={<Navigate to="/cod" replace />} />
            <Route path="register" element={<SignUpPage />} />
            <Route path="cod" element={<COD />} />
            <Route path="cod/slips" element={<CODSlips />} />
            <Route path="cod/pairs" element={<CODPairs />} />
            <Route path="csgo" element={<CSGO />} />
            <Route path="csgo/slips" element={<CSGOSlips />} />
            <Route path="csgo/pairs" element={<CSGOPairs />} />
            <Route path="lol" element={<LOL />} />
            <Route path="lol/pairs" element={<LOLPairs />} />
            <Route path="lol/slips" element={<LOLSlips />} />
            <Route path="val" element={<VAL />} />
            <Route path="val/pairs" element={<VALPairs />} />
            <Route path="val/slips" element={<VALSlips />} />
            <Route path="dota" element={<DOTA />} />
            <Route path="halo" element={<HALO />} />
            <Route path="halo/pairs" element={<HALOPairs />} />
            <Route path="halo/slips" element={<HALOSlips />} />
            {isAdmin() ? (
                <>
                    <Route
                        exact
                        path="admin/user-manage"
                        element={<UserManage />}
                    />
                    <Route
                        exact
                        path="admin/invite-code"
                        element={<InviteCode />}
                    />
                    <Route
                        exact
                        path="admin/user-manage/add-user"
                        element={<AddUser />}
                    />
                    <Route
                        path="admin/user-manage/edit-user/:userId"
                        element={<EditUser />}
                    />
                    <Route
                        path="admin/invite-code/add-code"
                        element={<AddInviteCode />}
                    />
                    <Route
                        path="admin/message"
                        element={<AdminMessageForm />}
                    />
                </>
            ) : (
                <>
                    <Route
                        path="admin/*"
                        element={<Navigate to="/login" replace />}
                    />
                </>
            )}
        </Route>
    ) : (
        <>
            <Route exact path="/user/expired" element={<ExpiredPage />} />
            <Route
                exact
                path="/user/forgot-password"
                element={<ForgotPasswordPage />}
            />
            <Route
                exact
                path="/user/change-password"
                element={<ResetPasswordPage />}
            />
            <Route element={<Layout />}>
                <Route path="/" element={<Navigate to="login" replace />} />
                <Route path="login" element={<LoginPage />} />
                <Route path="register" element={<SignUpPage />} />
                <Route path="*" element={<Navigate to="login" replace />} />
            </Route>
        </>
    )
);

export default createBrowserRouter(routes);
