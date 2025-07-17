package main

import (
	"GymBro/planner"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Podaj partie mięśniowe oddzielone spacją: ")
	rawInput, _ := reader.ReadString('\n')

	input, err := planner.ParseInput(rawInput)
	if err != nil {
		fmt.Println("Błąd parsowania:", err)
		return
	}

	_, err = planner.ValidateInput(input)
	if err != nil {
		fmt.Println("Błąd walidacji:", err)
		return
	}

	realClient := &planner.RealGeminiClient{}

	plan, err := planner.CallGemini(realClient, input)
	if err != nil {
		fmt.Println("Błąd generowania planu:", err)
		return
	}

	response, err := planner.ParseAndValidateResponse(plan, input)
	if err != nil {
		fmt.Println("Błąd parsowania odpowiedzi z bota:", err)
		return
	}

	fmt.Println("\nWygenerowany plan treningowy:")
	fmt.Println(response.String())
}
