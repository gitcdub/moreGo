package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface { //type is keyword human is identifier
	speak() // interface is the type
	// a value can be of more then one type
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{ //line 28-34 is of type secretAgent and human
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Maybe",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	bar(sa1)
	bar(sa2)
	bar(p1)
}
