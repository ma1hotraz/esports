package notifier

import (
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/prizepicks"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/sirupsen/logrus"
)

type NewPrizepicksData struct {
	COD  []prizepicks.RelevantData
	CSGO []prizepicks.RelevantData
	LOL  []prizepicks.RelevantData
	VAL  []prizepicks.RelevantData
	DOTA []prizepicks.RelevantData
	HALO []prizepicks.RelevantData
}

func (nl NewPrizepicksData) announce(t topic) {
	var cod, csgo, val, halo, lol = []prizepicks.RelevantData{}, []prizepicks.RelevantData{}, []prizepicks.RelevantData{}, []prizepicks.RelevantData{}, []prizepicks.RelevantData{}

	var codCnt, csgoCnt, valCnt, haloCnt, lolCnt = 0, 0, 0, 0, 0

	informedRecords, err := repo.GetInformedPrizepricksLines()

	if err != nil {
		return
	}

	for _, c := range nl.COD {
		if !prizepicks.IsRecordInList(c, informedRecords.PrizepicksLines) {
			codCnt = codCnt + 1
			cod = append(cod, c)
		}
	}

	for _, c := range nl.CSGO {
		if !prizepicks.IsRecordInList(c, informedRecords.PrizepicksLines) {
			csgoCnt = csgoCnt + 1
			csgo = append(csgo, c)
		}
	}

	for _, c := range nl.VAL {
		if !prizepicks.IsRecordInList(c, informedRecords.PrizepicksLines) {
			valCnt = valCnt + 1
			val = append(val, c)
		}
	}

	for _, c := range nl.HALO {
		if !prizepicks.IsRecordInList(c, informedRecords.PrizepicksLines) {
			haloCnt = haloCnt + 1
			halo = append(halo, c)
		}
	}

	for _, c := range nl.LOL {
		if !prizepicks.IsRecordInList(c, informedRecords.PrizepicksLines) {
			lolCnt = lolCnt + 1
			lol = append(lol, c)
		}
	}
	logrus.Infof("[NewPrizepicksData] Original data COD %d, CSGO %d, VAL %d, HALO %d \n\t\t> Reduce to: COD %d, CSGO %d, VAL %d, HALO %d", len(nl.COD), len(nl.CSGO), len(nl.VAL), len(nl.HALO), codCnt, csgoCnt, valCnt, haloCnt)

	if codCnt > 0 {
		announceNewLinesForLeagueV2(codCnt, esport.COD, t)
	}
	if csgoCnt > 0 {
		announceNewLinesForLeagueV2(csgoCnt, esport.CSGO, t)
	}
	if valCnt > 0 {
		announceNewLinesForLeagueV2(valCnt, esport.VAL, t)
	}

	if haloCnt > 0 {
		announceNewLinesForLeagueV2(haloCnt, esport.HALO, t)
	}

	if lolCnt > 0 {
		announceNewLinesForLeagueV2(lolCnt, esport.LOL, t)
	}

	// FIXME
	repo.SaveInformedPrize(cod, csgo, val, halo, lol)
}
