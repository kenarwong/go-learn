package flowcontrols

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// no parentheses
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// short statement before execution
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// v is available in else statements
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// return v // v is undefined
	return lim
}

func If() {
	fmt.Println("If Statements")
	fmt.Println("------")

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
