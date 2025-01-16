import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

export const fetchAllUsed = createAsyncThunk(
    'used/fetchAllUsed',
    async () => {
        const userId = +localStorage.getItem("userId")
        const response = await axios.post(`/stats`, {
            user_id: userId,
        });
        return response.data.stats;
    }
);

export const addOneUsed = createAsyncThunk(
    'used/addOneUsed',
    async (payload) => {
        const response = await axios.post("/insert-used", {
            ...payload
        });

        return response.data.stat;
    }
);

export const removeOneUsedThunk = createAsyncThunk(
    'used/removeOneUsed',
    async ({ user_id, stat_id }) => {
        await axios.post("/remove-stat", { user_id, stat_id });
        return { user_id, stat_id };
    }
);

export const removeAllUsed = createAsyncThunk(
    'used/removeAllUsed',
    async (userId) => {
        const response = await axios.post(`/remove-stats`, {
            user_id: userId,
        });
        return response.data.stats; // Assuming this returns the updated list of stats
    }
);

const usedSlice = createSlice({
    name: 'used',
    initialState: {
        data: [],
        userData: {},
    },
    reducers: {
        setUserData: (state, action) => {
            state.userData = action.payload;
        },
        logOut: (state) => {
            state.userData = {
                isLoggedIn: false,
                isAdmin: false,
                email: null,
            }
        }
    },
    extraReducers: (builder) => {
        builder
            .addCase(addOneUsed.fulfilled, (state, action) => {
                state.data.push(action.payload);
            })
            .addCase(removeOneUsedThunk.fulfilled, (state, action) => {
                state.data = state.data.filter(
                    (item) => item.Id !== action.payload.stat_id
                );
            })
            .addCase(fetchAllUsed.pending, (state) => {
                state.loading = true;
            })
            .addCase(fetchAllUsed.fulfilled, (state, action) => {
                state.loading = false;
                state.data = action.payload;
            })
            .addCase(fetchAllUsed.rejected, (state) => {
                state.loading = false;
            })
            .addCase(removeAllUsed.fulfilled, (state) => {
                state.data = [];
            })
            ;

    },
});

export const { setUserData, logOut } = usedSlice.actions;
export const selectUsed = (state) => state.used.data;
export const selectUserData = (state) => state.used.userData;
export default usedSlice.reducer;
