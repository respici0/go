package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// how to use a pointer
	// jimPointer := &jim // &variable - appersand gives the memory address of the value the variable is pointing at
	// jimPointer.updateName("Jimmy")

	// shortcut pointer - how to write
	jim.updateName("Jimmy")
	jim.print()
}

// example of how to use a pointer in a function
// *person - is a type description, it means we're working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer - gives the value of this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
