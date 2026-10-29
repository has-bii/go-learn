package main

import "fmt"

func main() {
	// In Go, there are two primary ways to declare and assign variables:
	// 1. Shorthand assignment (using :=)
	// 2. Var assignment (using var keyword)

	// === Shorthand Assignment ===
	// The short form `name := value` is a convenient way to both declare
	// and initialize a variable in a single step. It automatically infers
	// the variable type from the right-hand side value.
	// - Can only be used inside functions (not at package level)
	// - Creates a new variable OR reassigns existing ones if at least one
	//   on the left side is already declared
	// - Most idiomatic for local variable declaration with initialization

	// Example: Shorthand assignment with implicit type inference
	a := 100_000 // a is inferred as int type
	fmt.Println("Shorthand assignment:", a)

	// === Var Assignment (Long Form) ===
	// The `var` keyword is Go's explicit way to declare variables.
	// Use this form when you need to declare variables at package level,
	// when you want to specify the type explicitly, or when initializing
	// later rather than immediately.
	// - Available everywhere (functions and package level)
	// - Requires type specification if not initializing immediately
	// - More verbose but more explicit about intent

	// Example 1: var with explicit type, then assignment
	var b int    // declare b as int (zero value: 0)
	b = 200_000  // assign value separately
	fmt.Println("Var assignment (separate):", b)

	// Example 2: var with immediate initialization (less idiomatic but valid)
	// var c int = 300_000  // explicit type + initialization
	// fmt.Println("Var with init:", c)

	// Example 3: var with short initialization (idiomatic for package-level)
	var d = 400_000 // type inferred, equivalent to: var d int = 400_000
	fmt.Println("Var with short init:", d)
}
