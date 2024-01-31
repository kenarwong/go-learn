package moretypes

import "fmt"

func Pointers() {
	fmt.Println("Pointers")
	fmt.Println("------")
	// Pointers hold the memory address of values
	// Type *T is a pointer for to a value of T
	// Zero-value is 'nil'
	var p *int
	i, j := 42, 2701

	// The & operator generates a pointer to its operand
	p = &i // point to i

	// The * operator denotes the pointer's underlying value
	// a.k.a. "dereferencing" / "indirecting"
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer p
	fmt.Println(j) // see the new value of j

	// Go has no pointer arithmetic
}
