package service

import "github.com/msoovali/personality-trait-survey/internal/domain"

type QuestionService interface {
	GetAll() []domain.Question
	Add(question domain.Question) (domain.Question, error)
	Edit(question domain.Question) (domain.Question, error)
	Get(id string) (domain.Question, error)
	Delete(id string)
}

type TraitService interface {
	GetAll() []domain.Trait
	Add(trait domain.Trait) (domain.Trait, error)
	Edit(trait domain.Trait) (domain.Trait, error)
	Get(id string) (domain.Trait, error)
	Delete(id string)
}

type SurveyService interface {
	GetResults(id string) (domain.SurveyResponse, error)
	Finish(survey domain.Survey) (domain.SurveyResponse, error)
}
