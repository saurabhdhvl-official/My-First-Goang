package main

import "fmt"

func main() {

	temp := 30

	if temp < 10 {
		fmt.Println("Very Cold.")
	} else if temp >= 10 && temp <= 20 {
		fmt.Println("Cold")
	} else if temp > 20 && temp < 30 {
		fmt.Println("Warm")
	} else {
		fmt.Println("Hot")
	}
}
