package planner

import (
	"GymBro/models"
	"encoding/json"
	"errors"
	"strings"
)

func ParseAndValidateResponse(jsonStr string, expectedBodyParts []string) (*models.TrainingPlan, error) {
	var trainingPlan models.TrainingPlan
	err := json.Unmarshal([]byte(jsonStr), &trainingPlan)
	if err != nil {
		return nil, err
	}

	foundParts := make(map[string]bool)

	for _, day := range trainingPlan.TrainingPlan {
		normalized := strings.ToLower(strings.TrimSpace(day.BodyPart))
		foundParts[normalized] = true
	}

	for _, inputPart := range expectedBodyParts {
		if !foundParts[strings.ToLower(inputPart)] {
			return nil, errors.New("missing body part in response: " + inputPart)
		}
	}

	return &trainingPlan, nil
}
