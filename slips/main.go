package slips

import "github.com/iloginow/esportsdifference/compare"

type Result struct {
	COD  [][][]compare.Record `json:"cod"`
	CSGO [][][]compare.Record `json:"csgo"`
	VAL  [][][]compare.Record `json:"val"`
	HALO [][][]compare.Record `json:"halo"`
}

type PairsResult struct {
	COD  [][]compare.Record `json:"cod"`
	CSGO [][]compare.Record `json:"csgo"`
	VAL  [][]compare.Record `json:"val"`
	HALO [][]compare.Record `json:"halo"`
	LOL  [][]compare.Record `json:"lol"`
}

func Find(cr compare.Result) Result {
	result := Result{}
	result.COD = findSlipsForSport(cr.COD)
	result.CSGO = findSlipsForSport(cr.CSGO)
	result.VAL = findSlipsForSport(cr.VAL)
	result.HALO = findSlipsForSport(cr.HALO)

	return result
}

func FindPairs(cr compare.Result) PairsResult {
	result := PairsResult{}
	result.COD = findPairsForSport(cr.COD)
	result.CSGO = findPairsForSport(cr.CSGO)
	result.VAL = findPairsForSport(cr.VAL)
	result.HALO = findPairsForSport(cr.HALO)
	result.LOL = findPairsForSport(cr.LOL)
	return result
}

func (r Result) IsNotEmpty() bool {
	return len(r.COD) > 0 || len(r.CSGO) > 0 || len(r.VAL) > 0
}

func (r PairsResult) IsNotEmpty() bool {
	return len(r.COD) > 0 || len(r.CSGO) > 0 || len(r.VAL) > 0
}

func findSlipsForSport(cr []compare.Record) [][][]compare.Record {
	matches := NewMatches(cr)
	pairedElements := matches.Find()
	pairs := NewPairs(pairedElements)
	return pairs.Find()
}

func findPairsForSport(cr []compare.Record) [][]compare.Record {
	matches := NewMatches(cr)
	return matches.Find()
}

func ConvertPairsToSlips(cr [][]compare.Record) [][][]compare.Record {
	pairs := NewPairs(cr)
	return pairs.Find()
}
