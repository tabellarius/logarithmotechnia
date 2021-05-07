package vector

import (
	"fmt"
	"testing"
)

func TestNewIntegerVector(t *testing.T) {
	vec := NewIntegerPayload([]int{1, 2, 3, 4, 5})

	fmt.Println(vec)
	fmt.Println(vec.Names())
}
