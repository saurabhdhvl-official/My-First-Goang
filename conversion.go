package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 5.5

	// convert 'a' to float64 before adding

	result := float64(a) + b

	fmt.Println("Result :- ", result)

}
