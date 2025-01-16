package underdog

type Appearance struct {
	// Badges         []string `json:"badges,omitempty"`
	Id             string `json:"id"`
	LineupStatusId string `json:"lineup_status_id"`
	MatchId        int    `json:"match_id"`
	MatchType      string `json:"match_type"`
	PlayerId       string `json:"player_id"`
	PositionId     string `json:"position_id"`
	TeamId         string `json:"team_id"`
}

type AppearanceStat struct {
	AppearanceId string `json:"appearance_id"`
	DisplayStat  string `json:"display_stat"`
	GradedBy     string `json:"graded_by"`
	Id           string `json:"id"`
	PickemStatId string `json:"pickem_stat_id"`
	Stat         string `json:"stat"`
}
