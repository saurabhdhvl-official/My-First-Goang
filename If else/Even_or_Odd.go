package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter a Number :- ")
	fmt.Scanln(&a)

	if a%2 == 0 {
		fmt.Println("It's a Even Number !!")
	} else {
		fmt.Println("It's a odd Number !!")
	}
}
