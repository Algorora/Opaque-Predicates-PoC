package main

import (
	"fmt"
)

func OpaquePredicate1(x int) bool {
	return (x*x + x) > x
}

func OpaquePredicate2(x int) bool {
	return (x^x) == 0
}

func PerformOperation(x int) int {
	if OpaquePredicate1(x) {
		if OpaquePredicate2(x) {
			return x * 2
		}
	}
	return x - 1
}

func main() {
	x := 10
	result := PerformOperation(x)
	fmt.Printf("Result: %d\n", result)
}
