package underdog

import (
	"time"

	"github.com/iloginow/esportsdifference/esport"
)

type OverUnderLine struct {
	ExpiresAt     time.Time             `json:"expires_at"`
	Id            string                `json:"id"`
	LiveEvent     bool                  `json:"live_event"`
	LiveEventStat interface{}           `json:"live_event_stat"`
	Options       []OverUnderLineOption `json:"options"`
	OverUnder     OverUnder             `json:"over_under"`
	OverUnderId   string                `json:"over_under_id"`
	Rank          int                   `json:"rank"`
	StatValue     string                `json:"stat_value"`
	Status        string                `json:"status"`
}

type OverUnderLineOption struct {
	Choice           string `json:"choice"`
	ChoiceDisplay    string `json:"choice_display"`
	Id               string `json:"id"`
	OverUnderLineId  string `json:"over_under_line_id"`
	PayoutMultiplier string `json:"payout_multiplier"`
	Type             string `json:"type"`
}

type OverUnder struct {
	AppearanceStat AppearanceStat `json:"appearance_stat"`
	Boost          string         `json:"boost"`
	Id             string         `json:"id"`
	OptionPriority string         `json:"option_priority"`
	ScoringTypeId  string         `json:"scoring_type_id"`
	Title          string         `json:"title"`
}

func (ou OverUnder) GetLeague() esport.League {
	var league esport.League

	for _, l := range esport.RelevantLeagues {
		if l.RecognizeByEventTitle(ou.Title) {
			return l
		}
	}

	return league
}

func (ou OverUnder) GetProp(relevantProps []esport.PlayerProp) esport.PlayerProp {
	var prop esport.PlayerProp
	for _, p := range relevantProps {
		if p.RecognizeByName(ou.AppearanceStat.DisplayStat) {
			return p
		}
	}
	return prop
}

func (ou OverUnder) IsRelevant() bool {
	league := ou.GetLeague()
	if len(league.Names) > 0 && len(league.PlayerProps) > 0 {
		prop := ou.GetProp(league.PlayerProps)
		if len(prop.Names) > 0 {
			return true
		}
	}
	return false
}
