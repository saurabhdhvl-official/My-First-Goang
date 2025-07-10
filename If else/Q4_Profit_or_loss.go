package main

import "fmt"

func main() {
	var sell float64
	var cost float64

	fmt.Println("Enter Selling Price :- ")
	fmt.Scanln(&sell)

	fmt.Println("Enter a Cost Price :- ")
	fmt.Scanln(&cost)

	if sell > cost {
		fmt.Printf("Profit Is :-  Rs. %.2f\n", sell-cost)
	} else if sell < cost {
		fmt.Printf("Loss Is :-  Rs. %.2f\n", cost-sell)
	} else {
		fmt.Printf("\nNo profit No Loss !!")
	}
}
