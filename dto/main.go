package dto

import (
	"time"

	"github.com/iloginow/esportsdifference/esport"
)

type UnderdogRelevantData struct {
	Player   string
	Time     time.Time
	Sport    esport.LeagueType
	Team     string
	Opponent string
	StatType esport.PlayerPropType
	Value    float64
}

type PrizePicksRelevantData struct {
	ProjectionId string
	Player       string
	Time         time.Time
	Sport        esport.LeagueType
	StatType     esport.PlayerPropType
	Value        float64
}

type SleeperRelevantData struct {
	Player     string
	Time       time.Time
	Sport      esport.LeagueType
	StatType   esport.PlayerPropType
	Value      float64
	Multiplier float64
	OverUnder  string
}
