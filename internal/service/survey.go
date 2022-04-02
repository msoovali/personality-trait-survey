package service

import (
	"errors"

	"github.com/msoovali/personality-trait-survey/internal/domain"
	_errors "github.com/msoovali/personality-trait-survey/internal/errors"
	"github.com/msoovali/personality-trait-survey/internal/repository"
)

var (
	SurveyQuestionIdEmpty    = errors.New("SURVEY_QUESTION_ID_EMPTY")
	SurveyOptionIdEmpty      = errors.New("SURVEY_OPTION_ID_EMPTY")
	SurveyQuestionIdNotFound = errors.New("SURVEY_QUESTION_ID_NOT_FOUND")
	SurveyOptionIdNotFound   = errors.New("SURVEY_OPTION_ID_NOT_FOUND")
	SurveyMissingAnswers     = errors.New("SURVEY_MISSING_ANSWERS")
	SurveyNotFound           = errors.New("SURVEY_NOT_FOUND")
	SurveyResultExpired      = errors.New("SURVEY_RESULT_EXPIRED")
)

type surveyService struct {
	repository      repository.SurveyRepository
	questionService QuestionService
	traitService    TraitService
}

func NewSurveyService(repository repository.SurveyRepository, questionService QuestionService, traitService TraitService) *surveyService {
	return &surveyService{
		repository:      repository,
		questionService: questionService,
		traitService:    traitService,
	}
}

func (s *surveyService) GetResults(id string) (domain.SurveyResponse, error) {
	survey, err := s.repository.Find(id)
	if errors.Is(err, _errors.NoRecordFound) {
		return domain.SurveyResponse{}, SurveyNotFound
	}
	trait, err := s.traitService.Get(survey.TraitId)
	if errors.Is(err, _errors.NoRecordFound) {
		return domain.SurveyResponse{}, SurveyResultExpired
	}

	return domain.SurveyResponse{
		Id:    survey.Id,
		Score: survey.Score,
		Trait: trait,
	}, nil
}

func (s *surveyService) Finish(survey domain.Survey) (domain.SurveyResponse, error) {
	// calculate score
	questions := s.questionService.GetAll()
	score := 0
	for _, q := range questions {
		surveyIncludesAnswer := false
		for _, a := range survey.Answers {
			if a.QuestionId == q.Id {
				surveyIncludesAnswer = true
				option, err := findOptionById(a.OptionId, q.Options)
				if err != nil {
					return domain.SurveyResponse{}, err
				}
				score += option.Score
				break
			}
		}
		if !surveyIncludesAnswer {
			return domain.SurveyResponse{}, SurveyMissingAnswers
		}
	}
	survey.Score = score
	// determine trait
	traits := s.traitService.GetAll()
	var surveyTrait domain.Trait
	for _, t := range traits {
		if surveyTrait.MinScoreRequirement <= t.MinScoreRequirement && score >= t.MinScoreRequirement {
			surveyTrait = t
		}
	}
	survey.TraitId = surveyTrait.Id
	// save and return survey
	survey = s.repository.Save(survey)
	return domain.SurveyResponse{
		Id:    survey.Id,
		Score: survey.Score,
		Trait: surveyTrait,
	}, nil
}

func findOptionById(id string, options []domain.Option) (domain.Option, error) {
	for _, o := range options {
		if o.Id == id {
			return o, nil
		}
	}
	return domain.Option{}, SurveyOptionIdNotFound
}
