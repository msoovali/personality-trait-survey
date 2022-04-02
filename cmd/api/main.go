package main

import (
	"github.com/msoovali/personality-trait-survey/internal/app"
)

func main() {
	a := app.NewAPI()
	a.StartAPI()
}
