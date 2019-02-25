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

// correct use of methods func(r receiver) identifier(params)
// (returns(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
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

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
