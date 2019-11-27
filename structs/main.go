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

type personA struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Bad way of declaring struct since order is dependent
	// on struct definition order....
	// peter := person{"Peter", "Pan"}
	// fmt.Println(peter)
	// fmt.Printf("%+v\n", peter)

	// Better
	alex := person{firstName: "Alex",
		lastName: "Anderson"}

	fmt.Printf("My first name is %v and my lastname is %v.\n",
		alex.firstName, alex.lastName)
	fmt.Printf("%+v\n", alex)

	// More generic
	var robin person

	// Note the zero values (empty string)
	fmt.Printf("%+v\n", robin)

	robin.firstName = "Robin"
	robin.lastName = "Huiser"
	fmt.Printf("%+v\n", robin)

	// Now with contactInfo as a separate field
	david := personA{
		firstName: "David",
		lastName:  "Huiser",
		contact: contactInfo{
			email:   "david@google.com",
			zipCode: 654321,
		},
	}

	fmt.Printf("%+v\n", david)

	// Now with contactInfo as same name field
	jim := person{
		firstName: "Jim",
		lastName:  "Henderson",
		contactInfo: contactInfo{
			email:   "jim@google.com",
			zipCode: 12356,
		},
	}

	// This does not work (copy of variable)
	// jim.updateName("Jimmy")

	// This works - but a pain to write...
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	// This is a shortcut in Go - allows us to
	// call function with pointer or byValue!
	jim.updateName("Jimmy")
	jim.print()
}

// This does not work since p is a COPY of the object, not the object itself!
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// This does work
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
