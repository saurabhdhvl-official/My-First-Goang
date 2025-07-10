package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter an any Number :- ")
	fmt.Scanln(&a)

	if a%5 == 0 {
		if a%7 == 0 {
			fmt.Printf("Number %d is Divisible by both 5 and 7\n", a)
		} else {
			fmt.Printf("Number %d is Divisible by 5 !!\n", a)
		}
	} else {
		fmt.Printf("Number %d is not Divisible by both 5 and 7\n", a)
	}
}
