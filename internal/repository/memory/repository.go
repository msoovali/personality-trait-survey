package memory

import (
	"github.com/msoovali/personality-trait-survey/internal/domain"
	"github.com/msoovali/personality-trait-survey/internal/errors"
)

type repository[T domain.IEntity] struct {
	store map[string]T
}

func newRepository[T domain.IEntity]() *repository[T] {
	return &repository[T]{
		store: make(map[string]T),
	}
}

func (r *repository[T]) Find(id string) (T, error) {
	entity, exists := r.store[id]
	if exists {
		return entity, nil
	}

	return entity, errors.NoRecordFound
}

func (r *repository[T]) FindAll() []T {
	entities := make([]T, 0, len(r.store))
	for _, q := range r.store {
		entities = append(entities, q)
	}

	return entities
}

func (r *repository[T]) save(entity T) {
	r.store[entity.GetId()] = entity
}

func (r *repository[T]) Delete(id string) {
	delete(r.store, id)
}
