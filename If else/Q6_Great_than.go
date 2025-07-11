package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Println("Enter first  Number :- ")
	fmt.Scanln(&a)

	fmt.Println("Enter second Number :- ")
	fmt.Scanln(&b)

	if a > b {
		fmt.Printf("A is Greater than B !!")
	} else if a < b {
		fmt.Printf("B is Greater than A !!")
	} else {
		fmt.Printf("Both NUmbers are equal !!")
	}
}
