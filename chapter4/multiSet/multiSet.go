package multiSet

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

func NewMultiSet() MultiSet {
	return make(map[string]int)
}

func Insert(m MultiSet, val string) {
	m[val]++
}

func Erase(m MultiSet, val string) {
	if m[val]--; m[val] < 0 {
		m[val] = 0
	}
}

func Count(m MultiSet, val string) int {
	return m[val]
}

func String(m MultiSet) string {
	s := "{ "
	for key, value := range m {
		for i := 0; i < value; i++ {
			s += key + " "
		}
	}
	s += "}"

	return s
}

func InsertFunc(m MultiSet) func(val string) {
	return func(val string) {
		Insert(m, val)
	}
}
