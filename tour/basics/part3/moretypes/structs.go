package moretypes

import "fmt"

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

// Struct literals denote newly allocated struct values by listing the values of its fields
// Name: syntax, order irrelevant
// &Struct returns a pointer to the value instead
var (
	v1      = Vertex{1, 2}
	v2      = Vertex{X: 1}  // Y: 0 is implicit
	v3      = Vertex{}      // X: 0 and Y: 0
	structP = &Vertex{1, 2} // has type *Vertex
)

func Structs() {
	fmt.Println("Structs")
	fmt.Println("------")

	fmt.Println(Vertex{1, 2})

	// Struct fields are accessed using a dot
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Struct fields can be accessed through struct pointers
	p := &v
	p.X = 1e9 // Equivalent to (*p).X
	fmt.Println(v)

	fmt.Println(v1, structP, v2, v3)
}
