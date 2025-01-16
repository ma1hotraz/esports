package esport

import (
	"strings"
)

type LeagueType string

const (
	COD  LeagueType = "COD"
	CSGO LeagueType = "CSGO"
	LOL  LeagueType = "LOL"
	VAL  LeagueType = "VAL"
	DOTA LeagueType = "DOTA"
	HALO LeagueType = "HALO"
)

type League struct {
	Type        LeagueType
	Names       []string
	Ids         []string
	PlayerProps []PlayerProp
}

func (l League) RecognizeByName(name string) bool {
	for _, n := range l.Names {
		if strings.ToLower(name) == n {
			return true
		}
	}
	return false
}

func (l League) RecognizeById(id string) bool {
	for _, i := range l.Ids {
		if id == i {
			return true
		}
	}
	return false
}

func (l League) RecognizeByEventTitle(title string) bool {
	parts := strings.Split(title, ": ")
	if len(parts) < 2 {
		return false
	}
	for _, n := range l.Names {
		if strings.ToLower(parts[0]) == n {
			return true
		}
	}
	return false
}
