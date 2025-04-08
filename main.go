package main

import (
	"fmt"
)

//Struct
type person struct {
	name string
	age int
}

func main() {
	person := person{name: "John", age: 25}

	fmt.Printf("This is out person %+v\n", person)
	
	// anonymous Struct
	employee := struct {
		name string
		id int
	}{
		name: "alice",
		id: 123,
	}

	fmt.Println("This is employee", employee)
}

