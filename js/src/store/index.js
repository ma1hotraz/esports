import { configureStore } from '@reduxjs/toolkit';
import userSlice from './userSlice';
import userMessageSlice from './userMessageSlice';
import adminMessageSlice from './adminMessageSlice';
import lineSlice from './lineSlice';
import usedSlice from './usedSlice';
import filterSlice from './filterSlice';
import pairsSlice from './pairsSlice';
import prizepickProjectionSlice from './prizepickProjectionSlice';

const store = configureStore({
    reducer: {
        users: userSlice,
        messages: userMessageSlice,
        adminMessages: adminMessageSlice,
        lines: lineSlice,
        used: usedSlice,
        pairs: pairsSlice,
        filters: filterSlice,
        projections: prizepickProjectionSlice,
    },
});

export default store;