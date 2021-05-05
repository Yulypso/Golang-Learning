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
	thierry := person{
		firstName: "Thierry", lastName: "Khamphousone",
		contactInfo: contactInfo{
			email:   "khamphousone.thierry@outlook.fr",
			zipCode: 94200,
		},
	}
	thierry.print()
	/* Two ways of doing it */
	thierry.updateFirstName("thierry2")
	thierry.print()

	tpointer := &thierry
	tpointer.updateFirstName("thierry3")
	thierry.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}
