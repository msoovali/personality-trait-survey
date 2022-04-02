package service

import (
	"errors"

	"github.com/msoovali/personality-trait-survey/internal/domain"
	_errors "github.com/msoovali/personality-trait-survey/internal/errors"
	"github.com/msoovali/personality-trait-survey/internal/repository"
)

var (
	QuestionIdEmptyError           = errors.New("QUESTION_ID_EMPTY")
	QuestionDescriptionEmptyError  = errors.New("QUESTION_DESCRIPTION_EMPTY")
	AtleastTwoOptionsRequiredError = errors.New("ATLEAST_TWO_OPTIONS_REQUIRED")
	OptionDescriptionEmptyError    = errors.New("OPTION_DESCRIPTION_EMPTY")
	OptionScoreNegativeError       = errors.New("OPTION_SCORE_MUST_BE_POSITIVE")
	QuestionNotFoundError          = errors.New("QUESTION_NOT_FOUND")
)

type questionService struct {
	repository repository.QuestionRepository
}

func NewQuestionService(repository repository.QuestionRepository) *questionService {
	return &questionService{
		repository: repository,
	}
}

func (s *questionService) GetAll() []domain.Question {
	return s.repository.FindAll()
}

func (s *questionService) Add(question domain.Question) (domain.Question, error) {
	if err := questionValidator(question); err != nil {
		return question, err
	}
	question.Id = ""
	for _, o := range question.Options {
		o.Id = ""
	}
	question = s.repository.Save(question)

	return question, nil
}

func (s *questionService) Edit(question domain.Question) (domain.Question, error) {
	if err := questionEditValidator(question); err != nil {
		return question, err
	}
	existingQuestion, err := s.Get(question.Id)
	if err != nil {
		return question, err
	}
	for _, o := range question.Options {
		optionExists := false
		for _, existingOption := range existingQuestion.Options {
			if o.Id == existingOption.Id {
				optionExists = true
			}
		}
		if !optionExists {
			o.Id = ""
		}
	}
	question = s.repository.Save(question)

	return question, nil
}

func (s *questionService) Get(id string) (domain.Question, error) {
	q, err := s.repository.Find(id)
	if errors.Is(err, _errors.NoRecordFound) {
		return q, QuestionNotFoundError
	}

	return q, err
}

func (s *questionService) Delete(id string) {
	s.repository.Delete(id)
}

var questionEditValidator = validateQuestionEdit

func validateQuestionEdit(question domain.Question) error {
	if question.Id == "" {
		return QuestionIdEmptyError
	}

	return questionValidator(question)
}

var questionValidator = validateQuestion

func validateQuestion(question domain.Question) error {
	if question.Description == "" {
		return QuestionDescriptionEmptyError
	}
	if len(question.Options) < 2 {
		return AtleastTwoOptionsRequiredError
	}
	for _, o := range question.Options {
		if o.Description == "" {
			return OptionDescriptionEmptyError
		}
		if o.Score < 0 {
			return OptionScoreNegativeError
		}
	}

	return nil
}
