package method

import (
	"fmt"
	"strings"
)

type MultiSet map[string]int

func (m MultiSet) Insert(val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "{ "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}

func ExampleMultiSet() {
	m := MultiSet{}
	fmt.Println(m)
	fmt.Println(m.Count("3"))
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	fmt.Println(m)
	fmt.Println(m.Count("3"))
	m.Insert("1")
	m.Insert("2")
	m.Insert("5")
	m.Insert("7")
	m.Erase("3")
	m.Erase("5")
	fmt.Println(m.Count("3"))
	fmt.Println(m.Count("1"))
	fmt.Println(m.Count("2"))
	fmt.Println(m.Count("5"))
	// Output:
	// { }
	// 0
	// { 3 3 3 3 }
	// 4
	// 3
	// 1
	// 1
	// 0
}
