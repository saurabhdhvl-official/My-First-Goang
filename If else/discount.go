package main

import (
	"fmt"
	"strconv"
)

func main() {
	var Price float64 = 3200.0
	var Discount float64 = 18.0

	var discountAmount = (Price * Discount) / 100
	var finalPrice = Price - discountAmount

	finalStr := strconv.FormatFloat(finalPrice, 'f', 1, 64)

	fmt.Println("Final Price After 18x% Discount :-  ₹" + finalStr)
}
