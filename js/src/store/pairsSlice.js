
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
    if (!data) {
        return []
    }
    if (!filterSet) { 
        return data;
    }
    const { filters, isShowingHeadshot } = filterSet;
    if (!isShowingHeadshot) {
        data = data.filter((v) => v.flat().some(v1 => !v1.stat_type.includes("Headshot")));
    }
    return data.filter((v) => v.flat().every(v1 => isLineMatchedFilter(v1, filters)));
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

export const fetchpairs = createAsyncThunk("pairs/fetchpairs", async () => {
    const response = await axios.get("/pairs");
    return response.data;
});

const pairsSlice = createSlice({
    name: 'pairs',
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
            .addCase(fetchpairs.fulfilled, (state, action) => {
                state.cod = action.payload.cod ?? [];
                state.csgo = action.payload.csgo ?? [];
                state.lol = action.payload.lol ?? [];
                state.val = action.payload.val ?? [];
                state.dota = action.payload.dota ?? [];
                state.halo = action.payload.halo ?? [];
                state.loading = false;
            })
            .addCase(fetchpairs.pending, (state, action) => {
                state.loading = true;
            })
            .addCase(fetchpairs.rejected, (state, action) => {
                state.loading = false;
                state.error = "error";
            })
    }
});

export const {
    clearData
} = pairsSlice.actions;

export const selectPairsOfEsport = createSelector(
    state => state.pairs,
    (state) => state.filters,
    (state, esport) => esport,
    (pairs, filters, esport) => {
        if (!esport || !pairs[esport]) {
            return [];
        }
        const res = applyFilters(pairs[esport], filters);
        // console.log(res, pairs[esport], filters);
        return res;
    },
)

export const selectSlipsOfEsport = createSelector(
    state => state.pairs,
    (state) => state.filters,
    (state, esport) => esport,
    (pairs, filters, esport) => {
        if (!esport) {
            return [];
        }

        const filteredPairs = applyFilters(pairs[esport], filters);

        return generateSlips(filteredPairs);
    },
)

export const selectPairsOfEsport1 = createSelector(
    state => state.pairs,
    (state) => state.filters,
    (state, esport) => esport,
    (pairs, filters, esport) => {
        if (!esport || !pairs[esport]) {
            return [];
        }
        return applyFilters(pairs[esport], null);
    },
)

export const selectSlipsOfEsport1 = createSelector(
    state => state.pairs,
    (state) => state.filters,
    (state, esport) => esport,
    (pairs, filters, esport) => {
        if (!esport) {
            return [];
        }

        const filteredPairs = applyFilters(pairs[esport], null);

        return generateSlips(filteredPairs);
    },
)

function generateSlips(pairs) {
    // let firstFilter = pairs
    //     .filter((v) => {
    //         return v.flat().some((v1) => isLineMatchedFilter(v1, filterObj));
    //     })
    //     .flat()
    //     .filter(v => !v[0].stat_type.includes("Assists"));

    // if (firstFilter.length > 2) {
    //     firstFilter = firstFilter.filter((v) => {
    //         if (v.flat().some((v1) => isLineMatchedFilter(v1, filterObj))) {
    //             return true;
    //         }
    //     });
    // }
    if (!pairs || pairs.length === 0) {
        return [];
    }
    const sortedPairs = [...pairs]
    sortedPairs.sort((pair1, pair2) => {
        const diff1 = pair1[0].difference + pair1[1].difference;
        const diff2 = pair2[0].difference + pair2[1].difference;

        return diff2 - diff1;
    })
    return divideArrayV2(
        sortedPairs
    );
}


function divideArrayV2(array) {
    if (!array || array.length === 0) { return []; }
    let subArrays = [];

    for (let i = 0; i < array.length; i += 2) {
        if (i + 1 < array.length && !compare(array[i], array[i + 1])) {
            // Get two elements starting from the current index only if they are not the same
            let subArray = [array[i], array[i + 1]];
            subArrays.push(subArray);
        } else if (i + 1 < array.length) {
            // If they are the same, look for the next different element
            let j = i + 1;
            while (j < array.length && compare(array[i], array[i + 1])) {
                j++;
            }
            if (j < array.length) {
                let subArray = [array[i], array[j]];
                subArrays.push(subArray);
                i = j - 1; // Adjust the index to skip over the paired elements
            }
        } else {
            // If it's the last single element, just push it as a single element subarray
            subArrays.push([array[i]]);
        }
    }

    return subArrays.filter((s) => s.length === 2);
}

function compare(a, b) {
    const aTeams = [a[0].team, a[1].team, a[0].opponent, a[1].opponent].sort();
    const bTeams = [b[0].team, b[1].team, b[0].opponent, b[1].opponent].sort();

    for (let i = 0; i < aTeams.length; i++) {
        if (aTeams[i] !== bTeams[i]) {
            return false;
        }
    }

    return true;
}

export default pairsSlice.reducer;