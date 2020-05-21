package method

import "fmt"

type VertexID int

func (id VertexID) String() string {
	return fmt.Sprintf("VertextID(%d)", id)
}

func ExampleVertexID_print() {
	i := VertexID(100)
	fmt.Println(i)
	// Output:
	// VertextID(100)
}
