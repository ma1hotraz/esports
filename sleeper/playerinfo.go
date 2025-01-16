package sleeper

type PlayerData map[string]PlayerInfo

type PlayerInfo struct {
	Active           bool         `json:"active"`
	Position         interface{}  `json:"position"`
	Number           interface{}  `json:"number"`
	Metadata         UserMetadata `json:"metadata"`
	Username         string       `json:"username"`
	Age              interface{}  `json:"age"`
	Sport            string       `json:"sport"`
	PlayerID         string       `json:"player_id"`
	FantasyPositions interface{}  `json:"fantasy_positions"`
	LastName         string       `json:"last_name"`
	FirstName        string       `json:"first_name"`
	Team             string       `json:"team"`
	TeamAbbr         interface{}  `json:"team_abbr"`
	TeamChangedAt    interface{}  `json:"team_changed_at"`
	Hashtag          string       `json:"hashtag"`
	BirthDate        interface{}  `json:"birth_date"`
	YearsExp         interface{}  `json:"years_exp"`
}

type UserMetadata struct {
	Username string `json:"username"`
}
