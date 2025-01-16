package underdog

type Player struct {
	Country    string `json:"country"`
	FirstName  string `json:"first_name"`
	Id         string `json:"id"`
	ImageUrl   string `json:"image_url"`
	LastName   string `json:"last_name"`
	PositionId string `json:"position_id"`
	SportId    string `json:"sport_id"`
	TeamId     string `json:"team_id"`
}
