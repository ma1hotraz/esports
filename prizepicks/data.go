package prizepicks

type Data struct {
	Data     []Entity `json:"data"`
	Included []Entity `json:"included"`
	Links    Links    `json:"links"`
	Meta     Meta     `json:"meta"`
}

func (d *Data) Filter() []RelevantData {
	data := []RelevantData{}
	for _, e := range d.Data {
		if e.IsRelevant() {
			p := d.getPlayer(e.Relationships.NewPlayer.Data.Id)
			league := e.GetLeague()

			prop := e.GetProp(league.PlayerProps)
			if p.Type == NEW_PLAYER {
				data = append(data, RelevantData{
					ProjectionId: e.Id,
					Player:       p.Attributes.Name,
					Time:         e.Attributes.StartTime,
					Sport:        league.Type,
					StatType:     prop.Type,
					Value:        e.Attributes.LineScore,
				})
			}
		}
	}
	return data
}

func (d *Data) getPlayer(id string) Entity {
	var player Entity
	for _, p := range d.Included {
		if p.Type == NEW_PLAYER && p.Id == id {
			return p
		}
	}
	return player
}
