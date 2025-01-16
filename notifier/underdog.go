package notifier

import (
	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/iloginow/esportsdifference/underdog"
	"github.com/sirupsen/logrus"
)

type NewUnderdogData struct {
	COD  []dto.UnderdogRelevantData
	CSGO []dto.UnderdogRelevantData
	LOL  []dto.UnderdogRelevantData
	VAL  []dto.UnderdogRelevantData
	DOTA []dto.UnderdogRelevantData
	HALO []dto.UnderdogRelevantData
}

func (nl NewUnderdogData) announce(t topic) {
	var cod, csgo, val, halo = []dto.UnderdogRelevantData{}, []dto.UnderdogRelevantData{}, []dto.UnderdogRelevantData{}, []dto.UnderdogRelevantData{}

	var codCnt, csgoCnt, valCnt, haloCount = 0, 0, 0, 0

	informedRecords, err := repo.GetInformedUnderdogLines()

	if err != nil {
		return
	}

	for _, c := range nl.COD {
		if !underdog.IsRecordInList(c, informedRecords.UnderdogLines) {
			codCnt = codCnt + 1
			cod = append(cod, c)
		} else {
			logrus.Debugf("Detected duplicated record informed in COD: %v", c)
		}
	}

	for _, c := range nl.CSGO {
		if !underdog.IsRecordInList(c, informedRecords.UnderdogLines) {
			csgoCnt = csgoCnt + 1
			csgo = append(csgo, c)
		} else {
			logrus.Debugf("Detected duplicated record informed in CSGO: %v", c)
		}
	}

	for _, c := range nl.VAL {
		if !underdog.IsRecordInList(c, informedRecords.UnderdogLines) {
			valCnt = valCnt + 1
			val = append(val, c)
		} else {
			logrus.Debugf("Detected duplicated record informed in VAL: %v", c)
		}
	}

	for _, c := range nl.HALO {
		if !underdog.IsRecordInList(c, informedRecords.UnderdogLines) {
			haloCount = haloCount + 1
			halo = append(halo, c)
		} else {
			logrus.Debugf("Detected duplicated record informed in HALO: %v", c)
		}
	}

	logrus.Infof("[NewUnderdogData] Original data COD %d, CSGO %d, VAL %d, HALO %d \n\t\t> Reduce to: COD %d, CSGO %d, VAL %d, HALO %d", len(nl.COD), len(nl.CSGO), len(nl.VAL), len(nl.HALO), codCnt, csgoCnt, valCnt, haloCount)

	if codCnt > 0 {
		announceNewLinesForLeagueV2(codCnt, esport.COD, t)
	}
	if csgoCnt > 0 {
		announceNewLinesForLeagueV2(csgoCnt, esport.CSGO, t)
	}
	if valCnt > 0 {
		announceNewLinesForLeagueV2(valCnt, esport.VAL, t)
	}

	if haloCount > 0 {
		announceNewLinesForLeagueV2(haloCount, esport.HALO, t)
	}

	// FIXME
	repo.SaveInformedUnderdog(cod, csgo, val, halo)
}
