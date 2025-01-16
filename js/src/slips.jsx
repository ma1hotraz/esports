import React from "react";
import { CustomScroll } from "react-custom-scroll";
import { useSelector } from "react-redux";
import { Button, Panel, Table } from "rsuite";
import { CsgoDiffCell, DiffCell, TimeCell, ValCell } from "./cell";
import { EmptyDataPage } from "./lines";
import { selectSlipsOfEsport } from "./store/pairsSlice";
import { PrizePickBetCheckedCell, TeamUsedCell, UsedCell, UsedRow } from "./used";
import { getProjectionUrl } from "./utils";
import { CopyLinkButton } from "./header";
import { GetCsGoFilters } from "./store/filterSlice";

function SlipsTable({ data, header }) {
    const isLoading = useSelector(state => state.pairs.isLoading);
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
            {/* <Table.Column width={90}>
                <Table.HeaderCell>% Difference</Table.HeaderCell>
                <Table.Cell dataKey="percent_difference" />
            </Table.Column> */}
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
            <Table.Column width={100}>
                <Table.HeaderCell>PrizePicks Bet</Table.HeaderCell>
                <PrizePickBetCheckedCell />
            </Table.Column>
        </Table>
    );
}

function CSGOSlipsTable({ data, header }) {
    const isLoading = useSelector(state => state.pairs.isLoading);
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
            <Table.Column width={100} sortable>
                <Table.HeaderCell>Time on Game</Table.HeaderCell>
                <TimeCell dataKey="timestamp" />
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


function Slips({ league }) {

    const data = useSelector(state => selectSlipsOfEsport(state, league));

    const handeGenClick = (slip) => {
        const url = getProjectionUrl(slip.flat().map((p) => p.projection_string));
        open(url, "_blank");

    }

    return <>{data && data.length > 0 ? (
        <CustomScroll>
            <div className="custom-scroll-container">
                {data.map((slip, i) => (
                    <Panel
                        header={<>
                            <div>
                                <strong style={{ marginRight: "1rem" }}>{`Slip #${i + 1}`}</strong>
                            </div>
                            <div style={{ marginTop: "0.5rem" }}>
                                <Button style={{ marginRight: "0.5rem" }} size="sm" appearance="primary"
                                    color="blue" onClick={() => handeGenClick(slip)}>Bet on PrizePicks</Button>
                                <CopyLinkButton ppCheckedList={slip.flat().map((p) => p.projection_string)} />
                            </div>
                        </>}
                        key={`panel_${i}`}
                        className="slips-panel"
                    >
                        {slip?.map((pair, j) => (league === "csgo" ? <CSGOSlipsTable
                            data={pair}
                            header={j === 0}
                            key={`table_${i}_${j}`}
                        /> :
                            <SlipsTable
                                data={pair}
                                header={j === 0}
                                key={`table_${i}_${j}`}
                            />
                        ))}
                    </Panel>
                ))}
            </div>
        </CustomScroll>
    ) : <EmptyDataPage />}</>;
}

export function LOLSlips() {
    return <Slips league="lol" />;
}


export function CODSlips() {
    return <Slips league="cod" />;
}

export function CSGOSlips() {
    return <Slips league="csgo" />;
}

export function VALSlips() {
    return <Slips league="val" />;
}

export function HALOSlips() {
    return <Slips league="halo" />;
}

