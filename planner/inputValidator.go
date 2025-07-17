package planner

import "errors"

func ValidateInput(input []string) ([]string, error) {
	if ok, err := ValidateEmptyInput(input); !ok {
		return nil, err
	}
	if ok, err := ValidateTooManyInputs(input); !ok {
		return nil, err
	}
	if ok, err := ValidateInvalidInput(input); !ok {
		return nil, err
	}
	if ok, err := ValidateDuplicateInput(input); !ok {
		return nil, err
	}
	return input, nil
}

func ValidateEmptyInput(input []string) (output bool, err error) {
	if len(input) == 0 {
		return false, errors.New("input cannot be empty")
	}
	return true, nil
}

func ValidateTooManyInputs(input []string) (output bool, err error) {
	if len(input) > 4 {
		return false, errors.New("input cannot be more than 4 body parts")
	}
	return true, nil
}

var allowedBodyParts = map[string]bool{
	"legs":       true,
	"arms":       true,
	"chest":      true,
	"core":       true,
	"hamstrings": true,
	"calves":     true,
	"shoulders":  true,
	"back":       true,
}

func ValidateInvalidInput(input []string) (output bool, err error) {
	for _, word := range input {
		if !allowedBodyParts[word] {
			return false, errors.New("invalid input word: " + word)
		}
	}
	return true, nil
}

func ValidateDuplicateInput(input []string) (output bool, err error) {
	used := map[string]bool{}
	for _, word := range input {
		if used[word] {
			return false, errors.New("entered more than once: " + word)
		}
		used[word] = true
	}
	return true, nil
}
