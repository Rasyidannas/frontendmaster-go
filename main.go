package main

import (
	"fmt"
)

func main() {
	// CONDITIONAL
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

	// LOOP
	for i := 0; i < 5; i++ {
		fmt.Println("this is i", i)
	}

	counter := 0
	for counter < 3 {
		fmt.Println("This is the counter ", counter)
		counter++
	}

	// ARRAY
	numbers := [5]int{10, 20, 30, 40, 5}
	// numbers := [...]int{10, 20, 30, 40 ,5} // this is same like abovebut use ...
	numbers[4] = 60 // reassign new value

	fmt.Println("This is the all length of numbers ", len(numbers))
	
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("This is the matrix: ", matrix)
}

func add(a int, b int) int {
	return a + b
}

func calculateAndProducts(a, b int) (int, int) {
	return a + b, a * b
}
