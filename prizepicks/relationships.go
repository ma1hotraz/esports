package prizepicks

type Relationships struct {
	Duration       Relationship `json:"duration"`
	League         Relationship `json:"league"`
	NewPlayer      Relationship `json:"new_player"`
	ProjectionType Relationship `json:"projection_type"`
	StatType       Relationship `json:"stat_type"`
}

type Relationship struct {
	Data RelationshipData `json:"data"`
}

type RelationshipData struct {
	Id   string     `json:"id"`
	Type EntityType `json:"type"`
}
