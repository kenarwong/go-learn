package moretypes

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Range() {
	fmt.Println("Range")
	fmt.Println("------")

	// range iterates over a slice or map
	// range over slice returns two values: index and copy
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	// Omit value
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// Omit index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
