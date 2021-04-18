package main

import (
	"fmt"
	"logarithmotechnia.com/logarithmotechnia/vector"
)

func main() {
	v := vector.NewIntegerVector([]int{1, 2, 3, 4, 5})

	fmt.Printf("%T %v\n", v, v)

	var a []int

	fmt.Printf("a: %T %v, len %v", a, a, len(a))
}
