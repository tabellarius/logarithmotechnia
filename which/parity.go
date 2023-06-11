package which

import "logarithmotechnia/vector"

func Even(v vector.Vector) []bool {
	booleans := make([]bool, v.Len())

	for i := 0; i < v.Len(); i++ {
		booleans[i] = i%2 == 1
	}

	return booleans
}

func Odd(v vector.Vector) []bool {
	booleans := make([]bool, v.Len())

	for i := 0; i < v.Len(); i++ {
		booleans[i] = i%2 == 0
	}

	return booleans
}

func Nth(v vector.Vector, nth int) []bool {
	booleans := make([]bool, v.Len()+1)

	for i := 1; i <= v.Len(); i++ {
		if i%nth == 0 {
			booleans[i] = true
		}
	}

	return booleans[1:]
}
