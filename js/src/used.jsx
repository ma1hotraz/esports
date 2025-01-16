import React from "react";
import { Checkbox, Table } from 'rsuite';
import { isTeamUsed, matchUsed } from './store/utils';
import ErrorOutlineIcon from '@mui/icons-material/ErrorOutline';
import Badge from '@mui/material/Badge';
import { styled } from '@mui/material/styles';
import { useDispatch, useSelector } from "react-redux";
import { addProjection, removeProjection, selectProjections } from "./store/prizepickProjectionSlice";
import { addOneUsed, removeOneUsedThunk, selectUsed } from "./store/usedSlice";
import { selectUserId } from "./store/userSlice";

export function UsedCell({ rowData, ...props }) {
  const usedList = useSelector(selectUsed);
  const dispatch = useDispatch();
  const isUsed = usedList.some((item) => matchUsed(rowData, item));
  const userId = useSelector(selectUserId);

  const addUsed = async () => {
    const payload = {
      user_id: +userId,
      player: rowData.player,
      stat_type: rowData.stat_type,
      timestamp: rowData.timestamp,
      team: rowData.team,
      opponent: rowData.opponent,
      sport: rowData.sport
    };

    dispatch(addOneUsed(payload));
  };

  const removeUsed = async () => {
    const thisUsed = isUsed ? usedList.filter((item) => matchUsed(rowData, item))[0] : null
    if (thisUsed) {
      dispatch(removeOneUsedThunk({ user_id: +userId, stat_id: thisUsed?.Id }))
    }
  }
  return (
    <Table.Cell {...{ className: 'table-checkbox-cell', ...props }}>
      <Checkbox checked={isUsed} onChange={isUsed ? removeUsed : addUsed} />
    </Table.Cell>
  );
};



export function PrizePickBetCheckedCell({ rowData, ...props }) {
  const checkedList = useSelector(state => selectProjections(state));
  const dispatch = useDispatch();
  const isUsed = checkedList.some((item) => item === rowData["projection_string"]);
  const removeUsed = () => {
    if (isUsed) {
      dispatch(removeProjection(rowData["projection_string"]));
    }
  }
  const addUsed = () => {
    if (!isUsed) {
      dispatch(addProjection(rowData["projection_string"]));
    }
  }
  return (
    <Table.Cell {...{ className: 'table-checkbox-cell', ...props }}>
      <Checkbox checked={isUsed} onChange={isUsed ? removeUsed : addUsed} />
    </Table.Cell>
  );
};


const UsedRowBadge = styled(Badge)(() => ({
  '& .MuiBadge-badge': {
    backgroundColor: "#c5c6c7",
  },
}));

export function TeamUsedCell({ rowData, ...props }) {
  const usedList = useSelector(selectUsed);
  const isUsed = rowData ? usedList.some((item) => matchUsed(rowData, item)) : false;

  const listUsedTeam = usedList.filter(used =>
    new Date(used.timestamp) > new Date() && used.team === rowData.team && used.component === rowData.component
  ).filter((item) => isTeamUsed(rowData, item))

  const CustomBadge = isUsed ? UsedRowBadge : Badge;

  return (
    <Table.Cell {...{ className: 'table-checkbox-cell', ...props }}>
      {listUsedTeam.length > 0 ? <CustomBadge badgeContent={listUsedTeam.length} color="primary">
        <ErrorOutlineIcon />
      </CustomBadge> : null}
    </Table.Cell>
  );
};


export function UsedRow({ rowData, children }) {
  const usedList = useSelector(selectUsed);
  const isUsed = rowData ? usedList.some((item) => matchUsed(rowData, item)) : false;

  const classNames = [];

  if (isUsed && rowData && !rowData.is_new) {
    classNames.push('used-row');
  }
  if (rowData && rowData.is_new) {
    classNames.push('new-row');
  }
  return (
    <div className={classNames.join(' ')}>{children}</div>
  );
}
