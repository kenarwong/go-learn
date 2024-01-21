package part1

import "fmt"

// Define a function add
// Takes two parameters of type int
// Return type int
func add1(x int, y int) int {
	return x + y
}

// If multiple parameters have the same type, can exclude the type on all except last one
func add2(x, y int) int {
	return x + y
}

// A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Define named values to be returned
// return statements without variables is a "naked" return, should only be used in short functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Functions() {
	fmt.Println("Functions")
	fmt.Println("------")

	fmt.Println(add1(42, 13))
	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
