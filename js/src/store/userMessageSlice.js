import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "axios";

export const fetchMyMessages = createAsyncThunk(
    "usersMessage/fetchMyMessages",
    async () => {
        const userId = localStorage.getItem("userId")
        if (!userId) return null;
        const response = await axios.get(`/api/my-notification/${userId}`);
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

const messageSlice = createSlice({
    name: "messages",
    initialState: {
        data: [],
        loading: false,
        error: null,
    },
    reducers: {
        clearMessages: (state) => {
            state.data = [];
        }
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchMyMessages.fulfilled, (state, action) => {
                state.loading = false;
                state.data = action.payload;
            })
            .addCase(fetchMyMessages.pending, (state, action) => {
                state.loading = false;
            })
            .addCase(fetchMyMessages.rejected, (state, action) => {
                state.loading = false;
                state.error = "error";
            })
            .addCase(dismissAllMyMessages.fulfilled, (state, action) => {
                state.loading = false;
                state.data = [];
            })
            .addCase(dismissAllMyMessages.pending, (state, action) => {
                state.loading = true;
            })
            .addCase(dismissAllMyMessages.rejected, (state, action) => {
                state.loading = false;
                state.error = "error";
            })
            ;
    },
});

export const selectMessages = (state) => state.messages.data;

export const { clearMessages } = messageSlice.actions;

export default messageSlice.reducer;
