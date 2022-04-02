package domain

type Option struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Score       int    `json:"score"`
}

type Question struct {
	Id          string   `json:"id"`
	Description string   `json:"description"`
	Options     []Option `json:"options"`
}

func (q Question) GetId() string {
	return q.Id
}
