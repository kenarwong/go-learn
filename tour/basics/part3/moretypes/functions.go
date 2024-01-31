package moretypes

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Closures are function values that reference variables outside of its body
func adder() func(int) int {
	sum := 0 // The function can access this variable. It is "bound" to the variable
	return func(x int) int {
		sum += x
		return sum
	}
}

func Functions() {
	fmt.Println("Functions")
	fmt.Println("------")

	// Functions are values too
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Functions can be passed as arguments
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
