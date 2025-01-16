package update

import (
	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/notifier"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/iloginow/esportsdifference/underdog"
	"github.com/sirupsen/logrus"
)

func getNewUnderdogData() []dto.UnderdogRelevantData {
	rawData, err := underdog.GetData()
	if err != nil {
		logrus.Error(err)
	}
	return rawData.Filter()
}

func getPrevUnderdogData() []dto.UnderdogRelevantData {
	data, err := repo.GetUnderdogfantazyRelevant()
	if err != nil {
		logrus.Error(err)
	}
	return data
}

func countNewUnderdogRecords(
	newRecords []dto.UnderdogRelevantData,
	oldRecords []dto.UnderdogRelevantData,
) notifier.NewLinesCount {
	nl := notifier.NewLinesCount{}
	if len(oldRecords) < 1 {
		return nl
	}
	for _, r := range newRecords {
		if !findRelevantUnderdogRecord(r, oldRecords) {
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

func getNewUnderdogRecords(
	newRecords []dto.UnderdogRelevantData,
	oldRecords []dto.UnderdogRelevantData,
) notifier.NewUnderdogData {
	nl := notifier.NewUnderdogData{}
	if len(oldRecords) < 1 {
		return nl
	}
	for _, r := range newRecords {
		// if not in old records - add to new lines
		if !findRelevantUnderdogRecord(r, oldRecords) {
			if r.Sport == esport.COD {
				nl.COD = append(nl.COD, r)
			}
			if r.Sport == esport.CSGO {
				nl.CSGO = append(nl.CSGO, r)
			}
			if r.Sport == esport.VAL {
				nl.VAL = append(nl.VAL, r)
			}
			if r.Sport == esport.HALO {
				nl.HALO = append(nl.HALO, r)
			}
			if r.Sport == esport.LOL {
				nl.LOL = append(nl.LOL, r)
			}
		}
	}
	return nl
}

func findRelevantUnderdogRecord(record dto.UnderdogRelevantData, recordList []dto.UnderdogRelevantData) bool {
	if record.Team == "" || record.Opponent == "" {
		return false
	}

	for _, r := range recordList {
		if r.Player == record.Player && r.Sport == record.Sport && r.StatType == record.StatType && r.Time.Equal(record.Time) {
			return true
		}
	}
	return false
}
