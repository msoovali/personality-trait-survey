package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/msoovali/personality-trait-survey/internal/domain"
	"github.com/msoovali/personality-trait-survey/internal/service"
)

type surveyHandlers struct {
	service service.SurveyService
}

func NewSurveyHandlers(service service.SurveyService) *surveyHandlers {
	return &surveyHandlers{
		service: service,
	}
}

func (h *surveyHandlers) GetResults(c *fiber.Ctx) error {
	s, err := h.service.GetResults(c.Params("surveyId"))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(s)
}

func (h *surveyHandlers) FinishSurvey(c *fiber.Ctx) error {
	r := new(domain.Survey)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	s, err := h.service.Finish(createImmutableSurvey(*r))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(s)
}

func createImmutableSurvey(survey domain.Survey) domain.Survey {
	return domain.Survey{
		Answers: createImmutableAnswers(survey.Answers),
	}
}

func createImmutableAnswers(answers []domain.Answer) []domain.Answer {
	immutableAnswers := make([]domain.Answer, 0, len(answers))
	for _, a := range answers {
		immutableAnswers = append(immutableAnswers, domain.Answer{
			QuestionId: utils.CopyString(a.QuestionId),
			OptionId:   utils.CopyString(a.OptionId),
		})
	}

	return immutableAnswers
}
