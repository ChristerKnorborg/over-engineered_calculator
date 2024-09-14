package main

import (
	"fmt"
	"overengineered_calculator/calculator"
)

func main() {

	calc := calculator.Calculator{}

	resultAdd := calc.Add(1, 2)
	fmt.Println("Add 1 and 2: ", resultAdd)

	resultSubtract := calc.Subtract(1, 2)
	fmt.Println("Subtract 1 and 2: ", resultSubtract)

	resultMultiply := calc.Multiply(5, 2)
	fmt.Println("Multiply 5 and 2: ", resultMultiply)

	resultDivide, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Divide 10 and 2: ", resultDivide)
	}

	resultModulo, err := calc.Modulo(10, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Modulo 10 and 3: ", resultModulo)
	}

	resultPower := calc.Power(2, 3)
	fmt.Println("Power 2 and 3: ", resultPower)

	fmt.Println("History: ", calc.History)
}
