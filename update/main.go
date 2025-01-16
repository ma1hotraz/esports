package update

import (
	"github.com/iloginow/esportsdifference/compare"
	"github.com/iloginow/esportsdifference/notifier"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/iloginow/esportsdifference/slips"
)

func SyncNewDifferencesData() {
	newSleeperData := getNewSleeperData()
	newUnderdogData := getNewUnderdogData()
	newPrizepicksData := getNewPrizepicksData()

	newCompareResult := compare.CompareUnderdogToPrize(newUnderdogData, newPrizepicksData, newSleeperData)
	newSleeperCompareResult := compare.CompareUnderdogToSleeper(newUnderdogData, newSleeperData)

	pairsResult := slips.FindPairs(newCompareResult)
	slipsResult := slips.Find(newCompareResult)

	prevUnderdogData := getPrevUnderdogData()
	prevPrizepicksData := getPrevPrizepicksData()
	prevSleeperData := getPrevSleeperData()
	prevCompareResult := getPrevCompareResult()

	underdogDataResult := getNewUnderdogRecords(newUnderdogData, prevUnderdogData)
	prizepicksDataResult := getNewPrizepicksRecords(newPrizepicksData, prevPrizepicksData)
	sleeperDataResult := getNewSleeperRecords(newSleeperData, prevSleeperData)

	newCompareRecords := findNewCompareRecords(newCompareResult, prevCompareResult)

	compareDataResult := notifier.NewLinesData{
		COD:  newCompareRecords.COD,
		CSGO: newCompareRecords.CSGO,
		VAL:  newCompareRecords.VAL,
		// HALO: newCompareRecords.HALO,
	}

	finalCompareResult := markNewCompareRecords(newCompareResult, newCompareRecords)

	update := repo.UpdateData{
		UnderdogData:   newUnderdogData,
		PrizepicksData: newPrizepicksData,
		CompareResult:  finalCompareResult,
		SlipsResult:    slipsResult,
		PairsResult:    pairsResult,
		SleeperResult:  newSleeperCompareResult,
		SleeperData:    newSleeperData,
	}

	update.Save()
	notifier.AnnounceNewLines(underdogDataResult, prizepicksDataResult, compareDataResult, sleeperDataResult)
}

func CleanStaleInformed() {
	repo.CleanUpInformedData()
}
