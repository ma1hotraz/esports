package esport

var RelevantLeagues = []League{
	{COD, []string{"cod"}, []string{"145"}, RelevantPropsFull},
	{CSGO, []string{"cs", "cs2", "csgo", "cs:go"}, []string{"265"}, RelevantPropsFull},
	{LOL, []string{"lol"}, []string{"121"}, RelevantPropsFull},
	{VAL, []string{"val"}, []string{"159"}, RelevantPropsFull},
	{DOTA, []string{"dota", "dota2"}, []string{"174"}, RelevantPropsFull},
	{HALO, []string{"halo", "halo2"}, []string{"267"}, RelevantPropsFull},
}

var (
	RelevantLeaguesMap = map[string]League{
		"lol":  RelevantLeagues[2],
		"val":  RelevantLeagues[3],
		"halo": RelevantLeagues[5],
		"cs":   RelevantLeagues[1],
	}
)

var (
	MapKills1 = PlayerProp{
		MAP_KILLS_1,
		[]string{"kills in game 1", "kills on game 1", "kills in map 1", "kills on map 1", "map 1 kills"},
	}
	MapKills2 = PlayerProp{
		MAP_KILLS_2,
		[]string{"kills in game 2", "kills on game 2", "kills in map 2", "kills on map 2", "map 2 kills"},
	}
	MapKills3 = PlayerProp{
		MAP_KILLS_3,
		[]string{"kills in game 3", "kills on game 3", "kills in map 3", "kills on map 3", "map 3 kills"},
	}
	MapKills12 = PlayerProp{
		MAP_KILLS_1_2,
		[]string{"kills in game 1+2", "kills on game 1+2", "kills in maps 1+2", "kills on map 1+2", "kills on maps 1+2", "maps 1-2 kills"},
	}
	MapKills13 = PlayerProp{
		MAP_KILLS_1_3,
		[]string{"kills in game 1+2+3", "kills on game 1+2+3", "kills in maps 1+2+3", "kills on map 1+2+3", "kills on maps 1+2+3", "maps 1-3 kills"},
	}

	MapAssists1 = PlayerProp{
		MAP_ASSISTS_1,
		[]string{"assists in game 1", "assists on game 1", "assists in map 1", "assists on map 1", "map 1 assists"},
	}
	MapAssists2 = PlayerProp{
		MAP_ASSISTS_2,
		[]string{"assists in game 2", "assists on game 2", "assists in map 2", "assists on map 2", "map 2 assists"},
	}
	MapAssists3 = PlayerProp{
		MAP_ASSISTS_3,
		[]string{"assists in game 3", "assists on game 3", "assists in map 3", "assists on map 3", "map 3 assists"},
	}

	MapAssists12 = PlayerProp{
		MAP_ASSISTS_1_2,
		[]string{"assists in game 1+2", "assists on game 1+2", "assists in maps 1+2", "assists on map 1+2", "assists on maps 1+2", "maps 1-2 assists"},
	}

	MapAssists13 = PlayerProp{
		MAP_ASSISTS_1_3,
		[]string{"assists in game 1+2+3", "assists on game 1+2+3", "assists in maps 1+2+3", "assists on map 1+2+3", "assists on maps 1+2+3", "maps 1-3 assists"},
	}

	MapHeadshots1 = PlayerProp{
		MAP_HEADSHOTS_1,
		[]string{"headshots in game 1", "headshots on game 1", "headshots in map 1", "headshots on map 1", "map 1 headshots"},
	}
	MapHeadshots2 = PlayerProp{
		MAP_HEADSHOTS_2,
		[]string{"headshots in game 2", "headshots on game 2", "headshots in map 2", "headshots on map 2", "map 2 headshots"},
	}
	MapHeadshots3 = PlayerProp{
		MAP_HEADSHOTS_3,
		[]string{"headshots in game 3", "headshots on game 3", "headshots in map 3", "headshots on map 3", "map 3 headshots"},
	}
	MapHeadshots12 = PlayerProp{
		MAP_HEADSHOTS_1_2,
		[]string{"headshots in game 1+2", "headshots on game 1+2", "headshots in maps 1+2", "headshots on maps 1+2", "maps 1-2 headshots"},
	}
	MapHeadshots13 = PlayerProp{
		MAP_HEADSHOTS_1_3,
		[]string{"headshots in game 1+2+3", "headshots on game 1+2+3", "headshots in maps 1+2+3", "headshots on maps 1+2+3", "maps 1-3 headshots"},
	}

	MapKills1Combo = PlayerProp{
		MAP_KILLS_1_COMBO,
		[]string{"kills in game 1 (combo)", "kills on game 1 (combo)", "kills in map 1 combo", "kills on map 1 combo", "map 1 kills (combo)"},
	}
	MapKills2Combo = PlayerProp{
		MAP_KILLS_2_COMBO,
		[]string{"kills in game 2 (combo)", "kills on game 2 (combo)", "kills in map 2 combo", "kills on map 2 combo", "map 2 kills (combo)"},
	}

	MapKills3Combo = PlayerProp{
		MAP_KILLS_3_COMBO,
		[]string{"kills in game 3 (combo)", "kills on game 3 (combo)", "kills in map 3 combo", "kills on map 3 combo", "map 3 kills (combo)"},
	}

	MapKills12Combo = PlayerProp{
		MAP_KILLS_1_2_COMBO,
		[]string{"kills in game 1+2 (combo)", "kills on game 1+2 (combo)", "kills in maps 1+2 combo", "kills on maps 1+2 combo", "maps 1-2 kills (combo)"},
	}

	MapKills13Combo = PlayerProp{
		MAP_KILLS_1_3_COMBO,
		[]string{"kills in game 1+2+3 (combo)", "kills on game 1+2+3 (combo)", "kills in maps 1+2+3 combo", "kills on maps 1+2+3 combo", "maps 1-3 kills (combo)"},
	}

	RelevantPropsFull = []PlayerProp{
		MapKills1,
		MapKills2,
		MapKills3,
		MapKills12,
		MapKills13,
		MapHeadshots1,
		MapHeadshots2,
		MapHeadshots3,
		MapHeadshots12,
		MapHeadshots13,
		MapAssists1,
		MapAssists2,
		MapAssists3,
		MapAssists12,
		MapAssists13,

		MapKills1Combo,
		MapKills2Combo,
		MapKills3Combo,
		MapKills12Combo,
		MapKills13Combo,
	}
	RelevantPropsKills = []PlayerProp{
		MapKills1,
		MapKills2,
		MapKills3,
		MapKills12,
		MapKills13,
	}
)
