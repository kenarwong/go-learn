package part1

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Define variables of various types
// var declarations can be factored into blocks, like import statements
// int, uint, and uintptr are 32 bit on 32-bit systems and 64 bit on 64-bit systems. Use int unless there is a specific reason to use a sized or unsigned integer type
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Types() {
	fmt.Println("Types")
	fmt.Println("------")

	fmt.Printf("Type: %T, value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, value: %v\n", z, z)

	// Variables without an initial value are given their zero value
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %v\n", i, f, b, s)

	// T(v) converts value v to type T
	// Explicit conversion is necessary for variable assignment
	var x, y int = 3, 4
	f = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Variable declarations without explicit type are inferred from the value
	v1 := 42 // int
	fmt.Printf("The type of v1 is %T and the value is %v\n", v1, v1)
	v2 := 3.142 // float64
	fmt.Printf("The type of v2 is %T and the value is %v\n", v2, v2)
	v3 := 0.867 + 0.5i // complex128
	fmt.Printf("The type of v3 is %T and the value is %v\n", v3, v3)
}
