package slips

import (
	"sort"
	"strings"

	"github.com/iloginow/esportsdifference/compare"
	"github.com/iloginow/esportsdifference/esport"
)

type Matches struct {
	Elements []compare.Record
	Used     []compare.Record
}

func NewMatches(cr []compare.Record) *Matches {
	matches := new(Matches)
	matches.Elements = cr
	if matches.Elements != nil {
		sort.Slice(matches.Elements, func(i, j int) bool {
			return matches.Elements[i].Difference > matches.Elements[j].Difference
		})
	}
	return matches
}

func (m *Matches) Find() [][]compare.Record {
	paired := [][]compare.Record{}
	for _, r := range m.Elements {
		if m.findUsed(r) {
			continue
		}
		pair := m.findPair(r)
		if pair != nil {
			paired = append(paired, pair)
		}
	}
	// sort pari by difference
	sort.Slice(paired, func(i, j int) bool {
		return paired[i][0].Difference+paired[i][1].Difference > paired[j][0].Difference+paired[j][1].Difference
	})
	return paired
}

// Handle get pairs

func compareRecord(i, r compare.Record) bool {
	if i.Team == r.Opponent && i.Opponent == r.Team && i.Timestamp.Equal(r.Timestamp) {
		if i.StatType == r.StatType || strings.Contains(string(r.StatType), string(i.StatType)) || strings.Contains(string(r.StatType), string(i.StatType)) {
			return true
		}
	}
	return false
}

func (m *Matches) findPair(r compare.Record) []compare.Record {
	for _, i := range m.Elements {
		if m.findUsed(i) {
			continue
		}

		if i.Sport == esport.LOL && r.Sport == esport.LOL {
			if i.Player != r.Player && i.Team == r.Opponent && i.Opponent == r.Team && i.Timestamp.Equal(r.Timestamp) {
				s := Side{r, i}
				if s.match() {
					m.use(r)
					m.use(i)
					return []compare.Record{r, i}
				}
			}
		} else if i.Team == r.Opponent && i.Opponent == r.Team && i.Timestamp.Equal(r.Timestamp) {
			if i.StatType == r.StatType || strings.Contains(string(r.StatType), string(i.StatType)) || strings.Contains(string(r.StatType), string(i.StatType)) {
				s := Side{r, i}
				if s.match() {
					m.use(r)
					m.use(i)
					return []compare.Record{r, i}
				}
			}
		}
	}
	return nil
}

func (m *Matches) use(r compare.Record) {
	m.Used = append(m.Used, r)
}

func (m *Matches) findUsed(r compare.Record) bool {
	for _, u := range m.Used {
		if u.Player == r.Player && u.Timestamp.Equal(r.Timestamp) && u.Team == r.Team && u.Opponent == r.Opponent {
			return true
		}
	}
	return false
}
