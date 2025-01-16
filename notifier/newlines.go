package notifier

import (
	"strconv"

	"github.com/iloginow/esportsdifference/compare"
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/repo"
	"github.com/sirupsen/logrus"
)

type NewLinesCount struct {
	COD  int
	CSGO int
	VAL  int
	HALO int
}

type NewLinesData struct {
	COD  []compare.Record
	CSGO []compare.Record
	VAL  []compare.Record
	HALO []compare.Record
}

func (nl NewLinesCount) GetTotal() int {
	return nl.COD + nl.CSGO + nl.VAL + nl.HALO
}

func (nl NewLinesData) announce(t topic) {
	var cod, csgo, val, halo = []compare.Record{}, []compare.Record{}, []compare.Record{}, []compare.Record{}
	var codCnt, csgoCnt, valCnt, haloCnt = 0, 0, 0, 0

	informedRecords, err := repo.GetInformedCompareLines()

	if err != nil {
		return
	}

	for _, c := range nl.COD {
		if !compare.IsRecordInList(c, informedRecords.CompareLines) {
			codCnt = codCnt + 1
			cod = append(cod, c)
		}
	}

	for _, c := range nl.CSGO {
		if !compare.IsRecordInList(c, informedRecords.CompareLines) {
			csgoCnt = csgoCnt + 1
			csgo = append(csgo, c)
		}
	}

	for _, c := range nl.VAL {
		if !compare.IsRecordInList(c, informedRecords.CompareLines) {
			valCnt = valCnt + 1
			val = append(val, c)
		}
	}

	for _, c := range nl.HALO {
		if !compare.IsRecordInList(c, informedRecords.CompareLines) {
			haloCnt = haloCnt + 1
			halo = append(halo, c)
		}
	}
	logrus.Infof("[COMPARE] Original data COD %d, CSGO %d, VAL %d, HALO %d \n\t\t> Reduce to: COD %d, CSGO %d, VAL %d, HALO %d", len(nl.COD), len(nl.CSGO), len(nl.VAL), len(nl.HALO), codCnt, csgoCnt, valCnt, haloCnt)

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

	// FIXME
	repo.SaveInformedRecoredLines(cod, csgo, val, halo)
}

func announceNewLinesForLeagueV2(c int, l esport.LeagueType, t topic) {
	var description string
	if c == 0 {
		return
	}
	if c > 1 {
		description = topicDescriptionsPlural[t]
	} else {
		description = topicDescriptionsSingular[t]
	}
	msg := strconv.Itoa(c) + " new " + leagueNames[l] + " " + description
	SendDiscordMessageToAllBotV2(msg)
}
