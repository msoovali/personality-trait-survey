package handler

import "github.com/gofiber/fiber/v2"

type HealthHandlers interface {
	HealthCheck(c *fiber.Ctx) error
}

type QuestionHandlers interface {
	GetAll(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type TraitHandlers interface {
	GetAll(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type SurveyHandlers interface {
	GetResults(c *fiber.Ctx) error
	FinishSurvey(c *fiber.Ctx) error
}
