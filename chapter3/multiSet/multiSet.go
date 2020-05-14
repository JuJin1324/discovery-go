package multiSet

func NewMultiSet() map[string]int {
	return make(map[string]int)
}

func Insert(m map[string]int, val string) {
	m[val]++
}

func Erase(m map[string]int, val string) {
	if m[val]--; m[val] < 0 {
		m[val] = 0
	}
}

func Count(m map[string]int, val string) int {
	return m[val]
}

func String(m map[string]int) string {
	s := "{ "
	for key, value := range m {
		for i := 0; i < value; i++ {
			s += key + " "
		}
	}
	s += "}"

	return s
}
