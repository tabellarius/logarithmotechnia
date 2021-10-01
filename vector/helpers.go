package vector

import (
	"reflect"
)

func indicesArray(size int) []int {
	indices := make([]int, size)

	index := 0
	for i := range indices {
		indices[i] = index
		index++
	}

	return indices
}

func incIndices(indices []int) []int {
	for i := range indices {
		indices[i]++
	}

	return indices
}

func CompareVectorsForTest(one, two Vector) bool {
	if one == nil && two != nil || two == nil && one != nil {
		return false
	}

	if one == nil && two == nil {
		return true
	}

	vec1 := one.(*vector)
	vec2 := two.(*vector)
	if vec1.Type() != vec2.Type() {
		return false
	}

	if vec1.Len() != vec2.Len() {
		return false
	}

	var ok bool
	switch vec1.payload.(type) {
	case *booleanPayload:
		p1 := vec1.payload.(*booleanPayload)
		p2 := vec2.payload.(*booleanPayload)
		ok = reflect.DeepEqual(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *integerPayload:
		p1 := vec1.payload.(*integerPayload)
		p2 := vec2.payload.(*integerPayload)
		ok = reflect.DeepEqual(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *complexPayload:
		p1 := vec1.payload.(*complexPayload)
		p2 := vec2.payload.(*complexPayload)
		ok = reflect.DeepEqual(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *floatPayload:
		p1 := vec1.payload.(*floatPayload)
		p2 := vec2.payload.(*floatPayload)
		ok = reflect.DeepEqual(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *interfacePayload:
		p1 := vec1.payload.(*interfacePayload)
		p2 := vec2.payload.(*interfacePayload)
		ok = reflect.DeepEqual(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *naPayload:
		ok = vec1.length == vec2.length
	case *stringPayload:
		p1 := vec1.payload.(*stringPayload)
		p2 := vec2.payload.(*stringPayload)
		ok = reflect.DeepEqual(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *timePayload:
		p1 := vec1.payload.(*timePayload)
		p2 := vec2.payload.(*timePayload)
		ok = reflect.DeepEqual(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	default:
		ok = false
	}

	return ok
}

func CompareVectorArrs(arr1, arr2 []Vector) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if !CompareVectorsForTest(arr1[i], arr2[i]) {
			return false
		}
	}

	return true
}
