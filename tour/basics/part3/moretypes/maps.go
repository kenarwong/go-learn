package moretypes

import "fmt"

func Maps() {
	fmt.Println("Maps")
	fmt.Println("------")

	type Vertex struct {
		Lat, Long float64
	}

	// Maps map key to values
	// Zero value map is nil
	// nil map has no keys, nor can keys be added
	var m map[string]Vertex

	// make function returns a map of a given type, ready to use
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals use keys
	m = map[string]Vertex{
		"Bell Labs": {
			Lat:  40.68433,
			Long: -74.39967,
		},
		"Google": {
			Lat:  37.42202,
			Long: -122.08408,
		},
	}

	fmt.Println(m)

	// Mutating maps
	mapToMutate := make(map[string]int)

	// Insert
	mapToMutate["Answer"] = 42
	fmt.Println("The value:", mapToMutate["Answer"])

	// Update
	mapToMutate["Answer"] = 48
	fmt.Println("The value:", mapToMutate["Answer"])

	// Delete
	delete(mapToMutate, "Answer")
	fmt.Println("The value:", mapToMutate["Answer"])

	// Test
	// If doesn't exist, elem will default to map element type
	elem, ok := mapToMutate["Answer"]
	fmt.Println("The value:", elem, "Present?", ok)
}
