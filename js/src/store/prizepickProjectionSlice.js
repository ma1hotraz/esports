import { createSelector, createSlice } from "@reduxjs/toolkit";
import { getProjectionUrl } from "../utils";
import { selectAllLines } from "./lineSlice";

/**
 * This slice is responsible for managing the projections that the user has selected
 * to view on the PrizePicks website.
 */
const projectionSlice = createSlice({
    name: "projections",
    // The initial state of the slice should be get from local storage
    initialState: {
        data: localStorage.getItem("projections") ? JSON.parse(localStorage.getItem("projections")) : [],
        loading: false,
        error: null,
    },
    reducers: {
        addProjection: (state, action) => {
            state.data.push(action.payload);
            localStorage.setItem("projections", JSON.stringify(state.data));
        },
        removeProjection: (state, action) => {
            state.data = state.data.filter((item) => item !== action.payload);
            localStorage.setItem("projections", JSON.stringify(state.data));
        },
        clearProjections: (state) => {
            state.data = [];
            localStorage.setItem("projections", JSON.stringify(state.data));
        },
    }
}
);

export const selectCheckedProjectionLines = (state) => state.projections.data;

export const selectProjections = createSelector(
    (state) => state.projections.data,
    (state) => selectAllLines(state),
    (projections, alllines) => {
        return projections.filter((p) => {
            return alllines.some((line) => line.projection_string === p)
        });
    }
);

export const selectProjectionsUrl = (state) => {
    return getProjectionUrl(state.projections.data);
}
export const { addProjection, removeProjection, clearProjections } = projectionSlice.actions;
export default projectionSlice.reducer;
