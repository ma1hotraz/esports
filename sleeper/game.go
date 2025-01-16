package sleeper

type GameInDay []GameData

type GameData struct {
	Status         string       `json:"status"`
	Date           string       `json:"date"`
	Metadata       GameMetadata `json:"metadata"`
	StartTime      int64        `json:"start_time"`
	Week           int64        `json:"week"`
	Season         string       `json:"season"`
	SeasonType     string       `json:"season_type"`
	Sport          string       `json:"sport"`
	GameID         string       `json:"game_id"`
	UpdatedAt      int64        `json:"updated_at"`
	Reactions      interface{}  `json:"reactions"`
	LastUpdated    int64        `json:"last_updated"`
	ProviderID     string       `json:"provider_id"`
	ScheduleHash   interface{}  `json:"schedule_hash"`
	TotalReactions interface{}  `json:"total_reactions"`
	TotalViews     interface{}  `json:"total_views"`
	TotalComments  interface{}  `json:"total_comments"`
	ProviderWeek   interface{}  `json:"provider_week"`
}

type GameMetadata struct {
	AwayTeam         DataTeam `json:"away_team"`
	EventName        string   `json:"event_name"`
	HomeTeam         DataTeam `json:"home_team"`
	OriginalTimeUTCS int64    `json:"original_time_utc_s"`
	StartTimeUTCS    int64    `json:"start_time_utc_s"`
}

type DataTeam struct {
	Name string `json:"name"`
	Team string `json:"team"`
}
