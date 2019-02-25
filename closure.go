package main

import "fmt"

var x int //scope of var x is package level

func main() {
	var y int //scope of var y is only func main
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(y)
	{
		a := 50 //scope of a is limited to the block of code
		fmt.Println(a)
	}
}

func foo() {
	fmt.Println("gopher")

}
