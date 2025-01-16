import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Link, useLocation } from "react-router-dom";
import { Button, ButtonGroup, CheckPicker, Checkbox, CheckboxGroup } from "rsuite";
import { GetCsGoFilters, setCsgoFilters, setFilters, setShowHeadshot, ShowHeadshotOrNot } from "./store/filterSlice";

function show(pathname, pairsOrSlips) {
    const parts = (pathname || "").split("/");
    if (parts.length < 2) return false;
    if (pairsOrSlips?.length == 0) return false;
    return true;
}

function getLeague(pathname) {
    const parts = (pathname || "").split("/");
    if (parts.length < 2) return "";
    return parts[1];
}

function compare(a, b) {
    let nameA = a.toUpperCase();
    let nameB = b.toUpperCase();

    if (nameA < nameB) {
        return -1;
    }
    if (nameA > nameB) {
        return 1;
    }
    return 0;
}

export default function Subnav({ pathname, slips, pairs }) {
    const league = getLeague(pathname);
    const dispatch = useDispatch();
    const state = useSelector((state) => state.lines);
    const isShowHeadshot = useSelector(ShowHeadshotOrNot);
    const filterByFields = ["team", "player", "stat_type"];
    const toLabel = { team: "Team", player: "Player", stat_type: "Stat Type" };
    const fieldsMappingData = {};
    const stateLeague = isShowHeadshot ? state[league] : state[league]?.filter(v => !v.stat_type.includes("Headshot"));
    const { hidePP, hideUD, hideSP } = useSelector(GetCsGoFilters);
    filterByFields.forEach(
        (f) =>
        (fieldsMappingData[toLabel[f]] = [
            ...new Set(stateLeague?.map((v) => v[f])),
        ])
    );

    const filterCheckBoxesData = filterByFields
        .map((f) =>
            fieldsMappingData[toLabel[f]].map((item) => ({
                label: item,
                value: f + "_____" + item,
                role: toLabel[f],
            }))
        )
        .flat();

    return <>{state[league] && state[league].length > 0 ?
        <>
            <>
                {
                    league === "csgo" ?
                        <div style={{ marginRight: "1rem" }}>
                            <CheckboxGroup inline value={[`${hidePP ? "PP" : ""}`, `${hideUD ? "UD" : ""}`, `${hideSP ? "SP" : ""}`]}
                                onChange={value => {
                                    dispatch(setCsgoFilters({
                                        hidePP: value.includes("PP"),
                                        hideUD: value.includes("UD"),
                                        hideSP: value.includes("SP")
                                    }))

                                }} name="checkbox-group" >
                                <Checkbox value="PP">Hide PrizePicks</Checkbox>
                                <Checkbox value="UD">Hide Underdog</Checkbox>
                                <Checkbox value="SP">Hide Sleeper</Checkbox>
                            </CheckboxGroup>
                        </div> : null
                }
            </>
            <div
                style={{
                    display: "flex",
                    justifyContent: "space-between",
                    width: "45%",
                }}
            >

                <ButtonGroup size="sm">
                    {show(pathname, slips) || show(pathname, pairs) ? (
                        <Button
                            appearance={
                                basePathName(pathname) ? "primary" : "default"
                            }
                            className={basePathName(pathname) ? "subnav-btn" : ""}
                            color="blue"
                            as={Link}
                            to={league}
                        >
                            Lines
                        </Button>
                    ) : null}
                    {show(pathname, slips) ? (
                        <Button
                            appearance={
                                pathname.includes("slips") ? "primary" : "default"
                            }
                            className={
                                pathname.includes("slips") ? "subnav-btn" : ""
                            }
                            color="blue"
                            as={Link}
                            to={`${league}/slips`}
                        >
                            Slips
                        </Button>
                    ) : null}
                    {show(pathname, pairs) ? (
                        <Button
                            appearance={
                                pathname.includes("pairs") ? "primary" : "default"
                            }
                            className={
                                pathname.includes("pairs") ? "subnav-btn" : ""
                            }
                            color="blue"
                            as={Link}
                            to={`${league}/pairs`}
                        >
                            Pairs
                        </Button>
                    ) : null}
                </ButtonGroup>
                <div style={{ display: "flex", "flexDirection": "row", flexGrow: "1", marginLeft: "1.5rem" }}>
                    <FilterCheckBoxes
                        dispatch={dispatch}
                        filterCheckBoxesData={filterCheckBoxesData}
                    />
                    <div>
                        <Checkbox checked={isShowHeadshot} color="blue" onChange={(v, checked, e) => {
                            dispatch(setShowHeadshot(checked));
                        }}>
                            Headshots
                        </Checkbox>
                    </div>
                </div>
            </div>
        </> : null
    } </>;
}

function basePathName(pathname) {
    return !(pathname.includes("pairs") || pathname.includes("slips"));
}

function getFilterObject(filterSet) {
    const filterObj = {
        player: [],
        stat_type: [],
        team: [],
    };
    filterSet.map((v) =>
        filterObj[v.split("_____")[0]].push(v.split("_____")[1])
    );

    return filterObj;
}

function FilterCheckBoxes({ dispatch, filterCheckBoxesData }) {
    const location = useLocation();

    const [filterSet, setFilterSet] = useState([]);

    useEffect(() => {
        handleOnChange([]);
    }, [location]);

    const handleOnChange = (data) => {
        setFilterSet(() => data);
        dispatch(
            setFilters({
                filterSet: getFilterObject(data),
            })
        );
    };

    return (
        <div>
            <CheckPicker
                size="sm"
                data={filterCheckBoxesData}
                groupBy="role"
                sort={(isGroup) => {
                    if (isGroup) {
                        return (a, b) => {
                            return compare(a.groupTitle, b.groupTitle);
                        };
                    }

                    return (a, b) => {
                        return compare(a.value, b.value);
                    };
                }}
                style={{ width: 150 }}
                onChange={handleOnChange}
                value={filterSet}
            />
        </div>
    );
}
