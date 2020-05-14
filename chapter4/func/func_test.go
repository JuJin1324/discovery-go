package _func

import "fmt"

func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

func ExampleAddOne() {
	n := []int{1, 2, 3, 4}
	AddOne(n)
	fmt.Println(n)
	// Output:
	// [2 3 4 5]
}
