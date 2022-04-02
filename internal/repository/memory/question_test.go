package memory

import (
	"testing"

	"github.com/msoovali/personality-trait-survey/internal/domain"
)

const (
	questionId = "test-123"
)

func TestQuestionRepository_Save(t *testing.T) {
	r := NewQuestionRepository()

	type testCase struct {
		description        string
		input              domain.Question
		expectedStoreSize  int
		expectedQuestionId string
	}

	for _, scenario := range []testCase{
		{
			description:       "questionWithoutId_obtainsIdAndIncreasesStorage",
			input:             getQuestionMock(),
			expectedStoreSize: 1,
		},
		{
			description:        "questionWithId_increasesStorage",
			input:              getQuestionMockWithId(),
			expectedStoreSize:  2,
			expectedQuestionId: questionId,
		},
		{
			description:        "existingQuestionWithId_overwritesEntity",
			input:              getQuestionMockWithId(),
			expectedStoreSize:  2,
			expectedQuestionId: questionId,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			q := r.Save(scenario.input)

			if q.Id == "" {
				t.Errorf("Expected question to obtain id, but id is empty")
			}
			if scenario.expectedQuestionId != "" && q.Id != scenario.expectedQuestionId {
				t.Errorf("Expected question with id %s, got %s", scenario.expectedQuestionId, q.Id)
			}
			if len(r.store) != scenario.expectedStoreSize {
				t.Errorf("Expected store size %d, got %d", scenario.expectedStoreSize, len(r.store))
			}
		})
	}
}

func getQuestionMockWithId() domain.Question {
	q := getQuestionMock()
	q.Id = questionId

	return q
}

func getQuestionMock() domain.Question {
	return domain.Question{}
}
