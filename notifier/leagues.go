package notifier

import (
	"github.com/iloginow/esportsdifference/esport"
)

var leagueNames = map[esport.LeagueType]string{
	esport.COD:  "Call Of Duty",
	esport.CSGO: "CS2",
	esport.LOL:  "League of Legends",
	esport.VAL:  "Valorant",
	esport.DOTA: "Dota",
	esport.HALO: "Halo",
}
