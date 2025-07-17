package planner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedOutput []string
		expectedErr    string
	}{
		{
			name:           "normal input",
			input:          []string{"arms", "legs", "back", "calves"},
			expectedOutput: []string{"arms", "legs", "back", "calves"},
			expectedErr:    "",
		},
		{
			name:           "empty input",
			input:          nil,
			expectedOutput: nil,
			expectedErr:    "input cannot be empty",
		},
		{
			name:           "too many inputs",
			input:          []string{"arms", "legs", "back", "calves", "chest"},
			expectedOutput: nil,
			expectedErr:    "input cannot be more than 4 body parts",
		},
		{
			name:           "duplicated inputs",
			input:          []string{"arms", "arms", "back", "calves"},
			expectedOutput: nil,
			expectedErr:    "entered more than once: arms",
		},
		{
			name:           "wrong input",
			input:          []string{"arms", "legs", "back", "calve"},
			expectedOutput: nil,
			expectedErr:    "invalid input word: calve",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := ValidateInput(test.input)

			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
				assert.Empty(t, output)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedOutput, output)
			}
		})
	}
}
