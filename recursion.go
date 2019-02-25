package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1) // this is a sample of recursion
	n := factorial(4)          // you should write loops instead
	fmt.Println(n)

	n2 := loopFact(4)
	fmt.Println(n2)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// doing factorials with loops is better then using recursion
func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
