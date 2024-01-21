package part1

import "fmt"

// Constants are declared with const
const Pi = 3.14

// Numeric constants are high-precision values
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// A binary number that is 1 followed by 100 zeroes
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func Constants() {
	fmt.Println("Constants")
	fmt.Println("------")

	// Constants cannot be short-assigned
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// An untyped constant will take the type needed by the context
	fmt.Println(needInt(Small)) // int
	// fmt.Println(needInt(Big)) // int can store at maximum a 64-bit integer, sometimes less
	fmt.Println(needFloat(Small)) // float64
	fmt.Println(needFloat(Big))   // float64
}
