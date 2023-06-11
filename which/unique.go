package which

import "logarithmotechnia/vector"

// IsUnique returns a boolean array indicating whether each element of the vector is unique.
func IsUnique(v vector.Vector) []bool {
	if uniquer, ok := v.Payload().(vector.IsUniquer); ok {
		return uniquer.IsUnique()
	}

	return trueBooleanArr(v.Len())

}

// IsNotUnique returns a boolean array indicating whether each element of the vector is not unique.
func IsNotUnique(v vector.Vector) []bool {
	if uniquer, ok := v.Payload().(vector.IsUniquer); ok {
		return invertBooleanArr(uniquer.IsUnique())
	}

	return make([]bool, v.Len())
}
