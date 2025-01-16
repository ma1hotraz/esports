import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { BASE_URL } from "../constants";

// Define an async thunk action
export const fetchData = createAsyncThunk("users/fetchData", async () => {
  const response = await axios.get(`${BASE_URL}/api/allUser`);
  return response.data;
});
export const fetchUserById = createAsyncThunk(
  "users/fetchUserById",
  async (userId) => {
    const response = await axios.get(`/api/allUser/${userId}`);
    return response.data;
  }
);

export const fetchMyInfo = createAsyncThunk(
  "users/fetchMyInfo",
  async () => {
    const userId = localStorage.getItem("userId")
    if (!userId) return null;
    const response = await axios.get(`/api/allUser/${userId}`);
    return response.data;
  }
);

const userSlice = createSlice({
  name: "users",
  initialState: {
    data: [],
    user: null,
    myInfo: null,
    messages: [],
    loading: false,
    error: null,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchData.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchData.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
      })
      .addCase(fetchData.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message;
      })
      .addCase(fetchUserById.pending, (state) => {
        state.loading = true;
        state.user = null;
      })
      .addCase(fetchUserById.fulfilled, (state, action) => {
        state.loading = false;
        state.user = action.payload;
      })
      .addCase(fetchUserById.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message;
        state.user = null;
      }).addCase(fetchMyInfo.pending, (state) => {
        state.loading = true;
        state.myInfo = null;
      })
      .addCase(fetchMyInfo.fulfilled, (state, action) => {
        state.loading = false;
        // console.log("set new user info");
        state.myInfo = action.payload;
      })
      .addCase(fetchMyInfo.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message;
        state.myInfo = null;
      })
      ;
  },
});


export const selectData = (state) => state.users.data;

export const selectIsAdmin = (state) => state.users.myInfo?.is_admin == true;
export const selectUserId = (state) => state.users.myInfo?.id;
export default userSlice.reducer;
