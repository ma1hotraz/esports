import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "axios";

export const deleteAdminMessage = createAsyncThunk(
  "adminMessage/deleteAdminMessage",
  async (notification_id) => {
    await axios.delete(`/api/notification/${notification_id}`);
    return notification_id;
  }
)

export const fetchAdminMessages = createAsyncThunk(
  "adminMessage/fetchAdminMessages",
  async () => {
    const response = await axios.get(`/api/notification`);
    return response.data.notifications;
  }
)

export const dismissAllMyMessages = createAsyncThunk(
  "usersMessage/dismissAllMyMessages",
  async (arg, { getState }) => {
    const state = getState();
    const userId = localStorage.getItem("userId")
    if (!userId) return null;

    const response = await Promise.all(state.messages.data.map((item) => item.notification_id).map(async (notif_id) => {
      await axios.post(`/api/notification/dismiss`, {
        "user_id": +userId,
        "notification_id": notif_id
      });
    }))

    return response.data;
  }
)

const adminMessageSlice = createSlice({
  name: "adminMessages",
  initialState: {
    data: [],
    loading: false,
    error: null,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchAdminMessages.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
      })
      .addCase(fetchAdminMessages.pending, (state, action) => {
        state.loading = false;
      })
      .addCase(fetchAdminMessages.rejected, (state, action) => {
        state.loading = false;
        state.error = "error";
      })
      .addCase(deleteAdminMessage.fulfilled, (state, action) => {
        state.data = state.data.filter((item) => item.notification_id !== action.payload);
        state.loading = false;
      })
      .addCase(deleteAdminMessage.rejected, (state, action) => {
        state.error = "error";
      })
      .addCase(deleteAdminMessage.pending, (state, action) => {
        state.loading = true;
      })
  }
}
);

export const selectAdminMessages = (state) => state.adminMessages.data;
export const selectAdminMessagesLoading = (state) => state.adminMessages.loading;
export default adminMessageSlice.reducer;
