package main

import "fmt"

func main() {

	var temp int

	fmt.Println("Enter the Temperature in °C :- ")
	fmt.Scanln(&temp)

	fmt.Printf("You entered %d°C :- It's ", temp)

	if temp < 10 {
		fmt.Println("Very Cold")
	} else if temp >= 10 && temp <= 20 {
		fmt.Println("Cold")
	} else if temp > 20 && temp < 30 {
		fmt.Println("Warm")
	} else {
		fmt.Println("Hot")
	}
}
