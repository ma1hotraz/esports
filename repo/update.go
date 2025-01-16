package repo

import (
	"github.com/iloginow/esportsdifference/compare"
	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/prizepicks"
	"github.com/iloginow/esportsdifference/slips"
	"github.com/sirupsen/logrus"
)

type UpdateData struct {
	UnderdogData   []dto.UnderdogRelevantData
	SleeperData    []dto.SleeperRelevantData
	PrizepicksData []prizepicks.RelevantData
	CompareResult  compare.Result
	SlipsResult    slips.Result
	PairsResult    slips.PairsResult
	SleeperResult  compare.Result
}

func (u UpdateData) Save() {
	logrus.Infof("Update new data with: underdog %d, prize %d, compare_empty?: %t, slips_empty?: %t, pairs_empty?: %t", len(u.UnderdogData), len(u.PrizepicksData), u.CompareResult.IsNotEmpty(), u.SlipsResult.IsNotEmpty(), u.PairsResult.IsNotEmpty())

	SaveUnderdogfantazyRelevant(u.UnderdogData)

	SavePrizepicksRelevant(u.PrizepicksData)

	SaveSleeperRelevant(u.SleeperData)

	SaveCompareResult(u.CompareResult)

	SaveSlipsResult(u.SlipsResult)

	SavePairsResult(u.PairsResult)

	SaveSleeperResult(u.SleeperResult)
}
