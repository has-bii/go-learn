package main

import "fmt"

func main() {
	// Every type has a zero value.
	// We can see the zero values for the following types:
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %+v\n", a, a)
	fmt.Printf("var b %T = %q\n", b, b)
	fmt.Printf("var c %T = %+v\n", c, c)
	fmt.Printf("var d %T = %+v\n\n", d, d)

	// Note: Any allocated variable is usable
	// even if it never has a value assigned.
}
