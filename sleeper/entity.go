package sleeper

import "github.com/iloginow/esportsdifference/esport"

type Data []InnerStruct

type InnerStruct struct {
	Status      Status      `json:"status"`
	Options     []Option    `json:"options"`
	SubjectID   string      `json:"subject_id"`
	Sport       Sport       `json:"sport"`
	Season      string      `json:"season"`
	SeasonType  SeasonType  `json:"season_type"`
	GameID      string      `json:"game_id"`
	SubjectType SubjectType `json:"subject_type"`
	WagerType   WagerType   `json:"wager_type"`
	OutcomeType OutcomeType `json:"outcome_type"`
	LineType    LineType    `json:"line_type"`
	GameStatus  string      `json:"game_status"`
	UpdatedAt   int64       `json:"updated_at"`
	MarketType  string      `json:"market_type"`
	PickStats   *PickStats  `json:"pick_stats,omitempty"`
}

type Option struct {
	Status           Status      `json:"status"`
	Metadata         Metadata    `json:"metadata"`
	SubjectID        string      `json:"subject_id"`
	Outcome          Outcome     `json:"outcome"`
	Sport            Sport       `json:"sport"`
	Season           string      `json:"season"`
	SeasonType       SeasonType  `json:"season_type"`
	GameID           string      `json:"game_id"`
	SubjectType      SubjectType `json:"subject_type"`
	WagerType        WagerType   `json:"wager_type"`
	OutcomeType      OutcomeType `json:"outcome_type"`
	LineType         LineType    `json:"line_type"`
	GameStatus       string      `json:"game_status"`
	LineID           string      `json:"line_id"`
	MarketType       string      `json:"market_type"`
	OutcomeValue     float64     `json:"outcome_value"`
	PayoutMultiplier string     `json:"payout_multiplier"`
}

type Metadata struct {
}

type PickStats struct {
	Counts     Counts  `json:"counts"`
	Popularity float64 `json:"popularity"`
}

type Counts struct {
	Over  *int64 `json:"over,omitempty"`
	Total int64  `json:"total"`
	Under *int64 `json:"under,omitempty"`
}

const (
	PreGame string = "pre_game"
)

type LineType string

const (
	Normal LineType = "normal"
)

type Outcome string

const (
	Over  Outcome = "over"
	Under Outcome = "under"
)

type OutcomeType string

const (
	OverUnder OutcomeType = "over_under"
)

type SeasonType string

const (
	Regular SeasonType = "regular"
)

type Sport string

const (
	CS Sport = "cs"
)

type Status string

const (
	Active Status = "active"
)

type SubjectType string

const (
	Player SubjectType = "player"
)

type WagerType string

const (
	HeadshotsMap1    WagerType = "headshots_map_1"
	HeadshotsMaps1_2 WagerType = "headshots_maps_1_2"
	KillsMap1        WagerType = "kills_map_1"
	KillsMaps1_2     WagerType = "kills_maps_1_2"
)

func (e InnerStruct) GetStatType() esport.PlayerPropType {
	var res esport.PlayerPropType
	switch e.WagerType {
	case HeadshotsMap1:
		res = esport.MAP_HEADSHOTS_1
	case HeadshotsMaps1_2:
		res = esport.MAP_HEADSHOTS_1_2
	case KillsMap1:
		res = esport.MAP_KILLS_1
	case KillsMaps1_2:
		res = esport.MAP_KILLS_1_2
	default:
		// Handle default case if necessary
	}
	return res
}
