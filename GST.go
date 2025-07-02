package main

import (
	"fmt"
	"strconv"
)

func main() {
	var basePrice float64 = 49999.0
	var gstRate float64 = 18.0

	var gstAmount = (basePrice * gstRate) / 100
	var finalPrice = basePrice + gstAmount

	finalStr := strconv.FormatFloat(finalPrice, 'f', 1, 64)

	fmt.Println("Final Price After GST :- ₹" + finalStr)
}
