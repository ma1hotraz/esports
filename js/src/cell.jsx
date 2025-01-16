import React, { useEffect, useMemo } from "react";
import { Table, Button } from "rsuite";
import { formatTime } from "./store/utils";
import { useSelector } from "react-redux";
import { GetCsGoFilters } from "./store/filterSlice";
import { compareFields } from "./utils";

export function DiffCell({ rowData, dataKey, ...props }) {
    const color = (rowData.stat_type.includes("Combo")) ?
        "data-highlighted-orange combo-line" : "data-highlighted-orange"
    return (
        <Table.Cell {...{ className: "highlighted-cell", ...props }}>
            <span className={color}>{rowData[dataKey]}</span>
        </Table.Cell>
    );
}

export function calculateDifference(rowData, hidePP, hideUD, hideSP) {
    const prizePicks = hidePP ? 0 : rowData.prize_picks;
    const underdog = hideUD ? 0 : rowData.underdog;
    const sleeper = hideSP ? 0 : rowData.sleeper;

    // calculate difference between two non zero values to get the best value
    const prizePicksToUnderdog = prizePicks && underdog ? Math.abs(prizePicks - underdog) : 0;
    const prizePicksToSleeper = prizePicks && sleeper ? Math.abs(prizePicks - sleeper) : 0;
    const underdogToSleeper = underdog && sleeper ? Math.abs(underdog - sleeper) : 0;

    return Math.max(prizePicksToUnderdog, prizePicksToSleeper, underdogToSleeper);

}

export function calculateMininum(rowData, hidePP, hideUD, hideSP) {
    const fields = { PP: rowData["prize_picks"], UD: rowData.underdog, SP: rowData.sleeper, hidePP, hideUD, hideSP };
    return compareFields(fields);
}

export function CsgoDiffCell({ rowData, hidePP, hideUD, hideSP, ...props }) {
    const color = (rowData.stat_type.includes("Combo")) ?
        "data-highlighted-orange combo-line" : "data-highlighted-orange"
    const diff = calculateDifference(rowData, hidePP, hideUD, hideSP);
    return (
        <Table.Cell {...{ className: "highlighted-cell", ...props }}>
            <span className={color}>{diff}</span>
        </Table.Cell>
    );
}

export function ValCell({ rowData, dataKey, ...props }) {
    const { hidePP, hideUD, hideSP } = useSelector(GetCsGoFilters);
    const highlightRequirements = [
        dataKey === "prize_picks" && rowData[dataKey] < rowData.underdog && (rowData.sleeper ? rowData[dataKey] < rowData.sleeper : true),
        dataKey === "underdog" && rowData[dataKey] < rowData.prize_picks && (rowData.sleeper ? rowData[dataKey] < rowData.sleeper : true),
    ];
    const [className, setClassName] = React.useState("data-highlighted-none");
    useEffect(() => {
        let needGreen = false;

        if (rowData.sport === "CSGO" || rowData.sport === "csgo") {
            const result = calculateMininum(rowData, hidePP, hideUD, hideSP);
            needGreen = result[dataKey];
        } else {
            needGreen = highlightRequirements.some((req) => req)
        }
        const className = needGreen
            ? "data-highlighted-green"
            : "data-highlighted-none";
        setClassName(className);

    }, [rowData, hidePP, hideUD, hideSP]);

    if (dataKey === "sleeper_multiplier") {
        const text = rowData[dataKey] ? rowData["sleeper_over_under"].replace(" over", "x🔼").replace(" under", "x🔽").replace(";", " ") : "";
        return (
            <Table.Cell {...{ className: "highlighted-cell", ...props }}>
                <span style={{ justifyContent: "left" }} className={"data-highlighted-none"}>{text}</span>
            </Table.Cell>
        );
    }
    return (
        <Table.Cell {...{ className: "highlighted-cell", ...props }}>
            <span className={className}>{rowData[dataKey] ? rowData[dataKey] : ""}</span>
        </Table.Cell>
    );
}

export function TimeCell({ rowData, dataKey, ...props }) {
    const time = formatTime(rowData[dataKey]);
    return (
        <Table.Cell {...props}>
            <span>{time}</span>
        </Table.Cell>
    );
}

export function ButtonCell({ rowData, dataKey, onClick, ...props }) {
    return (
        <Table.Cell {...props}>
            <Button onClick={() => onClick(rowData[dataKey])}>🔍</Button>
        </Table.Cell>
    );
}
