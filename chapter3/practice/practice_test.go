package practice

import (
	"fmt"
	"github.com/jujin/discovery-go/chapter3/multiSet"
	"strings"
)

func less(a, b int) bool {
	return a > b
}

// 5 1 3 2 4
//
func sortAsc(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if less(a[i], a[j]) {
				swap(a, i, j)
			}
		}
	}
}

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func Example_sortAsc() {
	a := []int{5, 1, 3, 2, 4}
	fmt.Println("Before sort:", a)
	sortAsc(a)
	fmt.Println("After sort:", a)
	// Output:
	// Before sort: [5 1 3 2 4]
	// After sort: [1 2 3 4 5]
}

func find(s []string, str string) bool {
	var isFound bool
	for i := 0; i < len(s); i++ {
		if isFound = strings.EqualFold(s[i], str); isFound == true {
			break
		}
	}
	return isFound
}

func Example_findString() {
	s := []string{"abcd", "efgh", "ijkl", "mnop", "qrst"}
	abcd := "abcd"
	fmt.Println(find(s, abcd))
	notFound := "not found"
	fmt.Println(find(s, notFound))
	// Output:
	// true
	// false
}

func Example_queue() {
	array := [5]int{1, 2, 3, 4, 5}
	var q []int
	q = append(q, array[0:]...)
	array[3] = 100
	fmt.Println(q)
	q = q[1:]
	fmt.Println(q)
	q = q[1:]
	fmt.Println(q)
	q = q[1:]
	fmt.Println(q)
	q = q[1:]
	fmt.Println(q)
	q = q[1:]
	fmt.Println(q)
	// Output:
	//
}

func ExampleMultiSet() {
	m := multiSet.NewMultiSet()
	fmt.Println(multiSet.String(m))
	fmt.Println(multiSet.Count(m, "3"))
	multiSet.Insert(m, "3")
	multiSet.Insert(m, "3")
	multiSet.Insert(m, "3")
	multiSet.Insert(m, "3")
	fmt.Println(multiSet.String(m))
	fmt.Println(multiSet.Count(m, "3"))
	multiSet.Insert(m, "1")
	multiSet.Insert(m, "2")
	multiSet.Insert(m, "5")
	multiSet.Insert(m, "7")
	multiSet.Erase(m, "3")
	multiSet.Erase(m, "5")
	fmt.Println(multiSet.Count(m, "3"))
	fmt.Println(multiSet.Count(m, "1"))
	fmt.Println(multiSet.Count(m, "2"))
	fmt.Println(multiSet.Count(m, "5"))
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
