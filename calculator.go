package main

import (
	"fmt"
)

func main() {
	var operator string
	var num1, num2 float64
	//Infinite loop
	for {
		fmt.Println("Calculator")

		fmt.Print("Enter the first number: ")
		fmt.Scanln(&num1)

		fmt.Print("Enter an operator (+, -, *, / or ^): ")
		fmt.Scanln(&operator)

		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)

		// var result float64 = calculate(num1, num2, operator)
		//short variable declaration syntax, dont need infer type
		result := calculate(num1, num2, operator)

		fmt.Printf("Result: %.2f\n", result)

		fmt.Print("Do you want to exit? (y,n)")
		var userResponse string
		fmt.Scanln(&userResponse)
		if !getUserResponse(userResponse) {
			break
		}
	}

}

func calculate(num1, num2 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "^":
		result = power(num1, num2)
	default:
		fmt.Println("Invalid operator")
	}

	return result
}

func power(base, exponent float64) float64 {
	result := 1.0

	for i := 0; i < int(exponent); i++ {
		result *= base
	}

	return result
}

func getUserResponse(userResp string) bool {
	switch userResp {
	case "y":
		return false
	case "n":
		return true
	default:
		return true
	}
}
