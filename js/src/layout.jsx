import React, { useEffect } from "react";
import { Outlet, useLocation } from "react-router-dom";
import { Container, Content } from "rsuite";
import AutoLogout from "./components/AutoLogout";
import MainHeader from "./header";
import Navbar, { NavToggle } from "./navbar";
import NotificationBar from "./user/notification-bar";
import { useDispatch, useSelector } from "react-redux";
import { fetchLines } from "./store/lineSlice";
import { fetchAllUsed, selectUsed } from "./store/usedSlice";
import { selectUserId } from "./store/userSlice";
import { fetchpairs, selectPairsOfEsport1, selectSlipsOfEsport1 } from "./store/pairsSlice";

function getLeague(pathname) {
    const parts = (pathname || "").split("/");
    if (parts.length < 2) return "";
    return parts[1];
}

export default function Layout() {
    const dispatch = useDispatch();
    const location = useLocation();
    const league = getLeague(location.pathname);
    const { lol } = useSelector((state) => state.lines);
    const used = useSelector(selectUsed);
    const userId = useSelector(selectUserId);
    const usedList = useSelector(selectUsed);
    const isLoading = useSelector((state) => state.lines.loading);
    const pairs = useSelector(state => selectPairsOfEsport1(state, league))
    const slips = useSelector(state => selectSlipsOfEsport1(state, league))

    const [collapsed, setCollapsed] = React.useState(false);

    useEffect(() => {
        if (!lol || !lol.length) {
            dispatch(fetchLines())
            dispatch(fetchpairs())
        }
    }, []);
    useEffect(() => {
        if (!used || !used.length) {
            if (userId) {
                dispatch(fetchAllUsed(userId))
            }
        }
    }, [userId]);

    const loginOrSignUp =
        location.pathname === "/login" || location.pathname === "/register";

    return (
        <Container>
            {!loginOrSignUp && !collapsed && <Navbar />}
            {!loginOrSignUp && <NavToggle expand={collapsed} onChange={() => setCollapsed(collapsed => !collapsed)} />}
            {!loginOrSignUp && (
                <Container>
                    <MainHeader isLoading={isLoading} pairs={pairs}
                        slips={slips} usedList={usedList} userId={userId} />
                    <Content className="main-content">
                        <Outlet />
                        <NotificationBar />
                    </Content>
                </Container>
            )}
            {loginOrSignUp && (
                <Content>
                    <Outlet />
                </Content>
            )}
            <AutoLogout />
        </Container>
    );
}
