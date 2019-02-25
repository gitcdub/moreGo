package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var P person // p is person type

	P.name = "Jim"                                  // assign "Jim" to field 'name' of p
	P.age = 30                                      // assign 25 to field 'age' of p
	fmt.Printf("The person's name is %s\n", P.name) // access field 'name' of p
	fmt.Println("Jim is", P.age)
}
