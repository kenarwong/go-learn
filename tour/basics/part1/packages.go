package part1

import (
	"fmt"
	"math"
	"math/rand"
)

func Packages() {
	fmt.Println("Packages")
	fmt.Println("------")

	// Import math/rand
	fmt.Println("My favorite number is", rand.Intn(10))

	// Printf
	// Import just math
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))

	// Exported names are always capitalized
	//fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
