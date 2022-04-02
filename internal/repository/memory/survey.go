package memory

import (
	"github.com/google/uuid"
	"github.com/msoovali/personality-trait-survey/internal/domain"
)

type surveyRepository struct {
	repository[domain.Survey]
}

func NewSurveyRepository() *surveyRepository {
	return &surveyRepository{
		repository: *newRepository[domain.Survey](),
	}
}

func (r *surveyRepository) Save(survey domain.Survey) domain.Survey {
	if survey.Id == "" {
		survey.Id = uuid.New().String()
	}
	r.save(survey)

	return survey
}
