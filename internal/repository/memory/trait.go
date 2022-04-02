package memory

import (
	"github.com/google/uuid"
	"github.com/msoovali/personality-trait-survey/internal/domain"
)

type traitRepository struct {
	repository[domain.Trait]
}

func NewTraitRepository() *traitRepository {
	return &traitRepository{
		repository: *newRepository[domain.Trait](),
	}
}

func (r *traitRepository) Save(trait domain.Trait) domain.Trait {
	if trait.Id == "" {
		trait.Id = uuid.New().String()
	}
	r.save(trait)

	return trait
}
