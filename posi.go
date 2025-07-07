/*
package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter a Number :- ")
	fmt.Scanln(&a)

	if a > 0 {
		fmt.Println("It's a Positive Number !!")
	} else {
		fmt.Println("It's a Negative Number !!")
	}
}
*/

package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter a Number :- ")
	fmt.Scanln(&a)

	if a > 0 {
		fmt.Println("It's a Positive Number !!")
	} else if a < 0 {
		fmt.Println("It's a Negative Number !!")
	} else {
		fmt.Println("Zero is Niether Positive nor Negative !!")
	}
}
