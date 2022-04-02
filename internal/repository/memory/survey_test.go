package memory

import (
	"testing"

	"github.com/msoovali/personality-trait-survey/internal/domain"
)

const (
	surveyId = "survey-123"
)

func TestSurveyRepository_Save(t *testing.T) {
	r := NewSurveyRepository()

	type testCase struct {
		description       string
		input             domain.Survey
		expectedStoreSize int
		expectedId        string
	}

	for _, scenario := range []testCase{
		{
			description:       "surveyWithoutId_obtainsIdAndIncreasesStorage",
			input:             getSurveyMock(),
			expectedStoreSize: 1,
		},
		{
			description:       "surveyWithId_increasesStorage",
			input:             getSurveyMockWithId(),
			expectedStoreSize: 2,
			expectedId:        surveyId,
		},
		{
			description:       "existingSurveyWithId_overwritesEntity",
			input:             getSurveyMockWithId(),
			expectedStoreSize: 2,
			expectedId:        surveyId,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			q := r.Save(scenario.input)

			if q.Id == "" {
				t.Errorf("Expected survey to obtain id, but id is empty")
			}
			if scenario.expectedId != "" && q.Id != scenario.expectedId {
				t.Errorf("Expected survey with id %s, got %s", scenario.expectedId, q.Id)
			}
			if len(r.store) != scenario.expectedStoreSize {
				t.Errorf("Expected store size %d, got %d", scenario.expectedStoreSize, len(r.store))
			}
		})
	}
}

func getSurveyMockWithId() domain.Survey {
	s := getSurveyMock()
	s.Id = surveyId

	return s
}

func getSurveyMock() domain.Survey {
	return domain.Survey{}
}
