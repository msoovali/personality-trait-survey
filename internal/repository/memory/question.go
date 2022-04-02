package memory

import (
	"github.com/google/uuid"
	"github.com/msoovali/personality-trait-survey/internal/domain"
)

type questionRepository struct {
	repository[domain.Question]
}

func NewQuestionRepository() *questionRepository {
	return &questionRepository{
		repository: *newRepository[domain.Question](),
	}
}

func (r *questionRepository) Save(question domain.Question) domain.Question {
	if question.Id == "" {
		question.Id = uuid.New().String()
	}
	options := make([]domain.Option, 0, len(question.Options))
	for _, o := range question.Options {
		if o.Id == "" {
			o.Id = uuid.New().String()
		}
		options = append(options, o)
	}
	question.Options = options
	r.save(question)

	return question
}
