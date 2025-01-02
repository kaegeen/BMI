package main

import (
	"fmt"
)

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func bmiCategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi >= 18.5 && bmi < 24.9:
		return "Normal weight"
	case bmi >= 25 && bmi < 29.9:
		return "Overweight"
	default:
		return "Obesity"
	}
}

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("===============")
	fmt.Println("Enter your weight (in kilograms) and height (in meters) to calculate your BMI.")
	fmt.Println("Type 'exit' to quit.\n")

	for {
		var weight, height float64
		var input string

		fmt.Print("Enter your weight (kg): ")
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("Goodbye! Stay healthy!")
			break
		}

		_, err := fmt.Sscanf(input, "%f", &weight)
		if err != nil || weight <= 0 {
			fmt.Println("Invalid weight. Please enter a positive number.\n")
			continue
		}

		fmt.Print("Enter your height (m): ")
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("Goodbye! Stay healthy!")
			break
		}

		_, err = fmt.Sscanf(input, "%f", &height)
		if err != nil || height <= 0 {
			fmt.Println("Invalid height. Please enter a positive number.\n")
			continue
		}

		bmi := calculateBMI(weight, height)
		category := bmiCategory(bmi)

		fmt.Printf("\nYour BMI is: %.2f\n", bmi)
		fmt.Printf("Category: %s\n\n", category)
	}
}
