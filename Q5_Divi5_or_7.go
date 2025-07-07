/*

package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter an Any Number :- ")
	fmt.Scanln(&a)

	if a%5 == 0 && a%7 == 0 {
		fmt.Printf("Divisible by 5 or 7 !!")
	} else {
		fmt.Printf("It's not Divisible by 5 or 7 !!")
	}
}

*/

/*

// Logic 2

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

*/

// logic 3

package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter Any Number :- ")
	fmt.Scanln(&a)

	if a%5 == 0 {
		if a%7 == 0 {
			fmt.Printf("Number %d is Divisible by both 5 and 7 !!", a)
		} else {
			fmt.Printf("Number %d is Divisible by 5 !!", a)
		}
	} else {
		if a%7 == 0 {
			fmt.Printf("Number %d is Divisible by 7 !!", a)
		} else {
			fmt.Printf("Number %d is not Divisible by both 5 and 7 !!", a)
		}
	}
}
