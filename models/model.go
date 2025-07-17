package models

import (
	"fmt"
	"strings"
)

type TrainingPlan struct {
	TrainingPlan map[string]Day `json:"training_plan"`
}

type Day struct {
	BodyPart string     `json:"body_part"`
	Workout  []Exercise `json:"workout"`
}

type Exercise struct {
	Exercise    string `json:"exercise"`
	Sets        string `json:"sets"`
	Reps        string `json:"reps"`
	KgsGuidance string `json:"kgs_guidance"`
}

func (tp TrainingPlan) String() string {
	var b strings.Builder
	for day, d := range tp.TrainingPlan {
		b.WriteString(fmt.Sprintf("%s - %s\n", strings.Title(day), d.BodyPart))
		for _, w := range d.Workout {
			b.WriteString(fmt.Sprintf("  - Exercise: %s Sets x reps: %s x %s Kgs: (%s)\n", w.Exercise, w.Sets, w.Reps, w.KgsGuidance))
		}
		b.WriteString("\n")
	}
	return b.String()
}
