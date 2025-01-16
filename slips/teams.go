package slips

import "github.com/iloginow/esportsdifference/compare"

type Teams struct {
	Names  []string
	Lookup map[string]bool
}

func NewTeams(p []compare.Record) *Teams {
	teams := new(Teams)
	teams.Lookup = make(map[string]bool)

	for _, record := range p {
		// Add the team and opponent to Names if they are not already in Lookup
		if !teams.Lookup[record.Team] {
			teams.Names = append(teams.Names, record.Team)
			teams.Lookup[record.Team] = true
		}
		if !teams.Lookup[record.Opponent] {
			teams.Names = append(teams.Names, record.Opponent)
			teams.Lookup[record.Opponent] = true
		}
	}

	return teams
}

func (t *Teams) match(teams *Teams) bool {
	for _, n := range teams.Names {
		if t.Lookup[n] {
			return true
		}
	}

	return false
}
