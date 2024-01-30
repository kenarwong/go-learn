package part2

import (
	"fmt"
	"tour/basics/part2/flowcontrols"
	"tour/basics/part2/zrt"
)

func Flow() {
	fmt.Println("Flow Control Statements")
	fmt.Println("------")
	flowcontrols.For()
	fmt.Println()
	flowcontrols.If()
	fmt.Println()
	fmt.Println("Exercise: Square Root Estimation")
	fmt.Println("------")
	zrt.Sqrt(28)
	fmt.Println()
	flowcontrols.Switch()
	fmt.Println()
	flowcontrols.Defer()
	fmt.Println()
}
