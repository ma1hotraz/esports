package underdog

import (
	"strconv"
	"strings"

	"github.com/iloginow/esportsdifference/dto"
	"github.com/iloginow/esportsdifference/esport"
)

type Data struct {
	Appearances      []Appearance    `json:"appearances"`
	Games            []Game          `json:"games"`
	OpenedLinesCount int             `json:"opened_lines_count"`
	OverUnderLines   []OverUnderLine `json:"over_under_lines"`
	Players          []Player        `json:"players"`
	SoloGames        []SoloGame      `json:"solo_games"`
}

func (d *Data) Filter() []dto.UnderdogRelevantData {
	data := []dto.UnderdogRelevantData{}
	for _, l := range d.OverUnderLines {
		a := d.getAppearance(l.OverUnder.AppearanceStat.AppearanceId)
		p := d.getPlayer(a.PlayerId)
		g := d.getGame(a.MatchId)
		league := l.OverUnder.GetLeague()

		if league.Type == "" {
			sportId := strings.ToLower(g.SportId)
			otherLeague, ok := esport.RelevantLeaguesMap[sportId]
			if ok {
				league = otherLeague
			}
		}

		if l.OverUnder.IsRelevant() || league.Type != "" {
			prop := l.OverUnder.GetProp(league.PlayerProps)
			team, opponent := g.GetTeamNamesByPlayer(p)
			val, err := strconv.ParseFloat(l.StatValue, 64)
			if err == nil {
				data = append(data, dto.UnderdogRelevantData{
					Player:   p.LastName,
					Time:     g.ScheduledAt,
					Sport:    league.Type,
					Team:     team,
					Opponent: opponent,
					StatType: prop.Type,
					Value:    val,
				})
			}
		}
	}
	return data
}

func (d *Data) getAppearance(id string) Appearance {
	var appearance Appearance
	for _, a := range d.Appearances {
		if a.Id == id {
			return a
		}
	}
	return appearance
}

func (d *Data) getPlayer(id string) Player {
	var player Player
	for _, p := range d.Players {
		if p.Id == id {
			return p
		}
	}
	return player
}

func (d *Data) getGame(id int) Game {
	var game Game
	for _, g := range d.Games {
		if g.Id == id {
			return g
		}
	}
	return game
}
