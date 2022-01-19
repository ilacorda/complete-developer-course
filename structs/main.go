package main

import "fmt"

func main() {
	aContact := contactInfo{email: "name.surname@someStuff.com", zipCode: 23455}

	jim := person{firstname: "Jim", lastname: "Anderson", contact: aContact}

	fmt.Println(jim)

	// testing structs as receivers
	// testing how pointers work
	jim.updateName("Jimmy")
	jim.print()

	// Testing pointers with slices
	mySlice := []string{"Hi", "There", "All", "Good", "?"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

// structs
type person struct {
	firstname string
	lastname  string
	// example of an embedded struct
	contact contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p *person) print() {
	fmt.Printf("%+v", p)

}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstname = newFirstName
}

func updateSlice(s []string) {
	s[0] = "Boo"
}
