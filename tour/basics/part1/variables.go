package part1

import "fmt"

// var statements defines lists of variables
var c, python, java bool

// var declarations can have initializers
var j, k int = 1, 2

func Variables() {
	fmt.Println("Variables")
	fmt.Println("------")

	// var statements can be at package or function level
	var i int
	fmt.Println(i, c, python, java)

	// The type can be omitted in var declarations
	var c, python, java = true, false, "no!"
	fmt.Println(j, k, c, python, java)

	// := short assignment can be used in place of a var declaration
	var l, m int = 1, 2
	n := 3
	d, e, f := true, false, "no!"
	fmt.Println(l, m, n, d, e, f)
}
