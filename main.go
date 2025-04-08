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

	// SLICE
	allNumbers := numbers[:] // this will insert all numbers array value as allNumbers slice
	fmt.Println("This is allnumber values: ", allNumbers)

	firstThree := numbers[0:3] // this will get all 3 values from numbers array
	fmt.Println("This is firstThree", firstThree)

	fruits := []string{"apple", "banana", "strawberry"}
	fruits = append(fruits, "kiwi")
	fmt.Printf("These are my fruits %v\n", fruits)

	moreFruits := []string{"blueberries", "tomato"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("Combining 2 slices : %v\n", fruits)

  scores := make([]int, 3, 3)
	// scores = append(scores, 1)
  fmt.Println("score: ", scores, "length: ", len(scores), "cap: ", cap(scores))

	// MAP
	capitalCities := map[string]string{
		"USA": "Washington D.C",
		"India": "New Delhi",
		"UK": "London",
	}

	fmt.Println(capitalCities["USA"])
	// Checking if the key not exist
	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println("This is the capital", capital)
	} else {
		fmt.Println("Does not exist")
	}

	delete(capitalCities, "UK")
	fmt.Printf("this new deleted map: %v", capitalCities)
}

func add(a int, b int) int {
	return a + b
}

func calculateAndProducts(a, b int) (int, int) {
	return a + b, a * b
}
