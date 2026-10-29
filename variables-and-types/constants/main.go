package main

import (
	"fmt"
	"time"
)

func main() {
	// Constants in Go are immutable values that are known at compile time.
	// So the value is can't be changed.

	// Single constant
	const Pi = 3.14159
	fmt.Printf("Pi = %.5f\n", Pi)

	// Multiple constants
	const (
		StatusOK    = 200
		StatusError = 500
	)
	fmt.Printf("StatusOK = %d\n", StatusOK)
	fmt.Printf("StatusError = %d\n", StatusError)

	// Typed Constants
	const (
		Timeout time.Duration = 30 * time.Second
		MaxSize int           = 1024
		Debug   bool          = false
	)
	fmt.Printf("Timeout = %v\n", Timeout)
	fmt.Printf("MaxSize = %d\n", MaxSize)
	fmt.Printf("Debug = %v\n", Debug)

	// Iota
	// Go's iota is used in const declarations to simplify
	// definitions of incrementing numbers. It can be used in expressions as well.

	type ByteSize float64
	const (
		_ = iota // identified as 0
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Printf("KB = %v\n", KB)
	fmt.Printf("MB = %v\n", MB)
	fmt.Printf("GB = %v\n", GB)
	fmt.Printf("TB = %v\n", TB)
}
