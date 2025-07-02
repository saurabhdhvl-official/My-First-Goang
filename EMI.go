package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	principal := 150000.0
	var annualRate = 9.0
	var Months = 24.0

	var monthlyRate = annualRate / (12 * 100)
	var powerTerm = math.Pow(1+monthlyRate, Months)

	var emi = (principal * monthlyRate * powerTerm) / (powerTerm - 1)

	var emiStr = strconv.FormatFloat(emi, 'f', 1, 64)
	fmt.Println("\nFor a loan of ₹150000 over 24 months at 9% p.a.,\n")
	//fmt.Println("Your Monthly EMI Is :- " + emiStr + " ₹ ")
	fmt.Printf("Your Monthly EMI is: %.1f ₹\n", emiStr)

}
