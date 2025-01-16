package esport

import "strings"

type PlayerPropType string

const (
	// normal
	MAP_KILLS_1       PlayerPropType = "MAPS 1 Kills"
	MAP_KILLS_2       PlayerPropType = "MAPS 2 Kills"
	MAP_KILLS_3       PlayerPropType = "MAPS 3 Kills"
	MAP_KILLS_1_2     PlayerPropType = "MAPS 1-2 Kills"
	MAP_KILLS_1_3     PlayerPropType = "MAPS 1-3 Kills"
	MAP_ASSISTS_1     PlayerPropType = "MAPS 1 Assists"
	MAP_ASSISTS_2     PlayerPropType = "MAPS 2 Assists"
	MAP_ASSISTS_3     PlayerPropType = "MAPS 3 Assists"
	MAP_ASSISTS_1_2   PlayerPropType = "MAPS 1-2 Assists"
	MAP_ASSISTS_1_3   PlayerPropType = "MAPS 1-3 Assists"
	MAP_HEADSHOTS_1   PlayerPropType = "MAPS 1 Headshots"
	MAP_HEADSHOTS_2   PlayerPropType = "MAPS 2 Headshots"
	MAP_HEADSHOTS_3   PlayerPropType = "MAPS 3 Headshots"
	MAP_HEADSHOTS_1_2 PlayerPropType = "MAPS 1-2 Headshots"
	MAP_HEADSHOTS_1_3 PlayerPropType = "MAPS 1-3 Headshots"

	// combo
	MAP_KILLS_1_COMBO   PlayerPropType = "MAPS 1 Kills Combo"
	MAP_KILLS_2_COMBO   PlayerPropType = "MAPS 2 Kills Combo"
	MAP_KILLS_3_COMBO   PlayerPropType = "MAPS 3 Kills Combo"
	MAP_KILLS_1_2_COMBO PlayerPropType = "MAPS 1-2 Kills Combo"
	MAP_KILLS_1_3_COMBO PlayerPropType = "MAPS 1-3 Kills Combo"
	// MAP_ASSISTS_1_COMBO PlayerPropType = "MAPS 1 Assists Combo"
	// MAP_ASSISTS_2_COMBO PlayerPropType = "MAPS 2 Assists Combo"
	// MAP_ASSISTS_3_COMBO PlayerPropType = "MAPS 3 Assists Combo"
	// MAP_ASSISTS_1_2_COMBO PlayerPropType = "MAPS 1-2 Assists Combo"
	// MAP_ASSISTS_1_3_COMBO PlayerPropType = "MAPS 1-3 Assists Combo"
	// MAP_HEADSHOTS_1_COMBO PlayerPropType = "MAPS 1 Headshots Combo"
	// MAP_HEADSHOTS_2_COMBO PlayerPropType = "MAPS 2 Headshots Combo"
	// MAP_HEADSHOTS_3_COMBO PlayerPropType = "MAPS 3 Headshots Combo"
	// MAP_HEADSHOTS_1_2_COMBO PlayerPropType = "MAPS 1-2 Headshots Combo"
	// MAP_HEADSHOTS_1_3_COMBO PlayerPropType = "MAPS 1-3 Headshots Combo"
)

type PlayerProp struct {
	Type  PlayerPropType
	Names []string
}

func (p PlayerProp) RecognizeByName(name string) bool {
	for _, n := range p.Names {
		if strings.ToLower(name) == n {
			return true
		}
	}
	return false
}

var COMBO_TYPES_LIST = []PlayerPropType{
	MAP_KILLS_1_COMBO,
	MAP_KILLS_2_COMBO,
	MAP_KILLS_3_COMBO,
	MAP_KILLS_1_2_COMBO,
	MAP_KILLS_1_3_COMBO,
}

func IsComboType(t PlayerPropType) bool {
	for _, ct := range COMBO_TYPES_LIST {
		if ct == t {
			return true
		}
	}
	return false
}
