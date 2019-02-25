package main

import (
	"fmt"
)

var a int
var b string
var c bool
var d bool = true

func main() {
	e := 42
	f := "Different strings in go"
	g := `This is not an ordainary string` //This is a raw literal string
	h := `Can you say, "type is life?"`

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(f, h)

}
