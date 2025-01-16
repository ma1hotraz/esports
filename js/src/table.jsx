import React, { useState } from "react";
import { useSelector } from "react-redux";
import { Table } from 'rsuite';
import { calculateDifference, CsgoDiffCell, DiffCell, TimeCell, ValCell } from './cell';
import { PrizePickBetCheckedCell, TeamUsedCell, UsedCell, UsedRow } from './used';
import { GetCsGoFilters } from "./store/filterSlice";

const getData = (data, sortColumn, sortType) => {
  if (sortColumn && sortType) {
    return data.sort((a, b) => {
      const x = a[sortColumn];
      const y = b[sortColumn];
      if (typeof x === 'string' || typeof y === 'string') {
        if (sortType === 'asc') {
          return x.localeCompare(y);
        }
        return y.localeCompare(x);
      }
      if (sortType === 'asc') {
        return x - y;
      }
      return y - x;
    });
  }
  return data;
};

const getData1 = (data, sortColumn, sortType, hideUD, hidePP, hideSP) => {
  if (sortColumn && sortType) {
    return data.sort((a, b) => {
      let x = a[sortColumn];
      let y = b[sortColumn];
      if (sortColumn === 'difference') {
        x = calculateDifference(a, hidePP, hideUD, hideSP);
        y = calculateDifference(b, hidePP, hideUD, hideSP);
      }
      if (typeof x === 'string' || typeof y === 'string') {
        if (sortType === 'asc') {
          return x.localeCompare(y);
        }
        return y.localeCompare(x);
      }
      if (sortType === 'asc') {
        return x - y;
      }
      return y - x;
    });
  }
  return data;
};


export default function DataTable({ data }) {

  const isLoading = useSelector(state => state.lines.loading);
  const [sortColumn, setSortColumn] = useState('difference');
  const [sortType, setSortType] = useState('desc');
  const handleSortColumn = (col, type) => {
    setSortColumn(col);
    setSortType(type);
  };

  return (
    <div className="table-container">
      <Table
        data={getData(data, sortColumn, sortType)}
        sortColumn={sortColumn}
        sortType={sortType}
        onSortColumn={handleSortColumn}
        loading={isLoading}
        renderRow={(children, rowData) => <UsedRow rowData={rowData}>{children}</UsedRow>}
        fillHeight
      >
        <Table.Column width={60}>
          <Table.HeaderCell>Used</Table.HeaderCell>
          <UsedCell />
        </Table.Column>
        <Table.Column width={100} sortable>
          <Table.HeaderCell>Difference</Table.HeaderCell>
          <DiffCell dataKey="difference" />
        </Table.Column>
        <Table.Column width={90} sortable>
          <Table.HeaderCell>Exposure</Table.HeaderCell>
          <TeamUsedCell />
        </Table.Column>
        <Table.Column width={90} sortable>
          <Table.HeaderCell>Sport</Table.HeaderCell>
          <Table.Cell dataKey="sport" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
          <Table.HeaderCell>Player</Table.HeaderCell>
          <Table.Cell dataKey="player" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
          <Table.HeaderCell>Team</Table.HeaderCell>
          <Table.Cell dataKey="team" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
          <Table.HeaderCell>Opponent</Table.HeaderCell>
          <Table.Cell dataKey="opponent" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
          <Table.HeaderCell>Stat Type</Table.HeaderCell>
          <Table.Cell dataKey="stat_type" />
        </Table.Column>
        <Table.Column width={110} sortable>
          <Table.HeaderCell>Prize Picks</Table.HeaderCell>
          <ValCell dataKey="prize_picks" />
        </Table.Column>
        <Table.Column width={110} sortable>
          <Table.HeaderCell>Underdog</Table.HeaderCell>
          <ValCell dataKey="underdog" />
        </Table.Column>
        <Table.Column width={100} sortable>
          <Table.HeaderCell>Time on Game</Table.HeaderCell>
          <TimeCell dataKey="timestamp" />
        </Table.Column>
        <Table.Column width={100}>
          <Table.HeaderCell>PrizePicks Bet</Table.HeaderCell>
          <PrizePickBetCheckedCell />
        </Table.Column>
      </Table>
    </div >
  );
}


export function CsgoDataTable({ data }) {
  const isLoading = useSelector(state => state.lines.loading);
  const [sortColumn, setSortColumn] = useState('difference');
  const [sortType, setSortType] = useState('desc');
  const { hidePP, hideUD, hideSP } = useSelector(GetCsGoFilters);
  const handleSortColumn = (col, type) => {
    setSortColumn(col);
    setSortType(type);
  };

  return (
    <div className="table-container">
      <Table
        data={getData1(data, sortColumn, sortType, hideUD, hidePP, hideSP)}
        sortColumn={sortColumn}
        sortType={sortType}
        onSortColumn={handleSortColumn}
        loading={isLoading}
        renderRow={(children, rowData) => <UsedRow rowData={rowData}>{children}</UsedRow>}
        fillHeight
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
        {/* <Table.Column width={90}>
          <Table.HeaderCell>% Difference</Table.HeaderCell>
          <Table.Cell dataKey="percent_difference" />
        </Table.Column> */}
        <Table.Column width={90} sortable>
          <Table.HeaderCell>Sport</Table.HeaderCell>
          <Table.Cell dataKey="sport" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
          <Table.HeaderCell>Player</Table.HeaderCell>
          <Table.Cell dataKey="player" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
          <Table.HeaderCell>Team</Table.HeaderCell>
          <Table.Cell dataKey="team" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
          <Table.HeaderCell>Opponent</Table.HeaderCell>
          <Table.Cell dataKey="opponent" />
        </Table.Column>
        <Table.Column flexGrow={1} sortable>
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
    </div >
  );
}