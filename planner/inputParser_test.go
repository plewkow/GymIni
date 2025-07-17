package planner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput []string
		expectedErr    string
	}{
		{
			name:           "normal input",
			input:          "Arms Legs Back Abs",
			expectedOutput: []string{"arms", "legs", "back", "abs"},
			expectedErr:    "",
		},
		{
			name:           "extra spaces",
			input:          " Arms  Legs    Back  Abs    ",
			expectedOutput: []string{"arms", "legs", "back", "abs"},
			expectedErr:    "",
		},
		{
			name:           "empty input",
			input:          "",
			expectedOutput: nil,
			expectedErr:    "",
		},
		{
			name:           "empty input - spaces",
			input:          "   ",
			expectedOutput: nil,
			expectedErr:    "",
		},
		{
			name:           "one word capital letters",
			input:          "ARMS",
			expectedOutput: []string{"arms"},
			expectedErr:    "",
		},
		{
			name:           "mix letter size",
			input:          "ArMS",
			expectedOutput: []string{"arms"},
			expectedErr:    "",
		},
		{
			name:           "input with tabs and newlines",
			input:          "arms\tlegs\nback",
			expectedOutput: []string{"arms", "legs", "back"},
			expectedErr:    "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := ParseInput(test.input)

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
