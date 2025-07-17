package planner

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"log"
	"os"
)

type GeminiClient interface {
	GeneratePlan(input []string) (string, error)
}

func CallGemini(client GeminiClient, input []string) (string, error) {
	return client.GeneratePlan(input)
}

type RealGeminiClient struct{}

func (r *RealGeminiClient) GeneratePlan(input []string) (string, error) {
	_ = godotenv.Load()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set in environment")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatal(err)
	}

	prompt := fmt.Sprintf(
		"Make a one week training plan for these body parts: %s. "+
			"Your response should be in JSON format, with the structure:\n"+
			"{ \"training_plan\": { \"day 1\": { \"body_part\": \"Chest\", \"workout\": [ { \"exercise\": \"Barbell Bench Press\", \"sets\": \"3-4\", \"reps\": \"8-12\", \"kgs_guidance\": \"Challenging weight\" } ] } } }.\n"+
			"Include exercises, sets, reps, and kg guidance. "+
			"Do not provide triple backticks (```) around the JSON.",
		input,
	)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return result.Text(), nil
}
