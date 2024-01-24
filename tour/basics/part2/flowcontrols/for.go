package flowcontrols

import "fmt"

func For() {
	fmt.Println("For loops")
	fmt.Println("------")
	sum := 0

	// init; condition; post
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// init and post statements optional
	// for is Go's while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	// for {
	// }
}
