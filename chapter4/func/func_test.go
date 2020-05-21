package _func

import (
	"bufio"
	"fmt"
	"github.com/jujin/discovery-go/chapter4/multiSet"
	"io"
	"os"
	"strings"
)

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

func WriteTo(w io.Writer, lines ...string) (n int64, err error) {
	for _, line := range lines {
		var nw int
		nw, err = fmt.Fprintln(w, line)
		n += int64(nw)
		if err != nil {
			return
		}
	}
	return
}

func ExampleWriteTo() {
	w := os.Stdout
	lines := []string{"hello", "world", "Go language"}
	_, _ = WriteTo(w, lines...)
	// Output:
	// hello
	// world
	// Go language
}

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ExampleReadFrom_Print() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// ( bill )
	// ( tom )
	// ( jane )
}

func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
}

func ExampleNewIntGenerator_multiple() {
	gen1 := NewIntGenerator()
	gen2 := NewIntGenerator()
	fmt.Println(gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1(), gen1())
	// Output:
	// 1 2 3
	// 1 2 3 4 5
	// 4 5 6 7
}

type BinOp func(int, int) int
type BinSub func(int, int) int

func BinOpToBinSub(f BinOp) BinSub {
	var count int
	return func(a int, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}

func ExampleBinOpToBinSub() {
	sub := BinOpToBinSub(func(a int, b int) int {
		return a + b
	})
	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	// Output:
	// 12
	// 12
	// 12
	// count: 3
}

func ExampleInsertFunc() {
	r := strings.NewReader("hello\nworld\nGo language")
	m := multiSet.NewMultiSet()
	if err := ReadFrom(r, multiSet.InsertFunc(m)); err != nil {
		fmt.Println(err)
	}
}

func BindMap(f multiSet.SetOp, m multiSet.MultiSet) func(val string) {
	return func(val string) {
		f(m, val)
	}
}
