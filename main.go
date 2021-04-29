package main

import (
	"fmt"
	"logarithmotechnia.com/logarithmotechnia/vector"
)

func main() {
	vec := vector.NewIntegerVector([]int{100, 200, 300, 400, 500})
	vec.SetName("two", 2)

	fmt.Println(vec)
}
