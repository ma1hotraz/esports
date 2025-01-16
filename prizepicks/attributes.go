package prizepicks

import "time"

type Attributes struct {
	Active               bool      `json:"active"`
	AdjustedOdds         bool      `json:"adjusted_odds"`
	BoardTime            time.Time `json:"board_time"`
	Combo                bool      `json:"combo"`
	CustomImage          string    `json:"custom_image"`
	Description          string    `json:"description"`
	DisplayName          string    `json:"display_name"`
	EndTime              time.Time `json:"end_time"`
	FlashSaleLineScore   float64   `json:"flash_sale_line_score"`
	F2pEnabled           bool      `json:"f2p_enabled"`
	Hr20                 bool      `json:"hr_20"`
	Icon                 string    `json:"icon"`
	ImageUrl             string    `json:"image_url"`
	IsPromo              bool      `json:"is_promo"`
	LastFiveGamesEnabled bool      `json:"last_five_games_enabled"`
	League               string    `json:"league"`
	LeagueIconId         int       `json:"league_icon_id"`
	LeagueId             int       `json:"league_id"`
	LfgIgnoredLeagues    []int     `json:"lfg_ignored_leagues"`
	LineScore            float64   `json:"line_score"`
	Market               string    `json:"market"`
	Name                 string    `json:"name"`
	OddsType             string    `json:"odds_type"`
	ProjectionType       string    `json:"projection_type"`
	ProjectionsCount     int       `json:"projections_count"`
	Rank                 int       `json:"rank"`
	Refundable           bool      `json:"refundable"`
	ShowTrending         bool      `json:"show_trending"`
	StartTime            time.Time `json:"start_time"`
	StatType             string    `json:"stat_type"`
	Status               string    `json:"status"`
	Team                 string    `json:"team"`
	TeamName             string    `json:"team_name"`
	Today                bool      `json:"today"`
	TvChannel            string    `json:"tv_channel"`
	UpdatedAt            time.Time `json:"updated_at"`
}
