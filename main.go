package main

import (
	"fmt"
)

func main() {
	// decalre variables
	var name string = "Rasyid"
	fmt.Printf("This is my name %s\n", name) // for formating and string interpolation
	fmt.Println("My name is ", name)

	// short variable declaration operator (auto detected type) 
	age := 26
	fmt.Printf("May age is %d\n", age)

	// empty string
	var city string
	city = "Sidoarjo"
	fmt.Printf("I live in %s\n", city)

	hobbies := ""
	hobbies = "Reading"
	fmt.Printf("I like do %s\n", hobbies)
	
	// multiple declartion variables
	var country, continent string = "Indonesia", "East Java"
	fmt.Printf("I live in country %s and this is my continent %s\n", country, continent)

	var (
		isEmployed bool = true
		salary int = 50000
		position string = "developer" 
	)

	fmt.Printf("isEmployed: %t this is my salary: %d and this is my position %s\n", isEmployed, salary, position)

	// Zero Values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("int: %d float: %f  string: %s and bool:%t \n", defaultInt, defaultFloat, defaultString, defaultBool)
	
	// constant variables
	const pi = 3.14

	// multiple contant variables
	const (
		Monday = 1
		Tuesday = 2
		Wednesday = 3
	)

	fmt.Printf("Monday: %d - Tuesday: %d - Wednesday: %d\n", Monday, Tuesday, Wednesday)

	// iota for auto increments
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
	)

	fmt.Printf("jan - %d feb- %d mar - %d apr - %d\n", Jan, Feb, Mar, Apr)

	// calling function
	result := add(3, 4)
	fmt.Println("This is  the result", result)

	sum, product := calculateAndProducts(10, 10)
	fmt.Printf("this is sum: %d and this is product: %d\n", sum, product)
}

func add(a int, b int) int {
	return a + b
}

func calculateAndProducts(a, b int) (int, int) {
	return a + b, a * b
}
