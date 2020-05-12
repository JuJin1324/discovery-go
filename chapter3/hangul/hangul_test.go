package hangul

import (
	"fmt"
	"strconv"
)

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	//	Output:
	// false
	// true
	// false
}

func Example_printBytes() {
	s := "가나다"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"

	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2] += 4
	fmt.Println(string(b))
	// Output:
	// 간나다
}

func Example_strCat() {
	s := "adc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// adcdef
	// adcdef
}

func Example_strToNumber() {
	var i int
	var k int64
	var f float64
	var s string
	i, _ = strconv.Atoi("350")
	fmt.Println(i)
	k, _ = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(k)
	k, _ = strconv.ParseInt("0xcc7fdd", 0, 32)
	fmt.Println(k)
	f, _ = strconv.ParseFloat("3.14", 64)
	fmt.Println(f)
	s = strconv.Itoa(340)
	fmt.Println(s)
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s)
	var num int
	_, _ = fmt.Sscanf("57", "%d", &num)
	fmt.Println(num)
	s = fmt.Sprint(3.14)
	fmt.Println(s)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(s)
	// Output:
	// 350
	// 13402077
	// 13402077
	// 3.14
	// 340
	// cc7fdd
	// 57
	// 3.14
	// cc7fdd
}

func Example_array() {
	fruits := [...]string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		var suffix string
		if HasConsonantSuffix(fruit) {
			suffix = "은"
		} else {
			suffix = "는"
		}
		fmt.Printf("%s%s 맛있다.\n", fruit, suffix)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 수박은 맛있다.
}

func Example_slice() {
	fruits := make([]string, 3)
	fruits[0] = "사과"
	fruits[1] = "바나나"
	fruits[2] = "토마토"

	for _, fruit := range fruits[:len(fruits)-1] {
		var suffix string
		if HasConsonantSuffix(fruit) {
			suffix = "은"
		} else {
			suffix = "는"
		}
		fmt.Printf("%s%s 맛있다.\n", fruit, suffix)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
}

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
}

func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	// Output:
	// [사과 바나나 토마토]
	// [포도 딸기]
	// [사과 바나나 토마토 포도 딸기]
	// [사과 바나나 포도 딸기]
}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1)) // 3
	fmt.Println("cap:", cap(sliced1)) // 5
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2)) // 3
	fmt.Println("cap:", cap(sliced2)) // 3
	fmt.Println()

	sliced3 := nums[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3)) // 4
	fmt.Println("cap:", cap(sliced3)) // 5
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
	// Output:
	// [1 2 3 4 5]
	// len: 5
	// cap: 5
	//
	// [1 2 3]
	// len: 3
	// cap: 5
	//
	// [3 4 5]
	// len: 3
	// cap: 3
	//
	// [1 2 3 4]
	// len: 4
	// cap: 5
	//
	// [1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}
