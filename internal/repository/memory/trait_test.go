package memory

import (
	"testing"

	"github.com/msoovali/personality-trait-survey/internal/domain"
)

const (
	traitId = "trait-123"
)

func TestTraitRepository_Save(t *testing.T) {
	r := NewTraitRepository()

	type testCase struct {
		description       string
		input             domain.Trait
		expectedStoreSize int
		expectedId        string
	}

	for _, scenario := range []testCase{
		{
			description:       "surveyWithoutId_obtainsIdAndIncreasesStorage",
			input:             getTraitMock(),
			expectedStoreSize: 1,
		},
		{
			description:       "surveyWithId_increasesStorage",
			input:             getTraitMockWithId(),
			expectedStoreSize: 2,
			expectedId:        traitId,
		},
		{
			description:       "existingSurveyWithId_overwritesEntity",
			input:             getTraitMockWithId(),
			expectedStoreSize: 2,
			expectedId:        traitId,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			q := r.Save(scenario.input)

			if q.Id == "" {
				t.Errorf("Expected trait to obtain id, but id is empty")
			}
			if scenario.expectedId != "" && q.Id != scenario.expectedId {
				t.Errorf("Expected trait with id %s, got %s", scenario.expectedId, q.Id)
			}
			if len(r.store) != scenario.expectedStoreSize {
				t.Errorf("Expected store size %d, got %d", scenario.expectedStoreSize, len(r.store))
			}
		})
	}
}

func getTraitMockWithId() domain.Trait {
	t := getTraitMock()
	t.Id = traitId

	return t
}

func getTraitMock() domain.Trait {
	return domain.Trait{}
}
