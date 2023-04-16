package vector

import (
	"fmt"
	"logarithmotechnia/internal/util"
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

func trueBooleanArr(size int) []bool {
	booleans := make([]bool, size)

	for i := 0; i < size; i++ {
		booleans[i] = true
	}

	return booleans
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
		ok = util.EqualComplexArrays(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *floatPayload:
		p1 := vec1.payload.(*floatPayload)
		p2 := vec2.payload.(*floatPayload)
		ok = util.EqualFloatArrays(p1.data, p2.data) && reflect.DeepEqual(p1.na, p2.na)
	case *anyPayload:
		p1 := vec1.payload.(*anyPayload)
		p2 := vec2.payload.(*anyPayload)
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
	case *vectorPayload:
		p1 := vec1.payload.(*vectorPayload)
		p2 := vec2.payload.(*vectorPayload)
		ok = CompareVectorArrs(p1.data, p2.data)
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
			fmt.Println(arr1[i])
			fmt.Println(arr2[i])
			return false
		}
	}

	return true
}

func Combine(vecs ...Vector) Vector {
	if len(vecs) == 0 {
		return nil
	}

	cmbVec := vecs[0]
	for i := 1; i < len(vecs); i++ {
		cmbVec = cmbVec.Append(vecs[i])
	}

	return cmbVec
}
