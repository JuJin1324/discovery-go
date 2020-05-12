package slice

import "fmt"

func Example_sliceCopy1() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)
	// Output:
	// [30 20 50 10 40]
}

func Example_sliceCopy2() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	copiedNumber := copy(dest, src)
	fmt.Println(copiedNumber, dest)
	// Output:
	// 5 [30 20 50 10 40]
}

func Example_sliceCopy3() {
	src := []int{30, 20, 50, 10, 40}
	dest := append([]int(nil), src...)
	fmt.Println(dest)
	// Output:
	// [30 20 50 10 40]
}
