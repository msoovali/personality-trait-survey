package domain

type Trait struct {
	Id                  string `json:"id"`
	Type                string `json:"type"`
	MinScoreRequirement int    `json:"minScoreRequirement"`
}

func (t Trait) GetId() string {
	return t.Id
}
