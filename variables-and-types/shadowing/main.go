package main

import "fmt"

func main() {
	// When a variable is declared in an outer scope and another variable
	// with the same name is declared inside a function or block,
	// the inner variable shadows the outer one.

	x := 100
	fmt.Println(x)

	if true {
		x := 10
		fmt.Println(x)
	}

	// Outer scope: `x := 100` defines `x` in the main function's scope.
	// Inner scope: Inside the `if` block, `x := 10` creates a new variable `x`
	// that shadows the outer variable.
}
