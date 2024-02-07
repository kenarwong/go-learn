package interfaces

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures
type Abser interface {
	Abs() float64
}

// Interfaces are implemented implicitly
// There is no "implements" keyword
// This decouples the definition of an interface from its implementation
type I interface {
	M()
}

type T struct {
	S string
}

// This method automatically means that T implements I
// No need to declare explicitly
func (t T) M() {
	fmt.Println(t.S)
}

type U struct {
	S string
}

func (u *U) M() {
	fmt.Println(u.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

type W struct {
	S string
}

func (w *W) M() {
	// In go, it is common to write methods that gracefully handle being called when the receiver is nil
	if w == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(w.S)
}

func Interfaces() {
	fmt.Println("Interfaces")
	fmt.Println("------")

	// A value of an interface type can hold any value that implements those methods
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // MyFloat implements Abser
	fmt.Println(a.Abs())

	// Vertex has a pointer receiver, so it does not implement Abser
	//a = v // This will not compile

	// But, the value of a pointer of type Vertex does implement Abser
	a = &v // *Vertex implements Abser
	fmt.Println(a.Abs())

	// T implements I
	var i I = T{"hello"}
	i.M()

	// Interface values can be thought of as a tuple with a value and and a specific underlying concrete type: (value, type)
	i = &U{"hello"}

	// Call a method on an interface executes the method of the same name on its underlying type
	describe(i) // (&{hello}, *interfaces.U)
	i.M()

	i = F(math.Pi)
	describe(i) // (3.141592653589793, interfaces.F)
	i.M()

	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver
	var w *W // nil pointer

	i = w       // Set the interface value
	describe(i) // (<nil>, *interfaces.V)

	// Note that even though the receiver is nil, the method is still called with a non-nil value
	// The interface value that holds the nil receiver is non-nil
	i.M()

	w = &W{"hello"}
	i = w
	describe(i)
	i.M()

	// A nil interface value holds neither value nor concrete type
	var j I
	describe(j) // (<nil>, <nil>)
	//j.M() // Calling a method on a nil interface results in a runtime error. There is no underlying concrete type to indicate which method to call

	// An empty interface may hold values of any type
	var k interface{}
	describeAny(k) // (<nil>, <nil>)

	// Every type implements at least zero methods
	k = 42
	describeAny(k)
	k = "hello"
	describeAny(k)

	// Empty interfaces are used by code that handles values of unknown types
	// For example, fmt.Print takes any number of arguments of type interface{}
	fmt.Print("Hello, ", 42, "!\n")
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(i I) {
	// Describe the underlying value and type of the interface (tuple)
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeAny(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
