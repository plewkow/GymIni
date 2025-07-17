package planner

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockGeminiClient struct {
	//Response string
	//Err      error
	mock.Mock
}

func (m *MockGeminiClient) GeneratePlan(input []string) (string, error) {
	args := m.Called(input)
	return args.String(0), args.Error(1)
}

func TestAiCaller(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		mockResponse   string
		mockErr        error
		expectedOutput string
		expectedErr    string
	}{
		{
			name:  "normal input",
			input: []string{"arms", "legs", "back", "abs"},

			mockResponse: "good plan",
			mockErr:      nil,

			expectedOutput: "good plan",
			expectedErr:    "",
		},
		{
			name:  "api error",
			input: []string{"arms"},

			mockResponse: "",
			mockErr:      fmt.Errorf("mock API failure"),

			expectedOutput: "",
			expectedErr:    "mock API failure",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockClient := new(MockGeminiClient)

			mockResponse := test.mockResponse

			mockClient.ExpectedCalls = nil
			mockClient.On("GeneratePlan", mock.Anything).Return(mockResponse, test.mockErr)
			plan, err := mockClient.GeneratePlan(test.input)
			if err != nil {
				return
			}

			//mockClient := &MockGeminiClient{
			//	//Response: test.mockResponse,
			//	//Err:      test.mockErr,
			//
			//}

			//output, err := CallGemini(mockClient, test.input)

			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
				assert.Empty(t, plan)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedOutput, plan)
			}
			mockClient.AssertExpectations(t)
		})
	}
}
