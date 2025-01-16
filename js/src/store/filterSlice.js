import { createSlice } from '@reduxjs/toolkit'

const initialState = {
    isShowingHeadshot: localStorage.getItem("showHeadshot") === "true",
    filters: {
        player: [],
        stat_type: [],
        team: [],
    },
    csgoFilters: {
        hidePP: false,
        hideUD: false,
        hideSP: false
    },
}

const filtersSlice = createSlice({
    name: 'filters',
    initialState,
    reducers: {
        setShowHeadshot: (state, action) => {
            state.isShowingHeadshot = action.payload
            localStorage.setItem("showHeadshot", action.payload);
        },
        setFilters: (state, action) => {
            const { filterSet } = action.payload;
            state.filters = filterSet;
        },
        setCsgoFilters: (state, action) => {
            state.csgoFilters = {
                ...state.csgoFilters,
                ...action.payload
            };
        }
    },
})

export const { setShowHeadshot, setFilters, setCsgoFilters } = filtersSlice.actions

export const ShowHeadshotOrNot = (state) => state.filters.isShowingHeadshot
export const GetCsGoFilters = (state) => state.filters.csgoFilters

export default filtersSlice.reducer
