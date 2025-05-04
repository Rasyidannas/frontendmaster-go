package main

import (
	"fmt"
)

//Struct
type Person struct {
	Name string
	Age int
}

func main() {
	person := Person{Name: "John", Age: 25}

	fmt.Printf("This is out person %+v\n", person)
	
	// anonymous Struct
	employee := struct {
		name string
		id int
	}{
		name: "alice",
		id: 123,
	}
	
	type Address struct {
		Street string
		City string
	}
	
	// nested struct
	type Contact struct {
		Name string
		Address Address
		Phone string
	}

	contact := Contact {
		Name: "Rasyid",
		Address: Address{
			Street: "Jl Jeruck",
			City: "Sidoarjo",
		},
		Phone: "812731766123",
	}

	fmt.Println("This is contact", contact)

	fmt.Println("This is employee", employee)

	fmt.Println("Name before: ", person.Name)
	person.modifyPersonName("Rasyid")
	fmt.Println("Name after: ", person.Name)

	x := 20
	ptr := &x
	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)
	*ptr = 30
	fmt.Printf("value of x new: %d and address of x %p\n", x, ptr)
}

func(p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("inside scope: new name", p.Name)
}
