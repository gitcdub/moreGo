package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("this is a func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("Grorge Orwell's:", x)
	}
	g(1984)
}
