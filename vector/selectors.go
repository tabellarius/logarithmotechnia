package vector

func Odd() func(int, int, bool) bool {
	return func(idx int, _ int, _ bool) bool {
		return idx%2 == 1
	}
}

func Even() func(int, int, bool) bool {
	return func(idx int, _ int, _ bool) bool {
		return idx%2 == 0
	}
}

func Nth(n int) func(int, int, bool) bool {
	return func(idx int, _ int, _ bool) bool {
		return idx%n == 0
	}
}
