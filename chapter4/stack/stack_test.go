package stack

import (
	"fmt"
	"strconv"
	"strings"
)

type StrSet map[string]struct{}

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

type PrecMap map[string]StrSet
type BinOp func(int, int) int

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	ops := []string{"("}
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce(")")
	return nums[0]
}

func ExampleNewEvaluator() {
	eval := NewEvaluator(map[string]BinOp{
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}
			if b < 0 {
				return 0
			}
			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}
			return r
		},
		"+": func(a int, b int) int {
			return a + b
		},
		"-": func(a int, b int) int {
			return a - b
		},
		"*": func(a int, b int) int {
			return a * b
		},
		"/": func(a int, b int) int {
			return a / b
		},
		"mod": func(a int, b int) int {
			return a % b
		},
	}, PrecMap{
		"**":  NewStrSet(),
		"*":   NewStrSet("**", "*", "/", "mod"),
		"/":   NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+":   NewStrSet("**", "*", "/", "mod", "+", "-"),
		"-":   NewStrSet("**", "*", "/", "mod", "+", "-"),
	})

	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1 - 2 + 3"))
	fmt.Println(eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	fmt.Println(eval("1 + 2 ** 10 * 2"))
	fmt.Println(eval("2 ** 3 mod 3"))
	fmt.Println(eval("2 ** 2 ** 3"))
	// Output:
	// 5
	// 3
	// 2
	// 9
	// 18
	// 2049
	// 2
	// 256
}
