package service

import (
	"errors"

	"github.com/msoovali/personality-trait-survey/internal/domain"
	_errors "github.com/msoovali/personality-trait-survey/internal/errors"
	"github.com/msoovali/personality-trait-survey/internal/repository"
)

var (
	TraitIdEmptyError          = errors.New("TRAIT_ID_EMPTY")
	TraitTypeEmptyError        = errors.New("TRAIT_TYPE_EMPTY")
	TraitNotFoundError         = errors.New("TRAIT_NOT_FOUND")
	TraitMinScoreNegativeError = errors.New("TRAIT_MIN_SCORE_MUST_BE_POSITIVE")
)

type traitService struct {
	repository repository.TraitRepository
}

func NewTraitService(repository repository.TraitRepository) *traitService {
	return &traitService{
		repository: repository,
	}
}

func (s *traitService) GetAll() []domain.Trait {
	return s.repository.FindAll()
}

func (s *traitService) Add(trait domain.Trait) (domain.Trait, error) {
	if err := traitValidator(trait); err != nil {
		return trait, err
	}
	trait.Id = ""
	trait = s.repository.Save(trait)

	return trait, nil
}

func (s *traitService) Edit(trait domain.Trait) (domain.Trait, error) {
	if err := traitEditValidator(trait); err != nil {
		return trait, err
	}
	_, err := s.Get(trait.Id)
	if err != nil {
		return trait, err
	}
	trait = s.repository.Save(trait)

	return trait, nil
}

func (s *traitService) Get(id string) (domain.Trait, error) {
	q, err := s.repository.Find(id)
	if errors.Is(err, _errors.NoRecordFound) {
		return q, TraitNotFoundError
	}

	return q, err
}

func (s *traitService) Delete(id string) {
	s.repository.Delete(id)
}

var traitEditValidator = validateTraitEdit

func validateTraitEdit(trait domain.Trait) error {
	if trait.Id == "" {
		return TraitIdEmptyError
	}

	return traitValidator(trait)
}

var traitValidator = validateTrait

func validateTrait(trait domain.Trait) error {
	if trait.Type == "" {
		return TraitTypeEmptyError
	}
	if trait.MinScoreRequirement < 0 {
		return TraitMinScoreNegativeError
	}

	return nil
}
