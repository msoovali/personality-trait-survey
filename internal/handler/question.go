package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/msoovali/personality-trait-survey/internal/domain"
	"github.com/msoovali/personality-trait-survey/internal/service"
)

type questionHandlers struct {
	service service.QuestionService
}

func NewQuestionHandlers(service service.QuestionService) *questionHandlers {
	return &questionHandlers{
		service: service,
	}
}

func (h *questionHandlers) GetAll(c *fiber.Ctx) error {
	questions := h.service.GetAll()

	return c.JSON(questions)
}

func (h *questionHandlers) Get(c *fiber.Ctx) error {
	q, err := h.service.Get(c.Params("questionId"))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(q)
}

func (h *questionHandlers) Add(c *fiber.Ctx) error {
	r := new(domain.Question)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	q, err := h.service.Add(createImmutableQuestion(*r))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(q)
}

func (h *questionHandlers) Edit(c *fiber.Ctx) error {
	r := new(domain.Question)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	q, err := h.service.Edit(createImmutableQuestion(*r))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(q)
}

func (h *questionHandlers) Delete(c *fiber.Ctx) error {
	h.service.Delete(c.Params("questionId"))

	return c.SendStatus(fiber.StatusNoContent)
}

func createImmutableQuestion(q domain.Question) domain.Question {
	return domain.Question{
		Id:          utils.CopyString(q.Id),
		Description: utils.CopyString(q.Description),
		Options:     createImmutableOptions(q.Options),
	}
}

func createImmutableOptions(options []domain.Option) []domain.Option {
	immutableOptions := make([]domain.Option, 0, len(options))
	for _, o := range options {
		immutableOptions = append(immutableOptions, domain.Option{
			Id:          utils.CopyString(o.Id),
			Description: utils.CopyString(o.Description),
			Score:       o.Score,
		})
	}

	return immutableOptions
}
