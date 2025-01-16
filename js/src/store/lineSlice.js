
import { createAsyncThunk, createSelector, createSlice } from '@reduxjs/toolkit';
import axios from 'axios';

function isLineMatchedFilter(v, filterSet) {

    function isEmptyFilterObject(obj) {
        const { player, stat_type, team } = obj;
        return player.length == 0 && team.length == 0 && stat_type.length == 0;
    }


    if (!filterSet || isEmptyFilterObject(filterSet)) {
        return true;
    }

    const { player, team, stat_type } = filterSet;

    const teamOrComp =
        team.length == 0 || team.includes(v.team) || team.includes(v.component);
    const matches = [
        player.length == 0 || player.includes(v.player),
        teamOrComp,
        stat_type.length == 0 || stat_type.includes(v.stat_type),
    ];
    return !matches.includes(false);
}

function applyFilters(data, filterSet) {
    const { filters, isShowingHeadshot } = filterSet;
    if (!isShowingHeadshot) {
        data = data.filter((v) => !v.stat_type.includes("Headshot"));
    }
    return data.filter((v) => isLineMatchedFilter(v, filters));
}

const initialState = {
    cod: [],
    csgo: [],
    lol: [],
    val: [],
    dota: [],
    halo: [],
    loading: false,
};

export const fetchLines = createAsyncThunk("lines/fetchLines", async () => {
    const response = await axios.get("/compare");
    return response.data;
});

const lineSlice = createSlice({
    name: 'lines',
    initialState,
    reducers: {
        clearData: (state) => {
            state.cod = [];
            state.csgo = [];
            state.lol = [];
            state.val = [];
            state.dota = [];
            state.halo = [];
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchLines.fulfilled, (state, action) => {
                state.cod = action.payload.cod ?? [];
                state.csgo = action.payload.csgo ?? [];
                state.lol = action.payload.lol ?? [];
                state.val = action.payload.val ?? [];
                state.dota = action.payload.dota ?? [];
                state.halo = action.payload.halo ?? [];
                state.loading = false;
            })
            .addCase(fetchLines.pending, (state, action) => {
                state.loading = true;
            })
            .addCase(fetchLines.rejected, (state, action) => {
                state.loading = false;
                state.error = "error";
            })
    }
});

export const {
    clearData
} = lineSlice.actions;

export const selectAllLines = state => state.lines.cod.concat(state.lines.csgo, state.lines.lol, state.lines.val, state.lines.dota, state.lines.halo);

export const selectLinesOfEsport = createSelector(
    state => state.lines,
    (state) => state.filters,
    (state, esport) => esport,
    (lines, filters, esport) => {
        return applyFilters(lines[esport], filters);
    }
)

export const selectSlipsOfEsport = createSelector(
    state => state.lines,
    (state) => state.filters,
    (state, esport) => esport,
    (lines, filters, esport) => {
        return applyFilters(lines[esport], filters);
    },
)

export const selectFilteredLines = createSelector(
    state => state.lines,
    (lines) => {
        return lines;
    }
);

export default lineSlice.reducer;