package repository

import "github.com/msoovali/personality-trait-survey/internal/domain"

type repository[T domain.IEntity] interface {
	Find(id string) (T, error)
	FindAll() []T
	Delete(id string)
}

type QuestionRepository interface {
	repository[domain.Question]
	Save(question domain.Question) domain.Question
}

type SurveyRepository interface {
	repository[domain.Survey]
	Save(survey domain.Survey) domain.Survey
}

type TraitRepository interface {
	repository[domain.Trait]
	Save(trait domain.Trait) domain.Trait
}
