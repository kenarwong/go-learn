package zrt

import (
	"fmt"
	"math"
)

// Round to precision
func Round(f float64, p float64) float64 {
	return math.Round(f*math.Pow(10, p)) / math.Pow(10, p)
}

func Sqrt(x float64) {
	// Estimation precision
	var precision = float64(7)

	// Initial guesses
	// Convert to float64 to match type for calculations
	// z := float64(1)
	z := x / 2

	fmt.Printf("Iteration 1: %g\n", z)

	// Start iteration on second guess
	for i := 2; i < 11; i++ {
		zOld := z
		// Formula given to adjust guess
		// If with initial statement to break if estimation isn't changing
		if z -= (z*z - x) / (2 * z); Round(zOld, precision) == Round(z, precision) {
			break
		}
		fmt.Printf("Iteration %v: %g\n", i, z)
	}
}
