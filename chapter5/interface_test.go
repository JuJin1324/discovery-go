package chapter5

import (
	"fmt"
)

// Shape ...
type Shape interface {
	Area() float64
}

// Object ...
type Object interface {
	Volume() float64
}

// Skin ...
type Skin interface {
	Color() float64
}

// Cube ...
type Cube struct {
	side float64
}

// Area ...
func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

// Volume ...
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func Example_interface() {
	var s Shape = Cube{3}
	value1, ok1 := s.(Object)
	fmt.Printf("dynamic value of Shape s with value %v implements interface object? %v\n", value1, ok1)
	value2, ok2 := s.(Skin)
	fmt.Printf("dynamic value of Shape s with value %v implements interface object? %v\n", value2, ok2)
	// Output:
	//
}
