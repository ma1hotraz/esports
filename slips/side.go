package slips

import (
	"github.com/iloginow/esportsdifference/compare"
	"github.com/iloginow/esportsdifference/esport"
)

type Side struct {
	A compare.Record
	B compare.Record
}

func (s Side) match() bool {
	// LOL case
	if s.A.Sport == esport.LOL && s.B.Sport == esport.LOL {
		if s.A.Underdog > s.A.PrizePicks && s.B.Underdog < s.B.PrizePicks {
			return true
		}
		if s.A.Underdog < s.A.PrizePicks && s.B.Underdog > s.B.PrizePicks {
			return true
		}
	} else {
		// not LOL case
		if s.A.Underdog > s.A.PrizePicks && s.B.Underdog > s.B.PrizePicks {
			return true
		}
		if s.A.Underdog < s.A.PrizePicks && s.B.Underdog < s.B.PrizePicks {
			return true
		}
	}

	return false
}
