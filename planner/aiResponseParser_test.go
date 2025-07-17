package planner

import (
	"GymBro/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseResponse(t *testing.T) {
	tests := []struct {
		name           string
		jsonResponse   string
		inputParts     []string
		expectedOutput *models.TrainingPlan
		expectedErr    string
	}{
		{
			name: "normal input - one day",
			jsonResponse: `{
				"training_plan": {
					"day 1": { "body_part": "Arms", "workout": [{ "exercise": "Light Walk", "sets": "1", "reps": "30-60 minutes", "kgs_guidance": "Low intensity" }] }
				}
			}`,
			expectedOutput: &models.TrainingPlan{
				TrainingPlan: map[string]models.Day{
					"day 1": {
						BodyPart: "Arms",
						Workout: []models.Exercise{
							{
								Exercise:    "Light Walk",
								Sets:        "1",
								Reps:        "30-60 minutes",
								KgsGuidance: "Low intensity",
							},
						},
					},
				},
			},
			expectedErr: "",
		},
		{
			name: "all parts present",
			jsonResponse: `{
				"training_plan": {
					"day 1": { "body_part": "Arms", "workout": [{ "exercise": "Light Walk", "sets": "1", "reps": "30-60 minutes", "kgs_guidance": "Low intensity" }] }
				}
			}`,
			expectedOutput: &models.TrainingPlan{
				TrainingPlan: map[string]models.Day{
					"day 1": {
						BodyPart: "Arms",
						Workout: []models.Exercise{
							{
								Exercise:    "Light Walk",
								Sets:        "1",
								Reps:        "30-60 minutes",
								KgsGuidance: "Low intensity",
							},
						},
					},
				},
			},
			inputParts:  []string{"arms"},
			expectedErr: "",
		},
		{
			name: "missing body part",
			jsonResponse: `{
				"training_plan": {
					"day 1": { "body_part": "Arms", "workout": [{ "exercise": "Light Walk", "sets": "1", "reps": "30-60 minutes", "kgs_guidance": "Low intensity" }] }
				}
			}`,
			expectedOutput: &models.TrainingPlan{
				TrainingPlan: map[string]models.Day{
					"day 1": {
						BodyPart: "Arms",
						Workout: []models.Exercise{
							{
								Exercise:    "Light Walk",
								Sets:        "1",
								Reps:        "30-60 minutes",
								KgsGuidance: "Low intensity",
							},
						},
					},
				},
			},
			inputParts:  []string{"arms", "legs"},
			expectedErr: "missing body part in response: legs",
		},
		{
			name: "multiple days",
			jsonResponse: `{
				"training_plan": {
					"day 1": {
						"body_part": "Legs",
						"workout": [
							{ "exercise": "Squats", "sets": "4", "reps": "10", "kgs_guidance": "Heavy" }
						]
					},
					"day 2": {
						"body_part": "Back",
						"workout": [
							{ "exercise": "Deadlift", "sets": "3", "reps": "8", "kgs_guidance": "Challenging" }
						]
					}
				}
			}`,
			expectedOutput: &models.TrainingPlan{
				TrainingPlan: map[string]models.Day{
					"day 1": {
						BodyPart: "Legs",
						Workout: []models.Exercise{
							{
								Exercise:    "Squats",
								Sets:        "4",
								Reps:        "10",
								KgsGuidance: "Heavy",
							},
						},
					},
					"day 2": {
						BodyPart: "Back",
						Workout: []models.Exercise{
							{
								Exercise:    "Deadlift",
								Sets:        "3",
								Reps:        "8",
								KgsGuidance: "Challenging",
							},
						},
					},
				},
			},
			expectedErr: "",
		},
		{
			name: "invalid JSON format",
			jsonResponse: `{
				"training_plan": {
					"day 1": {
						"body_part": "Chest",
						"workout": [
							{ "exercise": "Bench Press", "sets": "3", "reps": "10", "kgs_guidance": "Medium" }
						]
					}
			}`,
			expectedOutput: nil,
			expectedErr:    "unexpected end of JSON input",
		},
		{
			name: "missing body_part",
			jsonResponse: `{
				"training_plan": {
					"day 1": {
						"workout": [
							{ "exercise": "Pull-up", "sets": "3", "reps": "10", "kgs_guidance": "Bodyweight" }
						]
					}
				}
			}`,
			expectedOutput: &models.TrainingPlan{
				TrainingPlan: map[string]models.Day{
					"day 1": {
						BodyPart: "",
						Workout: []models.Exercise{
							{
								Exercise:    "Pull-up",
								Sets:        "3",
								Reps:        "10",
								KgsGuidance: "Bodyweight",
							},
						},
					},
				},
			},
			expectedErr: "",
		},
		{
			name: "empty training plan",
			jsonResponse: `{
				"training_plan": {}
			}`,
			expectedOutput: &models.TrainingPlan{
				TrainingPlan: map[string]models.Day{},
			},
			expectedErr: "",
		},
		{
			name: "missing training_plan key",
			jsonResponse: `{
			"plan": {
				"day 1": {
					"body_part": "Chest",
					"workout": []
					}
				}
			}`,
			expectedOutput: &models.TrainingPlan{
				TrainingPlan: nil,
			},
			expectedErr: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := ParseAndValidateResponse(test.jsonResponse, test.inputParts)

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
