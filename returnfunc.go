package main

import (
	"fmt"
)

func main() {
	//	x := bar()
	//	fmt.Printf("%T\n", x)
	//	fmt.Println(x())

	// above code returns the same as below

	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 451
	}
}
