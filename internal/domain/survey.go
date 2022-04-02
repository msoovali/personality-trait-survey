package domain

type Survey struct {
	Id      string   `json:"id"`
	Score   int      `json:"score"`
	TraitId string   `json:"traitId"`
	Answers []Answer `json:"answers"`
}

func (s Survey) GetId() string {
	return s.Id
}

type Answer struct {
	QuestionId string `json:"questionId"`
	OptionId   string `json:"optionId"`
}

type SurveyResponse struct {
	Id    string `json:"id"`
	Score int    `json:"score"`
	Trait Trait  `json:"trait"`
}
