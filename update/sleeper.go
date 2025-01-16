package update

import (
	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/notifier"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/iloginow/esportsdifference/sleeper"
	"github.com/sirupsen/logrus"
)

func getNewSleeperData() []dto.SleeperRelevantData {
	rawData, err := sleeper.GetData()
	if err != nil {
		logrus.Error(err)
	}
	return rawData.Filter()
}

func getPrevSleeperData() []dto.SleeperRelevantData {
	data, err := repo.GetSleeperRelevant()
	if err != nil {
		logrus.Error(err)
	}
	return data
}

func countNewSleeperRecords(
	newRecords []dto.SleeperRelevantData,
	oldRecords []dto.SleeperRelevantData,
) notifier.NewLinesCount {
	nl := notifier.NewLinesCount{}
	if len(oldRecords) < 1 {
		return nl
	}
	for _, r := range newRecords {
		if !findRelevantSleeperRecord(r, oldRecords) {
			if r.Sport == esport.COD {
				nl.COD++
			}
			if r.Sport == esport.CSGO {
				nl.CSGO++
			}
			if r.Sport == esport.VAL {
				nl.VAL++
			}
			if r.Sport == esport.HALO {
				nl.HALO++
			}
		}
	}
	return nl
}

func getNewSleeperRecords(
	newRecords []dto.SleeperRelevantData,
	oldRecords []dto.SleeperRelevantData,
) []dto.SleeperRelevantData {
	nl := []dto.SleeperRelevantData{}
	if len(oldRecords) < 1 {
		return nl
	}
	for _, r := range newRecords {
		if !findRelevantSleeperRecord(r, oldRecords) {
			nl = append(nl, r)
		}
	}
	return nl
}

func findRelevantSleeperRecord(record dto.SleeperRelevantData, recordList []dto.SleeperRelevantData) bool {
	for _, r := range recordList {
		if r.Player == record.Player && r.Sport == record.Sport && r.StatType == record.StatType && r.Time.Equal(record.Time) {
			return true
		}
	}
	return false
}
