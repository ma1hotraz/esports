package prizepicks

import (
	"time"

	"github.com/iloginow/esportsdifference/esport"
)

type RelevantData struct {
	ProjectionId string
	Player   string
	Time     time.Time
	Sport    esport.LeagueType
	StatType esport.PlayerPropType
	Value    float64
}
