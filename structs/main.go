package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type name string

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
	}
	jim.UpdateName("Hüseyin")
	jim.PrintValues("**************************")

}

func (p person) PrintValues(footer string) {

	fmt.Println("**************************")
	fmt.Printf("%+v\n\n", p)
	fmt.Println(footer)

}

func (p *person) UpdateName(newFirstName string) {
	(*p).firstName = newFirstName
}
