package prizepicks

import "github.com/iloginow/esportsdifference/esport"

type EntityType string

const (
	PROJECTION      EntityType = "projection"
	PROJECTION_TYPE EntityType = "projection_type"
	LEAGUE          EntityType = "league"
	STAT_TYPE       EntityType = "stat_type"
	NEW_PLAYER      EntityType = "new_player"
	DURATION        EntityType = "duration"
)

type Entity struct {
	Type          EntityType    `json:"type"`
	Id            string        `json:"id"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

func (e Entity) GetLeague() esport.League {
	var league esport.League
	for _, l := range esport.RelevantLeagues {
		if l.RecognizeById(e.Relationships.League.Data.Id) {
			return l
		}
	}
	return league
}

func (e Entity) GetProp(relevantProps []esport.PlayerProp) esport.PlayerProp {
	var prop esport.PlayerProp
	for _, p := range relevantProps {
		if p.RecognizeByName(e.Attributes.StatType) {
			return p
		}
	}
	return prop
}

func (e Entity) IsRelevant() bool {
	if e.Type != PROJECTION {
		return false
	}
	league := e.GetLeague()
	if len(league.Names) > 0 && len(league.PlayerProps) > 0 {
		prop := e.GetProp(league.PlayerProps)
		if len(prop.Names) > 0 {
			return true
		}
	}
	return false
}
