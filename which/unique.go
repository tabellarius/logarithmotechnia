package which

import "logarithmotechnia/vector"

func IsUnique(v vector.Vector) []bool {
	if uniquer, ok := v.Payload().(vector.IsUniquer); ok {
		return uniquer.IsUnique()
	}

	return trueBooleanArr(v.Len())

}

func IsNotUnique(v vector.Vector) []bool {
	if uniquer, ok := v.Payload().(vector.IsUniquer); ok {
		return invertBooleanArr(uniquer.IsUnique())
	}

	return make([]bool, v.Len())
}
