package memory

import (
	"errors"
	"testing"

	"github.com/msoovali/personality-trait-survey/internal/domain"
	_errors "github.com/msoovali/personality-trait-survey/internal/errors"
)

func TestRepository_save(t *testing.T) {
	r := newRepository[domain.Question]()

	type testCase struct {
		description        string
		input              domain.Question
		expectedStoreSize  int
		expectedQuestionId string
	}

	for _, scenario := range []testCase{
		// {
		// 	description:       "questionWithoutId_obtainsIdAndIncreasesStorage",
		// 	input:             getQuestionMock(),
		// 	expectedStoreSize: 1,
		// },
		{
			description:        "questionWithId_increasesStorage",
			input:              getQuestionMockWithId(),
			expectedStoreSize:  1,
			expectedQuestionId: questionId,
		},
		{
			description:        "existingQuestionWithId_overwritesEntity",
			input:              getQuestionMockWithId(),
			expectedStoreSize:  1,
			expectedQuestionId: questionId,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			r.save(scenario.input)

			// if q.Id == "" {
			// 	t.Errorf("Expected question to obtain id, but id is empty")
			// }
			// if scenario.expectedQuestionId != "" && q.Id != scenario.expectedQuestionId {
			// 	t.Errorf("Expected question with id %s, got %s", scenario.expectedQuestionId, q.Id)
			// }
			if len(r.store) != scenario.expectedStoreSize {
				t.Errorf("Expected store size %d, got %d", scenario.expectedStoreSize, len(r.store))
			}
		})
	}
}

func TestRepository_Find(t *testing.T) {
	r := newRepository[domain.Question]()
	r.store[questionId] = getQuestionMockWithId()

	type testCase struct {
		description        string
		input              string
		expectedError      error
		expectedQuestionId string
	}

	for _, scenario := range []testCase{
		{
			description:        "questionExists_returnsEntityWithId",
			input:              questionId,
			expectedQuestionId: questionId,
		},
		{
			description:   "questionNotExist_returnsNil",
			input:         "hello",
			expectedError: _errors.NoRecordFound,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			q, err := r.Find(scenario.input)

			if !errors.Is(err, scenario.expectedError) {
				t.Errorf("Expected error %v, but got %v", scenario.expectedError, err)
			}
			if err == nil && scenario.expectedQuestionId != q.Id {
				t.Errorf("Expected question with id %s, got %s", scenario.expectedQuestionId, q.Id)
			}
		})
	}
}

func TestRepository_FindAll(t *testing.T) {
	r := newRepository[domain.Question]()
	r.store[questionId] = getQuestionMockWithId()
	r.store["1222"] = getQuestionMockWithId()

	type testCase struct {
		description string
	}

	for _, scenario := range []testCase{
		{
			description: "callMethod_returnsStoreValues",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			questions := r.FindAll()

			if len(questions) != len(r.store) {
				t.Errorf("Expected to get %d questions, but got %d", len(r.store), len(questions))
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	r := newRepository[domain.Question]()
	r.store[questionId] = getQuestionMockWithId()

	type testCase struct {
		description       string
		input             string
		expectedStoreSize int
	}

	for _, scenario := range []testCase{
		{
			description:       "nonExistentId_nothingIsDeletedFromStore",
			input:             "hi",
			expectedStoreSize: 1,
		},
		{
			description:       "idThatExists_questionIsDeletedFromStore",
			input:             questionId,
			expectedStoreSize: 0,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			r.Delete(scenario.input)

			if scenario.expectedStoreSize != len(r.store) {
				t.Errorf("Expected store size %d, but got %d", scenario.expectedStoreSize, len(r.store))
			}
		})
	}
}
