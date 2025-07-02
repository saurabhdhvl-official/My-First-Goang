package main

import (
	"fmt"
	"strconv"
)

func main() {
	var marks int = 87
	var percentage float64 = 87.5

	result := float64(marks) + percentage
	resultstr := strconv.FormatFloat(result, 'f', 1, 64)

	fmt.Println("Totol Score Is :- ", resultstr)

}
