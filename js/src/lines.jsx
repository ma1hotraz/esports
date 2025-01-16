import React from "react";
import { useSelector } from 'react-redux';
import { selectLinesOfEsport } from "./store/lineSlice";
import Table, { CsgoDataTable } from "./table";

export function COD() {
  const cod = useSelector(state => selectLinesOfEsport(state, "cod"));
  return cod && <Table data={cod} />;
}

export function CSGO() {
  const csgo = useSelector(state => selectLinesOfEsport(state, "csgo"));

  return csgo && <CsgoDataTable data={csgo} />;
}

export function LOL() {

  const lol = useSelector(state => selectLinesOfEsport(state, "lol"));

  return lol && <Table data={lol} />;
}

export function VAL() {
  const val = useSelector(state => selectLinesOfEsport(state, "val"));

  return val && <Table data={val} />;
}

export function DOTA() {
  const dota = useSelector(state => selectLinesOfEsport(state, "dota"));
  return dota && <Table data={dota} />;
}

export function HALO() {
  const halo = useSelector(state => selectLinesOfEsport(state, "halo"));
  return halo && <Table data={halo} />;
}

export function EmptyDataPage() {
  return <Table data={[]} />;
}