package update

import (
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/notifier"
	"github.com/iloginow/esportsdifference/prizepicks"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/sirupsen/logrus"
)

func getNewPrizepicksData() []prizepicks.RelevantData {
	rawData, err := prizepicks.GetData()
	if err != nil {
		logrus.Error(err)
	}
	return rawData.Filter()
}

func getPrevPrizepicksData() []prizepicks.RelevantData {
	data, err := repo.GetPrizepicksRelevant()
	if err != nil {
		logrus.Error(err)
	}
	return data
}

func countNewPrizepicksRecords(
	newRecords []prizepicks.RelevantData,
	oldRecords []prizepicks.RelevantData,
) notifier.NewLinesCount {
	nl := notifier.NewLinesCount{}
	if len(oldRecords) < 1 {
		return nl
	}
	for _, r := range newRecords {
		if !findRelevantPrizepicksRecord(r, oldRecords) {
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

func getNewPrizepicksRecords(
	newRecords []prizepicks.RelevantData,
	oldRecords []prizepicks.RelevantData,
) notifier.NewPrizepicksData {
	nl := notifier.NewPrizepicksData{}
	if len(oldRecords) < 1 {
		return nl
	}
	for _, r := range newRecords {
		if !findRelevantPrizepicksRecord(r, oldRecords) {
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

func findRelevantPrizepicksRecord(record prizepicks.RelevantData, recordList []prizepicks.RelevantData) bool {
	for _, r := range recordList {
		if r.Player == record.Player && r.Sport == record.Sport && r.StatType == record.StatType && r.Time.Equal(record.Time) {
			return true
		}
	}
	return false
}
