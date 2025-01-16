import React, { useEffect, useState } from "react";
import { CustomScroll } from "react-custom-scroll";
import { useSelector } from "react-redux";
import { Button, Panel, Table } from "rsuite";
import { ButtonCell, CsgoDiffCell, DiffCell, TimeCell, ValCell } from "./cell";
import { EmptyDataPage } from "./lines";
import { selectPairsOfEsport } from "./store/pairsSlice";
import { PrizePickBetCheckedCell, TeamUsedCell, UsedCell, UsedRow } from "./used";
import { GetCsGoFilters } from "./store/filterSlice";

function PairsTable({ data, header, onRowClick }) {
    const isLoading = useSelector(state => state.pairs.loading);
    return (
        <Table
            data={data}
            loading={isLoading}
            autoHeight
            className={!header ? "slips-table-no-header" : ""}
            renderRow={(children, rowData) => (
                <UsedRow rowData={rowData}>{children}</UsedRow>
            )}
        >
            <Table.Column width={60}>
                <Table.HeaderCell>Used</Table.HeaderCell>
                <UsedCell />
            </Table.Column>
            <Table.Column width={100}>
                <Table.HeaderCell>Difference</Table.HeaderCell>
                <DiffCell dataKey="difference" />
            </Table.Column>
            <Table.Column width={90} sortable>
                <Table.HeaderCell>Exposure</Table.HeaderCell>
                <TeamUsedCell />
            </Table.Column>
            <Table.Column width={90}>
                <Table.HeaderCell>Sport</Table.HeaderCell>
                <Table.Cell dataKey="sport" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Player</Table.HeaderCell>
                <Table.Cell dataKey="player" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Team</Table.HeaderCell>
                <Table.Cell dataKey="team" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Opponent</Table.HeaderCell>
                <Table.Cell dataKey="opponent" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Stat Type</Table.HeaderCell>
                <Table.Cell dataKey="stat_type" />
            </Table.Column>
            <Table.Column width={110}>
                <Table.HeaderCell>Prize Picks</Table.HeaderCell>
                <ValCell dataKey="prize_picks" />
            </Table.Column>
            <Table.Column width={110}>
                <Table.HeaderCell>Underdog</Table.HeaderCell>
                <ValCell dataKey="underdog" />
            </Table.Column>
            <Table.Column width={130}>
                <Table.HeaderCell>Time on Game</Table.HeaderCell>
                <TimeCell dataKey="timestamp" />
            </Table.Column>
            <Table.Column width={90}>
                <Table.HeaderCell>Filter</Table.HeaderCell>
                <ButtonCell dataKey="team" onClick={onRowClick}></ButtonCell>
            </Table.Column>
            <Table.Column width={100}>
                <Table.HeaderCell>PrizePicks Bet</Table.HeaderCell>
                <PrizePickBetCheckedCell />
            </Table.Column>
        </Table>
    );
}

function CSGOPairsTable({ data, header, onRowClick }) {
    const isLoading = useSelector(state => state.pairs.loading);
    const { hidePP, hideUD, hideSP } = useSelector(GetCsGoFilters);

    return (
        <Table
            data={data}
            loading={isLoading}
            autoHeight
            className={!header ? "slips-table-no-header" : ""}
            renderRow={(children, rowData) => (
                <UsedRow rowData={rowData}>{children}</UsedRow>
            )}
        >
            <Table.Column width={60}>
                <Table.HeaderCell>Used</Table.HeaderCell>
                <UsedCell />
            </Table.Column>
            <Table.Column width={100} sortable>
                <Table.HeaderCell>Difference</Table.HeaderCell>
                <CsgoDiffCell hidePP={hidePP} hideSP={hideSP} hideUD={hideUD} />
            </Table.Column>
            <Table.Column width={90} sortable>
                <Table.HeaderCell>Exposure</Table.HeaderCell>
                <TeamUsedCell />
            </Table.Column>
            <Table.Column width={90}>
                <Table.HeaderCell>Sport</Table.HeaderCell>
                <Table.Cell dataKey="sport" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Player</Table.HeaderCell>
                <Table.Cell dataKey="player" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Team</Table.HeaderCell>
                <Table.Cell dataKey="team" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Opponent</Table.HeaderCell>
                <Table.Cell dataKey="opponent" />
            </Table.Column>
            <Table.Column flexGrow={1}>
                <Table.HeaderCell>Stat Type</Table.HeaderCell>
                <Table.Cell dataKey="stat_type" />
            </Table.Column>
            {
                hideUD ? null : <>
                    <Table.Column width={110} sortable>
                        <Table.HeaderCell>Underdog</Table.HeaderCell>
                        <ValCell dataKey="underdog" />
                    </Table.Column>
                </>
            }
            {
                hidePP ? null : <>
                    <Table.Column width={110} sortable>
                        <Table.HeaderCell>Prize Picks</Table.HeaderCell>
                        <ValCell dataKey="prize_picks" />
                    </Table.Column>
                </>
            }
            {
                hideSP ? null : <>
                    <Table.Column width={110} sortable>
                        <Table.HeaderCell>Sleeper</Table.HeaderCell>
                        <ValCell dataKey="sleeper" />
                    </Table.Column>
                    <Table.Column width={130} sortable>
                        <Table.HeaderCell>Sleeper Mutiplier</Table.HeaderCell>
                        <ValCell dataKey="sleeper_multiplier" />
                    </Table.Column>
                </>
            }
            <Table.Column width={130}>
                <Table.HeaderCell>Time on Game</Table.HeaderCell>
                <TimeCell dataKey="timestamp" />
            </Table.Column>
            <Table.Column width={90}>
                <Table.HeaderCell>Filter</Table.HeaderCell>
                <ButtonCell dataKey="team" onClick={onRowClick}></ButtonCell>
            </Table.Column>
            {
                hidePP ? null : <Table.Column width={100}>
                    <Table.HeaderCell>PrizePicks Bet</Table.HeaderCell>
                    <PrizePickBetCheckedCell />
                </Table.Column>
            }
        </Table>
    );
}

function Pairs({ league }) {
    const pairs = useSelector(state => selectPairsOfEsport(state, league))
    const isLoading = useSelector(state => state.pairs.loading);
    const [tableData, setTableData] = useState(pairs);

    useEffect(() => {
        if (pairs) {
            setTableData(() => pairs);
        }
    }, [pairs]);

    const onRowClick = (team) => {
        setTableData(() =>
            pairs?.filter((v) => v[0].team === team || v[1].team === team)
        );
    };

    return (
        <>
            {tableData && tableData.length > 0 ? (
                <CustomScroll>
                    <Button
                        appearance="primary"
                        color="blue"
                        className="refresh-btn"
                        loading={isLoading}
                        onClick={() => {
                            setTableData(() => pairs);
                        }}
                    >
                        All
                    </Button>
                    <div className="custom-scroll-container">
                        {tableData.map((pair, i) => (
                            <Panel
                                header={<>
                                    <strong style={{ marginRight: "1rem" }}>{`Pair #${i + 1}`}</strong>
                                </>}
                                key={`pair_panel_${i}`}
                                className="pairs-panel"
                            >
                                {league !== "csgo" ? <PairsTable
                                    data={pair}
                                    header={true}
                                    key={`table_${i}_${i}`}
                                    onRowClick={onRowClick}
                                /> : <CSGOPairsTable
                                    data={pair}
                                    header={true}
                                    key={`table_${i}_${i}`}
                                    onRowClick={onRowClick}
                                />
                                }
                            </Panel>
                        ))}
                    </div>
                </CustomScroll>
            ) : (
                <EmptyDataPage />
            )}
        </>
    );
}

export function CODPairs() {
    return <Pairs league="cod" />;
}

export function LOLPairs() {
    return <Pairs league="lol" />;
}

export function CSGOPairs() {
    return <Pairs league="csgo" />;
}

export function VALPairs() {
    return <Pairs league="val" />;
}

export function HALOPairs() {
    return <Pairs league="halo" />;
}
