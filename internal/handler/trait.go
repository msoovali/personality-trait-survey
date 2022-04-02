package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/msoovali/personality-trait-survey/internal/domain"
	"github.com/msoovali/personality-trait-survey/internal/service"
)

type traitHandlers struct {
	service service.TraitService
}

func NewTraitHandlers(service service.TraitService) *traitHandlers {
	return &traitHandlers{
		service: service,
	}
}

func (h *traitHandlers) GetAll(c *fiber.Ctx) error {
	traits := h.service.GetAll()

	return c.JSON(traits)
}

func (h *traitHandlers) Get(c *fiber.Ctx) error {
	q, err := h.service.Get(c.Params("traitId"))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(q)
}

func (h *traitHandlers) Add(c *fiber.Ctx) error {
	r := new(domain.Trait)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	q, err := h.service.Add(createImmutableTrait(*r))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(q)
}

func (h *traitHandlers) Edit(c *fiber.Ctx) error {
	r := new(domain.Trait)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	q, err := h.service.Edit(createImmutableTrait(*r))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}

	return c.JSON(q)
}

func (h *traitHandlers) Delete(c *fiber.Ctx) error {
	h.service.Delete(c.Params("traitId"))

	return c.SendStatus(fiber.StatusNoContent)
}

func createImmutableTrait(t domain.Trait) domain.Trait {
	return domain.Trait{
		Id:                  utils.CopyString(t.Id),
		Type:                utils.CopyString(t.Type),
		MinScoreRequirement: t.MinScoreRequirement,
	}
}
