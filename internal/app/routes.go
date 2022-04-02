package app

func (a *application) registerRoutes() {
	a.router.Get("/health", a.handlers.health.HealthCheck)

	v1 := a.router.Group("/v1")
	{
		v1.Get("/questions", a.handlers.question.GetAll)
		v1.Get("/question/:questionId", a.handlers.question.Get)
		v1.Post("/question", a.handlers.question.Add)
		v1.Put("/question", a.handlers.question.Edit)
		v1.Delete("/question/:questionId", a.handlers.question.Delete)

		v1.Get("/traits", a.handlers.trait.GetAll)
		v1.Get("/trait/:traitId", a.handlers.trait.Get)
		v1.Post("/trait", a.handlers.trait.Add)
		v1.Put("/trait", a.handlers.trait.Edit)
		v1.Delete("/trait/:traitId", a.handlers.trait.Delete)

		v1.Get("/survey/result/:surveyId", a.handlers.survey.GetResults)
		v1.Post("/survey/finish", a.handlers.survey.FinishSurvey)
	}
}
