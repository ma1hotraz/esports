package update

import (
	"github.com/iloginow/esportsdifference/compare"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/sirupsen/logrus"
)

func getPrevCompareResult() compare.Result {
	result, err := repo.GetCompareResult()
	if err != nil {
		logrus.Error(err)
	}
	return result
}

func findNewCompareRecords(newResult, oldResult compare.Result) compare.Result {
	nl := compare.Result{}
	if oldResult.IsNotEmpty() {
		nl.COD = findNewCompareRecordsForSport(newResult.COD, oldResult.COD)
		nl.CSGO = findNewCompareRecordsForSport(newResult.CSGO, oldResult.CSGO)
		nl.VAL = findNewCompareRecordsForSport(newResult.VAL, oldResult.VAL)
		nl.HALO = findNewCompareRecordsForSport(newResult.HALO, oldResult.HALO)
		nl.LOL = findNewCompareRecordsForSport(newResult.LOL, oldResult.LOL)
		nl.DOTA = markNewCompareRecordsForSport(newResult.DOTA, oldResult.DOTA)
	}
	return nl
}

func markNewCompareRecords(newCompareResult, newRecords compare.Result) compare.Result {
	if newRecords.IsNotEmpty() {
		r := compare.Result{}
		r.COD = markNewCompareRecordsForSport(newCompareResult.COD, newRecords.COD)
		r.CSGO = markNewCompareRecordsForSport(newCompareResult.CSGO, newRecords.CSGO)
		r.VAL = markNewCompareRecordsForSport(newCompareResult.VAL, newRecords.VAL)
		r.HALO = markNewCompareRecordsForSport(newCompareResult.HALO, newRecords.HALO)
		r.LOL = markNewCompareRecordsForSport(newCompareResult.LOL, newRecords.LOL)
		r.DOTA = markNewCompareRecordsForSport(newCompareResult.DOTA, newRecords.DOTA)

		return r
	}
	return newCompareResult
}

func findNewCompareRecordsForSport(newRecords, oldRecords []compare.Record) []compare.Record {
	newCompareRecords := []compare.Record{}
	for _, newRecord := range newRecords {
		if !compare.IsRecordInList(newRecord, oldRecords) && newRecord.Difference > 0 && newRecord.Team != "" && newRecord.Opponent != "" {
			newCompareRecords = append(newCompareRecords, newRecord)
		}
	}
	return newCompareRecords
}

func markNewCompareRecordsForSport(newCompareResult, newFetchedRecords []compare.Record) []compare.Record {
	mr := []compare.Record{}
	for _, r := range newCompareResult {
		if compare.IsRecordInList(r, newFetchedRecords) && r.Difference > 0 {
			r.IsNew = true
		}
		mr = append(mr, r)
	}
	return mr
}
