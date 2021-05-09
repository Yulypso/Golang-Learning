package main

import "fmt"

type ContactInfo struct {
	Email   string
	ZipCode int
}

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

func main() {
	thierry := Person{
		FirstName: "Thierry",
		LastName:  "Khamphousone",
		ContactInfo: ContactInfo{
			Email:   "khamphousone.thierry@outlook.fr",
			ZipCode: 94200,
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

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) updateFirstName(firstName string) {
	(*p).FirstName = firstName
}
