package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

// func main() {
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Anderson",
// 		contact: contactInfo{
// 			email:   "janderson@a.com",
// 			zipCode: 94000,
// 		},
// 	}
// 	fmt.Println(jim)
// 	fmt.Printf("%+v", jim)
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "janderson@a.com",
			zipCode: 94000,
		},
	}
	// fmt.Println(jim)
	// jimPointer := &jim
	jim.updateName("Michael")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
