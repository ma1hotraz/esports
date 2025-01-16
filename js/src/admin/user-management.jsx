import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchData } from "../store/userSlice";
import UserDataTable from "./UserDataTable";

export default function UserManage() {
    const dispatch = useDispatch();

    const { data, loading, error } = useSelector((state) => state.users);

    useEffect(() => {
        dispatch(fetchData());
    }, [dispatch]);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div>
            <UserDataTable data={data} />
        </div>
    );
}
