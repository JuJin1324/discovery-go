package set

import "fmt"

func hasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false
}

func ExampleHasDupeRune() {
	fmt.Println(hasDupeRune("숨바꼭질"))
	fmt.Println(hasDupeRune("다시합시다"))
	// Output:
	// false
	// true
}
