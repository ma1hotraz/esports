package repo

import (
	"github.com/iloginow/esportsdifference/compare"
	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/prizepicks"
	"github.com/sirupsen/logrus"
)

const INFORMED_LINES_KEY = "INFORMED_LINES"
const INFORMED_UNDERDOG_KEY = "INFORMED_UNDERDOG_KEY"
const INFORMED_PRIZE_KEY = "INFORMED_PRIZE_KEY"

type InformedLines struct {
	CompareLines    []compare.Record           `json:"compare_pair_lines"`
	UnderdogLines   []dto.UnderdogRelevantData `json:"underdog_lines"`
	PrizepicksLines []prizepicks.RelevantData  `json:"prizepicks_lines"`
}

func SaveInformedRecoredLines(cod []compare.Record, csgo []compare.Record, val []compare.Record, halo []compare.Record) error {
	existedLines, err := GetInformedCompareLines()
	if err == nil {
		newLines := compare.MergeRecords(existedLines.CompareLines, cod)
		newLines = append(newLines, csgo...)
		newLines = append(newLines, val...)
		newLines = append(newLines, halo...)

		if len(cod) > 0 || len(csgo) > 0 || len(val) > 0 || len(halo) > 0 {
			logrus.Infof("Saving new compare informed lines: cod %d, csgo %d, val %d, halo %d", len(cod), len(csgo), len(val), len(halo))
		}
		return saveJsonData(INFORMED_LINES_KEY, InformedLines{
			CompareLines: newLines,
		})
	}

	return nil
}

func SaveInformedUnderdog(cod []dto.UnderdogRelevantData, csgo []dto.UnderdogRelevantData, val []dto.UnderdogRelevantData, halo []dto.UnderdogRelevantData) error {
	existedLines, err := GetInformedUnderdogLines()
	if err == nil {
		newLines := []dto.UnderdogRelevantData{}
		newLines = append(newLines, (existedLines.UnderdogLines)...)
		newLines = append(newLines, cod...)
		newLines = append(newLines, csgo...)
		newLines = append(newLines, val...)
		newLines = append(newLines, halo...)
		if len(cod) > 0 || len(csgo) > 0 || len(val) > 0 || len(halo) > 0 {
			logrus.Infof("Saving new underdog informed lines: cod %d, csgo %d, val %d, halo %d", len(cod), len(csgo), len(val), len(halo))
		}
		return saveJsonData(INFORMED_UNDERDOG_KEY, InformedLines{
			UnderdogLines: newLines,
		})
	}

	return nil
}

func SaveInformedPrize(cod []prizepicks.RelevantData, csgo []prizepicks.RelevantData, val []prizepicks.RelevantData, halo []prizepicks.RelevantData, lol []prizepicks.RelevantData) error {
	existedLines, err := GetInformedPrizepricksLines()
	if err == nil {
		newLines := []prizepicks.RelevantData{}
		newLines = append(newLines, (existedLines.PrizepicksLines)...)
		newLines = append(newLines, cod...)
		newLines = append(newLines, csgo...)
		newLines = append(newLines, val...)
		newLines = append(newLines, halo...)
		newLines = append(newLines, lol...)

		if len(cod) > 0 || len(csgo) > 0 || len(val) > 0 || len(halo) > 0 || len(lol) > 0 {
			logrus.Infof("Saving new prize informed lines: cod %d, csgo %d, val %d, halo %d, lol %d", len(cod), len(csgo), len(val), len(halo), len(lol))
		}
		return saveJsonData(INFORMED_PRIZE_KEY, InformedLines{
			PrizepicksLines: newLines,
		})
	}
	return nil
}

func GetInformedCompareLines() (InformedLines, error) {
	var lines InformedLines
	if err := getJsonData(INFORMED_LINES_KEY, &lines); err != nil {
		return lines, err
	}
	return lines, nil
}

func GetInformedUnderdogLines() (InformedLines, error) {
	var lines InformedLines
	if err := getJsonData(INFORMED_UNDERDOG_KEY, &lines); err != nil {
		return lines, err
	}
	return lines, nil
}

func GetInformedPrizepricksLines() (InformedLines, error) {
	var lines InformedLines
	if err := getJsonData(INFORMED_PRIZE_KEY, &lines); err != nil {
		return lines, err
	}
	return lines, nil
}

func CleanUpInformedData() {
	removeJsonData(INFORMED_LINES_KEY)
	removeJsonData(INFORMED_UNDERDOG_KEY)
	removeJsonData(INFORMED_PRIZE_KEY)
}

func GetAllInformedLines() (InformedLines, error) {
	underdog, _ := GetInformedUnderdogLines()
	prize, _ := GetInformedPrizepricksLines()
	compare, _ := GetInformedCompareLines()

	return InformedLines{
		UnderdogLines:   underdog.UnderdogLines,
		PrizepicksLines: prize.PrizepicksLines,
		CompareLines:    compare.CompareLines,
	}, nil
}
