import React from "react";
import {
    MdKeyboardArrowLeft,
    MdOutlineKeyboardArrowRight
} from 'react-icons/md';
import { useSelector } from "react-redux";
import { Link, useLocation } from "react-router-dom";
import { IconButton, Nav, Sidebar, Sidenav, Stack, Tag } from "rsuite";

function NavItemContent({ title, count }) {
    return (
        <div className="nav-item-content">
            <span>{title}</span>
            <Tag
                color={count > 0 ? "blue" : ""}
                className={count > 0 ? "nav-tag" : ""}
            >
                {count}
            </Tag>
        </div>
    );
}

export default function Navbar() {
    const { lol, cod, halo, csgo, val, dota } = useSelector(state => state.lines);
    const location = useLocation();

    const isAdminUrl = location.pathname.includes("admin");
    const isAdmin = localStorage.getItem("isAdmin") === "true";
    return (
        <Sidebar className="main-navbar" width={180}>
            <Sidenav.Header className="sidebar-header">
                <span>{isAdminUrl ? "ADMIN" : "ESPORTS"}</span>
            </Sidenav.Header>
            <Sidenav expanded appearance="subtle">
                <Sidenav.Body>
                    <Nav>
                        {isAdminUrl && (
                            <>
                                <Nav.Item
                                    as={Link}
                                    to="/cod"
                                    style={{ paddingLeft: 15 }}
                                >
                                    <NavItemContent title="Go to board" />
                                </Nav.Item>
                                <Nav.Item
                                    as={Link}
                                    to="admin/user-manage"
                                    active={location.pathname.includes(
                                        "user-manage"
                                    )}
                                    style={{ paddingLeft: 15 }}
                                >
                                    <NavItemContent title="User Manage" />
                                </Nav.Item>

                                <Nav.Item
                                    as={Link}
                                    to="admin/invite-code"
                                    active={location.pathname.includes(
                                        "invite-code"
                                    )}
                                    style={{ paddingLeft: 15 }}
                                >
                                    <NavItemContent title="Invite Code" />
                                </Nav.Item>

                                <Nav.Item
                                    as={Link}
                                    to="admin/message"
                                    active={location.pathname.includes(
                                        "admin/message"
                                    )}
                                    style={{ paddingLeft: 15 }}
                                >
                                    <NavItemContent title="Inform Users" />
                                </Nav.Item>
                            </>
                        )}
                        {!isAdminUrl && (
                            <>
                                {isAdmin ? (
                                    <Nav.Item as={Link} to="/admin/user-manage">
                                        <NavItemContent title="Admin" />
                                    </Nav.Item>
                                ) : null}
                                <Nav.Item
                                    as={Link}
                                    to="cod"
                                    active={location.pathname.includes("cod")}
                                >
                                    <NavItemContent
                                        title="COD"
                                        count={cod.length}
                                    />
                                </Nav.Item>
                                <Nav.Item
                                    as={Link}
                                    to="csgo"
                                    active={location.pathname.includes("csgo")}
                                >
                                    <NavItemContent
                                        title="CSGO"
                                        count={csgo.length}
                                    />
                                </Nav.Item>
                                <Nav.Item
                                    as={Link}
                                    to="lol"
                                    active={location.pathname.includes("lol")}
                                >
                                    <NavItemContent
                                        title="LOL"
                                        count={lol.length}
                                    />
                                </Nav.Item>
                                <Nav.Item
                                    as={Link}
                                    to="val"
                                    active={location.pathname.includes("val")}
                                >
                                    <NavItemContent
                                        title="VAL"
                                        count={val.length}
                                    />
                                </Nav.Item>
                                <Nav.Item
                                    as={Link}
                                    to="dota"
                                    active={location.pathname.includes("dota")}
                                >
                                    <NavItemContent
                                        title="DOTA"
                                        count={dota.length}
                                    />
                                </Nav.Item>
                                <Nav.Item
                                    as={Link}
                                    to="halo"
                                    active={location.pathname.includes("halo")}
                                >
                                    <NavItemContent
                                        title="HALO"
                                        count={halo.length}
                                    />
                                </Nav.Item>
                            </>
                        )}
                    </Nav>
                </Sidenav.Body>
            </Sidenav>
        </Sidebar>
    );
}

export function NavToggle({ expand, onChange }) {
    return (
        <Stack className="floating-toggle" justifyContent={expand ? 'flex-end' : 'center'}>
            <IconButton
                onClick={onChange}
                appearance="subtle"
                size="lg"
                icon={!expand ? <MdKeyboardArrowLeft /> : <MdOutlineKeyboardArrowRight />}
            />
        </Stack>
    );
};