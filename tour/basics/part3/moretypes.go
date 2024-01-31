package part3

import (
	"fmt"
	"tour/basics/part3/drawpic"
	"tour/basics/part3/fibonacci"
	"tour/basics/part3/moretypes"
	"tour/basics/part3/wordcount"
)

func MoreTypes() {
	fmt.Println("More Types")
	fmt.Println("------")
	moretypes.Pointers()
	fmt.Println()
	moretypes.Structs()
	fmt.Println()
	moretypes.Arrays()
	fmt.Println()
	moretypes.Slices()
	fmt.Println()
	moretypes.Range()
	fmt.Println()
	fmt.Println("Exercise: 2-D Picture Function")
	fmt.Println("------")
	drawpic.DrawPic()
	fmt.Println()
	moretypes.Maps()
	fmt.Println()
	fmt.Println("Exercise: Count Words")
	fmt.Println("------")
	wordcount.TestWordCount()
	fmt.Println()
	moretypes.Functions()
	fmt.Println()
	fmt.Println("Exercise: Fibonacci")
	fmt.Println("------")
	fibonacci.RunFibonacci()
	fmt.Println()
}
