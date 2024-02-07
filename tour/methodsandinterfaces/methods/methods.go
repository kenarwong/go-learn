package methods

import (
	"fmt"
	"math"
)

// Go doesn't have classes
type Vertex struct {
	X, Y float64
}

// Define methods on types
// A method is a function with a special receiver argument
// Abs method has a receiver of type Vertex named v
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods are just like functions
// Here is Abs written as a regular function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare a method on non-struct types
type MyFloat float64

// You can only declare a method with a receiver whose type is defined in the same package as the method
// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Declare a pointer receiver
// Pointer receivers can modify the value to which the receiver points
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Non-pointer receivers (value receiver) operates on a copy of the value of the receiver (as with normal functions)
// The Scale method must have a pointer receiver to change the Vertex value
// ScaleCopy will not change the Vertex value
func (v Vertex) ScaleCopy(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Here is Scale written as a normal function
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Here is ScaleCopy written as a normal function
func ScaleCopyFunc(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Here is Abs written as a normal function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Generally, using pointer receivers are more efficient
// Also, all methods on a given type should have either value or pointer receivers, but not a mixture of both
func (v *Vertex) AbsPointerReceiver() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods() {
	fmt.Println("Methods")
	fmt.Println("------")

	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	// No change to v
	v.ScaleCopy(10)
	fmt.Println(v.Abs())

	// Reset X and Y
	v.X = 3
	v.Y = 4

	ScaleFunc(&v, 10)
	fmt.Println(v.Abs())

	// No change to v
	ScaleCopyFunc(v, 10)
	fmt.Println(v.Abs())

	// Reset X and Y
	v.X = 3
	v.Y = 4

	// You can call a method with a pointer receiver using a value
	v.Scale(2)

	// But you must call a function with a pointer argument using a pointer
	//ScaleFunc(v, 5) // This will not compile
	ScaleFunc(&v, 5) // But you must call a function with a pointer argument using a pointer

	// Method with a pointer receiver can be called with a pointer or a value
	p := &Vertex{3, 4}
	p.Scale(3)

	// Function with a pointer argument must be called with a pointer
	ScaleFunc(p, 8)

	// Reset X and Y
	v.X = 3
	v.Y = 4

	// The same ting happens in reverse with value receivers
	// Value receiver and arguments, must be called with a value
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p.X = 4
	p.Y = 3

	// And you can call a method with a value receiver using a pointer
	fmt.Println(p.Abs())
	//fmt.Println(AbsFunc(p)) // This will not compile
	fmt.Println(AbsFunc(*p)) // But you must call a func with a value argument using a value

	// Reset X and Y
	p.X = 3
	p.Y = 4

	// Printf specifier +v is used to print the struct fields
	fmt.Printf("Before scaling: %+v, Abs: %v\n", p, p.AbsPointerReceiver())
	p.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", p, p.AbsPointerReceiver())
}
