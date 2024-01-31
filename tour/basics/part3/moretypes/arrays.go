package moretypes

import "fmt"

func Arrays() {
	fmt.Println("Arrays")
	fmt.Println("------")

	// Type [n]T is an array of n values of type T
	// Array length part of its type, cannot be resized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Array literal
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
