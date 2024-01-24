package flowcontrols

import "fmt"

func DeferDemo1() {
	// defer call arguments are evaluated, but not executed until surrounding function returns
	defer fmt.Println(" world")
	fmt.Print("hello")
}

func DeferDemo2() {
	// Deferred function calls are pushed onto a stack
	fmt.Println("counting")

	for i := 10; i > 0; i-- {
		// Executed in last-in-first-out order
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func Defer() {
	fmt.Println("Defer Statements")
	fmt.Println("------")

	DeferDemo1()
	DeferDemo2()
}
