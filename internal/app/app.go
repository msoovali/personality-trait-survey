package app

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/msoovali/personality-trait-survey/internal/conf"
	"github.com/msoovali/personality-trait-survey/internal/data"
	"github.com/msoovali/personality-trait-survey/internal/handler"
	"github.com/msoovali/personality-trait-survey/internal/repository"
	"github.com/msoovali/personality-trait-survey/internal/repository/memory"
	"github.com/msoovali/personality-trait-survey/internal/service"
)

type repositories struct {
	question repository.QuestionRepository
	trait    repository.TraitRepository
	survey   repository.SurveyRepository
}

type services struct {
	question service.QuestionService
	trait    service.TraitService
	survey   service.SurveyService
}

type handlers struct {
	health   handler.HealthHandlers
	question handler.QuestionHandlers
	trait    handler.TraitHandlers
	survey   handler.SurveyHandlers
}

type application struct {
	config       conf.ApplicationConfig
	repositories repositories
	services     services
	handlers     handlers
	router       *fiber.App
}

func NewAPI() *application {
	config := conf.ParseConfig()
	app := &application{
		config: config,
	}
	app.initRepositories()
	app.initServices()
	app.initHandlers()

	app.loadData()

	app.router = fiber.New(fiber.Config{
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	})
	app.router.Use(logger.New())
	app.registerRoutes()

	return app
}

func (a *application) StartAPI() {
	log.Fatal(a.router.Listen(a.config.Addr))
}

func (a *application) loadData() {
	data := data.LoadData()
	if data != nil {
		for _, q := range data.Questions {
			a.services.question.Add(q)
		}
		for _, t := range data.Traits {
			a.services.trait.Add(t)
		}
	}
}

func (a *application) initRepositories() {
	a.repositories = repositories{
		question: memory.NewQuestionRepository(),
		trait:    memory.NewTraitRepository(),
		survey:   memory.NewSurveyRepository(),
	}
}

func (a *application) initServices() {
	a.services = services{
		question: service.NewQuestionService(a.repositories.question),
		trait:    service.NewTraitService(a.repositories.trait),
	}
	a.services.survey = service.NewSurveyService(a.repositories.survey, a.services.question, a.services.trait)
}

func (a *application) initHandlers() {
	a.handlers = handlers{
		health:   handler.NewHealthHandlers(),
		question: handler.NewQuestionHandlers(a.services.question),
		trait:    handler.NewTraitHandlers(a.services.trait),
		survey:   handler.NewSurveyHandlers(a.services.survey),
	}
}
