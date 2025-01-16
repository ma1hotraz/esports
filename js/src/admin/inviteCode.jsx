import React from "react";
import { useState } from "react";
import { useEffect } from "react";
import axios from "axios";
import InviteDataTable from "./inviteDataTable";
import { BASE_URL } from "../constants";

export default function InviteCode() {
  const [userData, setUserData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(BASE_URL + "/api/getInviteCode");
        setUserData(response.data);
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    fetchData();
  }, []);
  return (
    <div>
      <InviteDataTable data={userData} />
    </div>
  );
}
