package main

import (
	"fmt"
)

func main() {
	age := 30

	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenasger")
	} else {
		fmt.Println("you are a child")
	}

	day := "Tuesday"

	switch day {
		case "Monday":
			fmt.Println("start of the week")
		case "Tuesday", "Wednesday", "Thursday":
			fmt.Println("Midweek")
		case "Friday":
			fmt.Println("TGIF")
		default:
			fmt.Println("it's this weekend")
	}


}

func add(a int, b int) int {
	return a + b
}

func calculateAndProducts(a, b int) (int, int) {
	return a + b, a * b
}
