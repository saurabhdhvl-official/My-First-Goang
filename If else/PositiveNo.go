package main

import "fmt"

func main() {
	var n int

	fmt.Println("Enter a Number :- ")
	fmt.Scanln(&n)

	if n > 0 {
		fmt.Println("Positive Number !!")
	} else {
		fmt.Println("Its Not a Positive Number !!")
	}

}
