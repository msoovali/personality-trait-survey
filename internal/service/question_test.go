package service

import (
	"errors"
	"testing"

	"github.com/msoovali/personality-trait-survey/internal/domain"
	_errors "github.com/msoovali/personality-trait-survey/internal/errors"
	"github.com/msoovali/personality-trait-survey/internal/repository"
)

const (
	questionId          = "question-123"
	questionDescription = "question-description"
	optionDescription   = "option-description"
)

type questionRepositoryMock struct {
	repository.QuestionRepository
	fakeFind    func(id string) (domain.Question, error)
	fakeFindAll func() []domain.Question
	fakeSave    func(question domain.Question) domain.Question
	fakeDelete  func(id string)
}

func (r *questionRepositoryMock) Find(id string) (domain.Question, error) {
	if r.fakeFind != nil {
		return r.fakeFind(id)
	}

	return domain.Question{}, nil
}

func (r *questionRepositoryMock) FindAll() []domain.Question {
	if r.fakeFindAll != nil {
		return r.fakeFindAll()
	}

	return make([]domain.Question, 0)
}

func (r *questionRepositoryMock) Save(question domain.Question) domain.Question {
	if r.fakeSave != nil {
		return r.fakeSave(question)
	}

	return question
}

func (r *questionRepositoryMock) Delete(id string) {
	if r.fakeDelete != nil {
		r.fakeDelete(id)
	}
}

func TestQuestionService_validateQuestion(t *testing.T) {
	type testCase struct {
		description string
		input       domain.Question
		expectedErr error
	}

	for _, scenario := range []testCase{
		{
			description: "descriptionEmpty_returnsError",
			input:       domain.Question{},
			expectedErr: QuestionDescriptionEmptyError,
		},
		{
			description: "zeroOptions_returnsError",
			input: domain.Question{
				Description: questionDescription,
			},
			expectedErr: AtleastTwoOptionsRequiredError,
		},
		{
			description: "oneOption_returnsError",
			input: domain.Question{
				Description: questionDescription,
				Options: []domain.Option{
					getOptionMock(),
				},
			},
			expectedErr: AtleastTwoOptionsRequiredError,
		},
		{
			description: "twoOptionsButOneWithoutDescription_returnsError",
			input: domain.Question{
				Description: questionDescription,
				Options: []domain.Option{
					getOptionMock(),
					{},
				},
			},
			expectedErr: OptionDescriptionEmptyError,
		},
		{
			description: "validQuestion_returnsNil",
			input: domain.Question{
				Description: questionDescription,
				Options: []domain.Option{
					getOptionMock(),
					getOptionMock(),
				},
			},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			err := validateQuestion(scenario.input)
			if !errors.Is(err, scenario.expectedErr) {
				t.Errorf("Expected %v, got %v", scenario.expectedErr, err)
			}
		})
	}
}

func TestQuestionService_validateQuestionEdit(t *testing.T) {
	type testCase struct {
		description                    string
		input                          domain.Question
		expectedErr                    error
		expectedQuestionValidatorCalls int
	}
	questionValidatorCalled := 0
	questionValidator = func(question domain.Question) error {
		questionValidatorCalled++
		return nil
	}
	for _, scenario := range []testCase{
		{
			description: "idEmpty_returnsError",
			input:       domain.Question{},
			expectedErr: QuestionIdEmptyError,
		},
		{
			description: "idFilled_returnsNil",
			input: domain.Question{
				Id: questionId,
			},
			expectedQuestionValidatorCalls: 1,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			questionValidatorCalled = 0
			err := validateQuestionEdit(scenario.input)

			if !errors.Is(err, scenario.expectedErr) {
				t.Errorf("Expected %v, got %v", scenario.expectedErr, err)
			}
			if questionValidatorCalled != scenario.expectedQuestionValidatorCalls {
				t.Errorf("Expected %d validator calls, got %d", scenario.expectedQuestionValidatorCalls, questionValidatorCalled)
			}
		})
	}
	questionValidator = validateQuestion
}

func TestQuestionService_Get(t *testing.T) {
	type testCase struct {
		description           string
		input                 string
		repository            questionRepositoryMock
		expectedErr           error
		expectedFindCallInput string
		expectedFindCalls     int
	}
	var findCallInput string
	var findCalls int

	for _, scenario := range []testCase{
		{
			description: "questionFound_success",
			input:       questionId,
			repository: questionRepositoryMock{
				fakeFind: func(id string) (domain.Question, error) {
					findCallInput = id
					findCalls++
					return domain.Question{}, nil
				},
			},
			expectedFindCallInput: questionId,
			expectedFindCalls:     1,
		},
		{
			description: "questionNotFound_returnsError",
			input:       questionId,
			repository: questionRepositoryMock{
				fakeFind: func(id string) (domain.Question, error) {
					findCallInput = id
					findCalls++
					return domain.Question{}, _errors.NoRecordFound
				},
			},
			expectedErr:           QuestionNotFoundError,
			expectedFindCallInput: questionId,
			expectedFindCalls:     1,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			findCallInput = ""
			findCalls = 0
			s := NewQuestionService(&scenario.repository)

			_, err := s.Get(scenario.input)

			if !errors.Is(err, scenario.expectedErr) {
				t.Errorf("Expected %v, got %v", scenario.expectedErr, err)
			}
			if findCallInput != scenario.expectedFindCallInput {
				t.Errorf("Expected find call input %s, got %s", scenario.expectedFindCallInput, findCallInput)
			}
			if findCalls != scenario.expectedFindCalls {
				t.Errorf("Expected find calls %d, got %d", scenario.expectedFindCalls, findCalls)
			}
		})
	}
}

func TestQuestionService_Delete(t *testing.T) {
	type testCase struct {
		description             string
		input                   string
		repository              questionRepositoryMock
		expectedDeleteCallInput string
		expectedDeleteCalls     int
	}
	var deleteCallInput string
	var deleteCalls int

	for _, scenario := range []testCase{
		{
			description: "repositoryDeleteCalled_success",
			input:       questionId,
			repository: questionRepositoryMock{
				fakeDelete: func(id string) {
					deleteCallInput = id
					deleteCalls++
				},
			},
			expectedDeleteCallInput: questionId,
			expectedDeleteCalls:     1,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			deleteCallInput = ""
			deleteCalls = 0
			s := NewQuestionService(&scenario.repository)

			s.Delete(scenario.input)

			if deleteCallInput != scenario.expectedDeleteCallInput {
				t.Errorf("Expected find call input %s, got %s", scenario.expectedDeleteCallInput, deleteCallInput)
			}
			if deleteCalls != scenario.expectedDeleteCalls {
				t.Errorf("Expected find calls %d, got %d", scenario.expectedDeleteCalls, deleteCalls)
			}
		})
	}
}

func TestQuestionService_GetAll(t *testing.T) {
	type testCase struct {
		description          string
		input                string
		repository           questionRepositoryMock
		expectedFindAllCalls int
		expectedReturnSize   int
	}
	var findAllCalls int

	for _, scenario := range []testCase{
		{
			description: "repositoryFindAllCalled_success",
			repository: questionRepositoryMock{
				fakeFindAll: func() []domain.Question {
					findAllCalls++
					return make([]domain.Question, 0)
				},
			},
			expectedFindAllCalls: 1,
			expectedReturnSize:   0,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			findAllCalls = 0
			s := NewQuestionService(&scenario.repository)

			questions := s.GetAll()

			if len(questions) != scenario.expectedReturnSize {
				t.Errorf("Expected no of questions %d, got %d", scenario.expectedReturnSize, len(questions))
			}
			if findAllCalls != scenario.expectedFindAllCalls {
				t.Errorf("Expected findAll calls %d, got %d", scenario.expectedFindAllCalls, findAllCalls)
			}
		})
	}
}

func TestQuestionService_Add(t *testing.T) {
	type testCase struct {
		description       string
		input             domain.Question
		validator         func(question domain.Question) error
		expectedSaveCalls int
		expectedErr       error
	}
	var saveCalls int
	var saveCallInputId string

	r := questionRepositoryMock{
		fakeSave: func(question domain.Question) domain.Question {
			saveCalls++
			saveCallInputId = question.Id
			return question
		},
	}
	s := NewQuestionService(&r)

	for _, scenario := range []testCase{
		{
			description:       "incorrectQuestion_returnsError",
			input:             domain.Question{},
			validator:         validateQuestion,
			expectedSaveCalls: 0,
			expectedErr:       QuestionDescriptionEmptyError,
		},
		{
			description: "incomingQuestionHasId_repositorySaveIsCalledWithEmptyId",
			input: domain.Question{
				Id: "123",
			},
			validator: func(question domain.Question) error {
				return nil
			},
			expectedSaveCalls: 1,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			saveCalls = 0
			saveCallInputId = ""

			questionValidator = scenario.validator
			_, err := s.Add(scenario.input)

			if !errors.Is(err, scenario.expectedErr) {
				t.Errorf("Expected error %v, got %v", scenario.expectedErr, err)
			}
			if saveCallInputId != "" {
				t.Errorf("Expected repository save called with empty id, but got %s", saveCallInputId)
			}
			if saveCalls != scenario.expectedSaveCalls {
				t.Errorf("Expected save calls %d, got %d", scenario.expectedSaveCalls, saveCalls)
			}
		})
	}
	questionValidator = validateQuestion
}

func TestQuestionService_Edit(t *testing.T) {
	type testCase struct {
		description       string
		input             domain.Question
		validator         func(question domain.Question) error
		expectedSaveCalls int
		expectedErr       error
	}
	var saveCalls int

	r := questionRepositoryMock{
		fakeSave: func(question domain.Question) domain.Question {
			saveCalls++
			return question
		},
	}
	s := NewQuestionService(&r)

	for _, scenario := range []testCase{
		{
			description:       "incorrectQuestion_returnsError",
			input:             domain.Question{},
			validator:         validateQuestionEdit,
			expectedSaveCalls: 0,
			expectedErr:       QuestionIdEmptyError,
		},
		{
			description: "validatorIsOk_callsRepositorySave",
			input:       domain.Question{},
			validator: func(question domain.Question) error {
				return nil
			},
			expectedSaveCalls: 1,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			saveCalls = 0

			questionEditValidator = scenario.validator
			_, err := s.Edit(scenario.input)

			if !errors.Is(err, scenario.expectedErr) {
				t.Errorf("Expected error %v, got %v", scenario.expectedErr, err)
			}
			if saveCalls != scenario.expectedSaveCalls {
				t.Errorf("Expected save calls %d, got %d", scenario.expectedSaveCalls, saveCalls)
			}
		})
	}
	questionEditValidator = validateQuestionEdit
}

func getOptionMock() domain.Option {
	return domain.Option{
		Description: optionDescription,
	}
}
