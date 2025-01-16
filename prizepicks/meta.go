package prizepicks

type Links struct {
	Next string `json:"next"`
	Self string `json:"self"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}
