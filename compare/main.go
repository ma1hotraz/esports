package compare

import (
	"fmt"
	"math"
	"strings"
	"time"
	"unicode"

	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/esport"
	"github.com/iloginow/esportsdifference/prizepicks"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Result struct {
	COD  []Record `json:"cod"`
	CSGO []Record `json:"csgo"`
	LOL  []Record `json:"lol"`
	VAL  []Record `json:"val"`
	DOTA []Record `json:"dota"`
	HALO []Record `json:"halo"`
}

type Record struct {
	Difference        float64               `json:"difference"`
	PercentDifference float64               `json:"percent_difference"`
	Player            string                `json:"player"`
	Sport             esport.LeagueType     `json:"sport"`
	Team              string                `json:"team"`
	Opponent          string                `json:"opponent"`
	StatType          esport.PlayerPropType `json:"stat_type"`
	Timestamp         time.Time             `json:"timestamp"`
	IsNew             bool                  `json:"is_new"`

	PrizePicks       float64 `json:"prize_picks"`
	ProjectionString string  `json:"projection_string"`

	SleeperOverUnder  string  `json:"sleeper_over_under"`
	Sleeper           float64 `json:"sleeper"`
	SleeperMultiplier float64 `json:"sleeper_multiplier"`

	Underdog float64 `json:"underdog"`
}

func CompareUnderdogToPrize(ud []dto.UnderdogRelevantData, pd []prizepicks.RelevantData, sd []dto.SleeperRelevantData) Result {
	underdog_pp := underdog_merge_prizepicks(ud, pd)
	underdog_sp := underdog_merge_Sleeper(ud, sd)
	underdog_pp = mergeCombine(underdog_pp, underdog_sp)
	return sort(underdog_pp)
}

func CompareUnderdogToSleeper(ud []dto.UnderdogRelevantData, pd []dto.SleeperRelevantData) Result {
	mergedData := underdog_merge_Sleeper(ud, pd)
	return sort(mergedData)
}

func (r Result) IsNotEmpty() bool {
	return len(r.COD) > 0 || len(r.CSGO) > 0 || len(r.LOL) > 0 || len(r.VAL) > 0 || len(r.DOTA) > 0 || len(r.HALO) > 0
}

func underdog_merge_prizepicks(underdogList []dto.UnderdogRelevantData, prizepicksList []prizepicks.RelevantData) []Record {
	result := []Record{}
	for _, underdogLine := range underdogList {
		prizepicksLine := findPrizepicksRecord(underdogLine, prizepicksList)
		if prizepicksLine.Value > 0 && getAbsoluteDifference(underdogLine.Value, prizepicksLine.Value) > 0 {
			if underdogLine.Team == "" || underdogLine.Opponent == "" {
				// logrus.Errorf("%v is having empty team name", underdogLine)
			} else {
				result = append(result, mergeRelevantData(prizepicksLine, underdogLine))
			}
		}
	}

	// handle combo line case
	for _, prizepicksLine := range prizepicksList {
		if esport.IsComboType(prizepicksLine.StatType) {
			comboLineRecord := findComboLineRecordFromUnderdog(prizepicksLine, underdogList)
			if comboLineRecord.Difference > 0 && comboLineRecord.Team != "" && comboLineRecord.Opponent != "" {
				result = append(result, comboLineRecord)
			}
		}
	}

	return result
}

func underdog_merge_Sleeper(underdogList []dto.UnderdogRelevantData, sleeperList []dto.SleeperRelevantData) []Record {
	result := []Record{}
	for _, underdogLine := range underdogList {
		sleeper := findSleeperRecord(underdogLine, sleeperList)
		if underdogLine.Player == "j3kie" {
			logrus.Debugf("underdogLine: %v, sleeper: %v", underdogLine, sleeper)
		}
		if sleeper.Value > 0 {
			if underdogLine.Team == "" || underdogLine.Opponent == "" {
				// logrus.Errorf("%v is having empty team name", underdogLine)
			} else {
				result = append(result, mergeRelevantDataSleeper(sleeper, underdogLine))
			}
		}
	}

	return result
}

func mergeCombine(udToPp []Record, udToSP []Record) []Record {
	res := []Record{}
	for _, udToPpRecord := range udToPp {
		tmp := udToPpRecord
		// if r is in udToPp
		if s := GetRecordInList(udToPpRecord, udToSP); s.Team != "" {
			if tmp.Player == "j3kie" {
				logrus.Debugf("udToPpRecord: %v, udToSP: %v", tmp, s)
			}
			tmp.Sleeper = s.Sleeper
			tmp.SleeperMultiplier = s.SleeperMultiplier
			tmp.SleeperOverUnder = s.SleeperOverUnder
			res = append(res, tmp)
		} else {
			res = append(res, tmp)
		}
	}
	return res
}

func findComboLineRecordFromUnderdog(prizepicksComboLine prizepicks.RelevantData, underdogList []dto.UnderdogRelevantData) Record {
	var result Record

	if !esport.IsComboType(prizepicksComboLine.StatType) {
		return result
	}

	subPlayers := strings.Split(prizepicksComboLine.Player, " + ")
	if len(subPlayers) < 1 {
		return result
	}

	var records []dto.UnderdogRelevantData
	for _, player := range subPlayers {
		var temp = prizepicksComboLine
		temp.Player = player
		record := findUnderdogRecord(temp, underdogList)
		records = append(records, record)
	}
	var totalValue float64
	for _, record := range records {
		if record.Value == 0 {
			return result
		}
		totalValue += record.Value
	}

	return mergeRelevantData(prizepicksComboLine, dto.UnderdogRelevantData{
		Player:   prizepicksComboLine.Player,
		Time:     records[0].Time,
		Sport:    prizepicksComboLine.Sport,
		Team:     records[0].Team,
		Opponent: records[0].Opponent,
		StatType: prizepicksComboLine.StatType,
		Value:    totalValue,
	})
}

func sort(d []Record) Result {
	var result Result
	for _, r := range d {
		if r.Sport == esport.COD {
			result.COD = append(result.COD, r)
		} else if r.Sport == esport.CSGO {
			result.CSGO = append(result.CSGO, r)
		} else if r.Sport == esport.LOL {
			result.LOL = append(result.LOL, r)
		} else if r.Sport == esport.VAL {
			result.VAL = append(result.VAL, r)
		} else if r.Sport == esport.DOTA {
			result.DOTA = append(result.DOTA, r)
		} else if r.Sport == esport.HALO {
			result.HALO = append(result.HALO, r)
		}
	}
	return result
}

func findPrizepicksRecord(u dto.UnderdogRelevantData, pd []prizepicks.RelevantData) prizepicks.RelevantData {
	var result prizepicks.RelevantData
	for _, r := range pd {
		if compareRecord(u, r) {
			return r
		}
	}
	return result
}
func findSleeperRecord(u dto.UnderdogRelevantData, pd []dto.SleeperRelevantData) dto.SleeperRelevantData {

	var result dto.SleeperRelevantData
	for _, r := range pd {
		if compareSleeperRecord(u, r) {
			return r
		}
	}
	return result
}

func findUnderdogRecord(p prizepicks.RelevantData, ud []dto.UnderdogRelevantData) dto.UnderdogRelevantData {
	var result dto.UnderdogRelevantData
	for _, r := range ud {
		if compareRecord(r, p) {
			return r
		}
	}
	return result
}

func mergeRelevantData(prizepicksLine prizepicks.RelevantData, underdogLine dto.UnderdogRelevantData) Record {
	overOrUnder := "o"
	if prizepicksLine.Value > underdogLine.Value {
		overOrUnder = "u"
	}
	return Record{
		Difference:        getAbsoluteDifference(underdogLine.Value, prizepicksLine.Value),
		PercentDifference: getPercentDifference(underdogLine.Value, prizepicksLine.Value),
		Player:            underdogLine.Player,
		Sport:             underdogLine.Sport,
		Team:              underdogLine.Team,
		Opponent:          underdogLine.Opponent,
		StatType:          underdogLine.StatType,
		PrizePicks:        prizepicksLine.Value,
		Underdog:          underdogLine.Value,
		Timestamp:         underdogLine.Time,
		IsNew:             false,
		ProjectionString:  fmt.Sprintf("%s-%s-%.1f", prizepicksLine.ProjectionId, overOrUnder, prizepicksLine.Value),
	}
}

func mergeRelevantDataSleeper(sleeperLine dto.SleeperRelevantData, underdogLine dto.UnderdogRelevantData) Record {
	return Record{
		Difference:        getAbsoluteDifference(underdogLine.Value, sleeperLine.Value),
		PercentDifference: getPercentDifference(underdogLine.Value, sleeperLine.Value),
		Player:            underdogLine.Player,
		Sport:             underdogLine.Sport,
		Team:              underdogLine.Team,
		Opponent:          underdogLine.Opponent,
		StatType:          underdogLine.StatType,

		Sleeper:           sleeperLine.Value,
		SleeperMultiplier: sleeperLine.Multiplier,
		SleeperOverUnder:  sleeperLine.OverUnder,

		Underdog:  underdogLine.Value,
		Timestamp: underdogLine.Time,
		IsNew:     false,
	}
}

var normalizeMapping = map[rune]string{
	'ø': "o", 'Ø': "O",
	'å': "a", 'Å': "A",
	'é': "e", 'É': "E",
	'ä': "a", 'Ä': "A",
	'ö': "o", 'Ö': "O",
	'ü': "u", 'Ü': "U",
	'ß': "ss",
}

func normalizeString(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if replacement, found := normalizeMapping[r]; found {
			builder.WriteString(replacement)
		} else if unicode.IsPrint(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func comparePlayerNames(a string, b string) bool {

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	norm_a, _, _ := transform.String(t, a)
	norm_b, _, _ := transform.String(t, b)
	// trim spaces
	norm_a = strings.TrimSpace(norm_a)
	norm_b = strings.TrimSpace(norm_b)

	return strings.EqualFold(normalizeString(norm_a), normalizeString(norm_b))
}

func compareRecord(u dto.UnderdogRelevantData, p prizepicks.RelevantData) bool {
	if u.Sport == p.Sport && (u.StatType != "" && strings.Contains(string(p.StatType), string(u.StatType)) && comparePlayerNames(u.Player, p.Player)) {
		if u.Time.Equal(p.Time) {
			return true

		} else if u.Sport == esport.COD && u.Time.UTC().Day() == p.Time.UTC().Day() && u.Time.UTC().Month() == p.Time.UTC().Month() && u.Time.UTC().Year() == p.Time.UTC().Year() {
			return true
		}
	}
	return false
}

func compareSleeperRecord(u dto.UnderdogRelevantData, p dto.SleeperRelevantData) bool {
	if u.Sport == p.Sport && (u.StatType != "" && strings.Contains(string(p.StatType), string(u.StatType)) && comparePlayerNames(u.Player, p.Player)) {
		if u.Time.Equal(p.Time) {
			return true
		} else if u.Sport == esport.COD && u.Time.UTC().Day() == p.Time.UTC().Day() && u.Time.UTC().Month() == p.Time.UTC().Month() && u.Time.UTC().Year() == p.Time.UTC().Year() {
			return true
		}
	}
	return false
}
func roundToTwoDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}

func getAbsoluteDifference(a, b float64) float64 {
	return roundToTwoDecimal(math.Abs(a - b))
}

func getPercentDifference(a, b float64) float64 {
	if a == 0 && b == 0 {
		return 0
	}
	return roundToTwoDecimal((math.Abs(a-b) / ((a + b) / 2)) * 100)
}
