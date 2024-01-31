package moretypes

import (
	"fmt"
	"strings"
)

func Slices() {
	fmt.Println("Slices")
	fmt.Println("------")

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// A slice is a dynamically-sized, flexible view into elements of an array
	// Much more common than arrays

	// Type []T is a slice of with elements of type T
	// Slices are formed by specifying a low bound and a high bound, separated by a colon
	var pSlice []int = primes[1:4]
	fmt.Println(pSlice)

	// Slices do not store data. They only reference a section of underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	firstTwo := names[0:2]
	middleTwo := names[1:3]
	fmt.Println(firstTwo, middleTwo)

	// Changes elements in a slice modifies the corresponding elements of the underlying array
	middleTwo[0] = "XXX"
	fmt.Println(firstTwo, middleTwo)
	fmt.Println(names)

	// Slice literal builds an array, then a slice that references it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// When slicing, high and low bounds may be omitted
	// Default low bound is 0, high bound is length of slice
	slice := []int{2, 3, 5, 7, 11, 13}
	slice = slice[1:4]
	fmt.Println(slice)
	slice = slice[:2]
	fmt.Println(slice)
	slice = slice[1:]
	fmt.Println(slice)

	// Slices have length and capacity
	slice = []int{2, 3, 5, 7, 11, 13}
	printSlice(slice)

	// 0 Length
	slice = slice[:0]
	printSlice(slice)

	// Extend length
	slice = slice[:4]
	printSlice(slice)

	// Drop first two values
	// Reduces capacity by 2
	slice = slice[2:]
	printSlice(slice)

	// Extending slice beyond capacity
	// "runtime error: slice bounds out of range [:5] with capacity 4"
	// slice = slice[:5]
	// printSlice(slice)

	// Zero value of slice is nil
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice)) // Zero length/capacity
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	// Use make to build dynamically-sized arrays
	// Allocates a 0'd array and returns a slice that refers to that array
	a := make([]int, 5) // len(a)=5 cap(a)=5
	fmt.Print("a: ")
	printSlice(a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Print("b: ")
	printSlice(b)

	c := b[:2] // len(c)=2, cap(c)=5
	fmt.Print("c: ")
	printSlice(c)

	d := c[2:5] // len(d)=3, cap(d)=3
	fmt.Print("d: ")
	printSlice(d)

	// Slices can contain any type, including other slices
	// Create a tic-tac-toe board
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// Players take turns
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"

	// My loop
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}

	// Their loop
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// It is common to append new elements to a slice
	var sliceToAppend []int
	printSlice(sliceToAppend)

	// func append(s []T, vs ...T) []T
	// First parameter is slice, second parameter are elements of type T to append
	// If array is too small, a bigger array will be allocated, and the slice will point to new array

	// Works on nil slices
	sliceToAppend = append(sliceToAppend, 0)
	printSlice(sliceToAppend)

	// The slice grows as needed
	sliceToAppend = append(sliceToAppend, 1)
	printSlice(sliceToAppend)

	// The slice grows as needed
	sliceToAppend = append(sliceToAppend, 2, 3, 4)
	printSlice(sliceToAppend)
}

func printSlice(s []int) {
	// length is number of elements
	// capacity number of elements in underlying array, counting from the first element in slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
