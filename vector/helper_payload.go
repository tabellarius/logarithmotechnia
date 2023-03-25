package vector

import (
	"golang.org/x/exp/constraints"
	"math"
	"math/cmplx"
	"time"
)

func pickValueWithNA[T any](idx int, data []T, na []bool, maxLen int) any {
	if idx < 1 || idx > maxLen {
		return nil
	}

	if na[idx-1] {
		return nil
	}

	return any(data[idx-1])
}

func dataWithNAToInterfaceArray[T any](data []T, na []bool) []any {
	dataLen := len(data)
	outData := make([]any, dataLen)

	for idx, val := range data {
		if na[idx] {
			outData[idx] = nil
		} else {
			outData[idx] = val
		}
	}

	return outData
}

func byIndicesWithNA[T any](indices []int, srcData []T, srcNA []bool, naDef T) ([]T, []bool) {
	data := make([]T, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		if idx == 0 {
			data = append(data, naDef)
			na = append(na, true)
		} else {
			data = append(data, srcData[idx-1])
			na = append(na, srcNA[idx-1])
		}
	}

	return data, na
}

func byIndicesWithoutNA[T any](indices []int, srcData []T, naDef T) []T {
	data := make([]T, len(indices))

	for i, idx := range indices {
		if idx == 0 {
			data[i] = naDef
		} else {
			data = append(data, srcData[idx-1])
		}
	}

	return data
}

func adjustToLesserSizeWithNA[T any](srcData []T, srcNA []bool, size int) ([]T, []bool) {
	data := make([]T, size)
	na := make([]bool, size)

	copy(data, srcData)
	copy(na, srcNA)

	return data, na
}

func adjustToLesserSizeWithoutNA[T any](srcData []T, size int) []T {
	data := make([]T, size)

	copy(data, srcData)

	return data
}

func adjustToBiggerSizeWithNA[T any](srcData []T, srcNA []bool, length int, size int) ([]T, []bool) {
	cycles := size / length
	if size%length > 0 {
		cycles++
	}

	data := make([]T, cycles*length)
	na := make([]bool, cycles*length)

	for i := 0; i < cycles; i++ {
		copy(data[i*length:], srcData)
		copy(na[i*length:], srcNA)
	}

	data = data[:size]
	na = na[:size]

	return data, na
}

func adjustToBiggerSizeWithoutNA[T any](srcData []T, length int, size int) []T {
	cycles := size / length
	if size%length > 0 {
		cycles++
	}

	data := make([]T, cycles*length)

	for i := 0; i < cycles; i++ {
		copy(data[i*length:], srcData)
	}

	data = data[:size]

	return data
}

func supportsWhicher[T any](whicher any) bool {
	if _, ok := whicher.(func(int, T, bool) bool); ok {
		return true
	}

	if _, ok := whicher.(func(T, bool) bool); ok {
		return true
	}

	if _, ok := whicher.(func(T) bool); ok {
		return true
	}

	return false
}

func whichWithNA[T any](inData []T, inNA []bool, whicher any) []bool {
	if byFunc, ok := whicher.(func(int, T, bool) bool); ok {
		return selectByFuncWithNA(inData, inNA, byFunc)
	}

	if byFunc, ok := whicher.(func(T, bool) bool); ok {
		return selectByCompactFuncWithNA(inData, inNA, byFunc)
	}

	if byFunc, ok := whicher.(func(T) bool); ok {
		return selectByBriefFuncWithNA(inData, inNA, byFunc)
	}

	return make([]bool, len(inData))
}

func whichWithoutNA[T any](inData []T, whicher any) []bool {
	if byFunc, ok := whicher.(func(int, T) bool); ok {
		return selectByFuncWithoutNA(inData, byFunc)
	}

	if byFunc, ok := whicher.(func(T) bool); ok {
		return selectByCompactFuncWithoutNA(inData, byFunc)
	}

	return make([]bool, len(inData))
}

func selectByFuncWithNA[T any](inData []T, inNA []bool, byFunc func(int, T, bool) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(idx+1, val, inNA[idx])
	}

	return booleans
}

func selectByFuncWithoutNA[T any](inData []T, byFunc func(int, T) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(idx+1, val)
	}

	return booleans
}

func selectByCompactFuncWithNA[T any](inData []T, inNA []bool, byFunc func(T, bool) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(val, inNA[idx])
	}

	return booleans
}

func selectByCompactFuncWithoutNA[T any](inData []T, byFunc func(T) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(val)
	}

	return booleans
}

func selectByBriefFuncWithNA[T any](inData []T, inNA []bool, byFunc func(T) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		if !inNA[idx] && byFunc(val) {
			booleans[idx] = true
		}
	}

	return booleans
}

func apply[T any](inData []T, inNA []bool, applier any, options []Option) Payload {
	if data, na, ok := applyType[T, bool](inData, inNA, applier, false); ok {
		return BooleanPayload(data, na, options...)
	}

	if data, na, ok := applyType[T, int](inData, inNA, applier, 0); ok {
		return IntegerPayload(data, na, options...)
	}

	if data, na, ok := applyType[T, float64](inData, inNA, applier, math.NaN()); ok {
		return FloatPayload(data, na, options...)
	}

	if data, na, ok := applyType[T, complex128](inData, inNA, applier, cmplx.NaN()); ok {
		return ComplexPayload(data, na, options...)
	}

	if data, na, ok := applyType[T, string](inData, inNA, applier, ""); ok {
		return StringPayload(data, na, options...)
	}

	if data, na, ok := applyType[T, time.Time](inData, inNA, applier, time.Time{}); ok {
		return TimePayload(data, na, options...)
	}

	if data, na, ok := applyType[T, any](inData, inNA, applier, nil); ok {
		return AnyPayload(data, na, options...)
	}

	return NAPayload(len(inData))
}

func applyType[T, S any](inData []T, inNA []bool, applier any, naDef S) ([]S, []bool, bool) {
	if applyFunc, ok := applier.(func(int, T, bool) (S, bool)); ok {
		data, na := applyByFunc[T, S](inData, inNA, applyFunc, naDef)
		return data, na, true
	}

	if applyFunc, ok := applier.(func(T, bool) (S, bool)); ok {
		data, na := applyByCompactFunc[T, S](inData, inNA, applyFunc, naDef)
		return data, na, true
	}

	if applyFunc, ok := applier.(func(T) S); ok {
		data, na := applyByBriefFunc[T, S](inData, inNA, applyFunc, naDef)
		return data, na, true
	}

	return nil, nil, false
}

func applyByFunc[T, S any](inData []T, inNA []bool,
	applyFunc func(int, T, bool) (S, bool), naDef S) ([]S, []bool) {
	length := len(inData)

	data := make([]S, length)
	na := make([]bool, length)

	for i := 0; i < length; i++ {
		dataVal, naVal := applyFunc(i+1, inData[i], inNA[i])
		if naVal {
			dataVal = naDef
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func applyByCompactFunc[T, S any](inData []T, inNA []bool,
	applyFunc func(T, bool) (S, bool), naDef S) ([]S, []bool) {
	length := len(inData)

	data := make([]S, length)
	na := make([]bool, length)

	for i := 0; i < length; i++ {
		dataVal, naVal := applyFunc(inData[i], inNA[i])
		if naVal {
			dataVal = naDef
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func applyByBriefFunc[T, S any](inData []T, inNA []bool,
	applyFunc func(T) S, naDef S) ([]S, []bool) {
	length := len(inData)

	data := make([]S, length)
	na := make([]bool, length)

	for i := 0; i < length; i++ {
		dataVal := naDef
		naVal := true
		if !inNA[i] {
			dataVal = applyFunc(inData[i])
			naVal = false
		}

		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func applyTo[T any](indices []int, inData []T, inNA []bool, applier any, naDef T) ([]T, []bool) {
	var data []T = nil
	var na []bool = nil

	switch applier.(type) {
	case func(int, T, bool) (T, bool):
		data, na = applyToByFunc(indices, inData, inNA, applier.(func(int, T, bool) (T, bool)), naDef)
	case func(T, bool) (T, bool):
		data, na = applyToByCompactFunc(indices, inData, inNA, applier.(func(T, bool) (T, bool)), naDef)
	case func(T) T:
		data, na = applyToByBriefFunc(indices, inData, inNA, applier.(func(T) T))
	case T:
		data, na = applyToByValue(indices, inData, inNA, applier.(T))
	}

	return data, na
}

func applyToByFunc[T any](indices []int, inData []T, inNA []bool,
	applyFunc func(int, T, bool) (T, bool), naDef T) ([]T, []bool) {
	length := len(inData)

	data := make([]T, length)
	na := make([]bool, length)

	copy(data, inData)
	copy(na, inNA)

	for _, idx := range indices {
		idx = idx - 1
		dataVal, naVal := applyFunc(idx+1, inData[idx], inNA[idx])
		if naVal {
			dataVal = naDef
		}
		data[idx] = dataVal
		na[idx] = naVal
	}

	return data, na
}

func applyToByCompactFunc[T any](indices []int, inData []T, inNA []bool,
	applyFunc func(T, bool) (T, bool), naDef T) ([]T, []bool) {
	length := len(inData)

	data := make([]T, length)
	na := make([]bool, length)

	copy(data, inData)
	copy(na, inNA)

	for _, idx := range indices {
		idx = idx - 1
		dataVal, naVal := applyFunc(inData[idx], inNA[idx])
		if naVal {
			dataVal = naDef
		}
		data[idx] = dataVal
		na[idx] = naVal
	}

	return data, na
}

func applyToByValue[T any](indices []int, inData []T, inNA []bool, val T) ([]T, []bool) {
	data := make([]T, len(inData))
	na := make([]bool, len(inData))

	copy(data, inData)
	copy(na, inNA)

	for _, idx := range indices {
		idx = idx - 1
		if inNA[idx] {
			continue
		}
		data[idx] = val
	}

	return data, na
}

func traverse[T any](inData []T, inNA []bool, traverser any) {
	length := len(inData)

	if fn, ok := traverser.(func(int, T, bool)); ok {
		for i := 0; i < length; i++ {
			fn(i, inData[i], inNA[i])
		}
	}

	if fn, ok := traverser.(func(T, bool)); ok {
		for i := 0; i < length; i++ {
			fn(inData[i], inNA[i])
		}
	}

	if fn, ok := traverser.(func(T)); ok {
		for i := 0; i < length; i++ {
			if !inNA[i] {
				fn(inData[i])
			}
		}
	}
}

func applyToByBriefFunc[T any](indices []int, inData []T, inNA []bool, applyFunc func(T) T) ([]T, []bool) {
	length := len(inData)

	data := make([]T, length)
	na := make([]bool, length)

	copy(data, inData)
	copy(na, inNA)

	for _, idx := range indices {
		idx = idx - 1
		if !inNA[idx] {
			data[idx] = applyFunc(inData[idx])
		}
	}

	return data, na
}

func groupsForData[T comparable](srcData []T, srcNA []bool) ([][]int, []any) {
	groupMap := map[T][]int{}
	ordered := []T{}
	na := []int{}

	for i, val := range srcData {
		idx := i + 1

		if srcNA[i] {
			na = append(na, idx)
			continue
		}

		if _, ok := groupMap[val]; !ok {
			groupMap[val] = []int{}
			ordered = append(ordered, val)
		}

		groupMap[val] = append(groupMap[val], idx)
	}

	groups := make([][]int, len(ordered))
	for i, val := range ordered {
		groups[i] = groupMap[val]
	}

	if len(na) > 0 {
		groups = append(groups, na)
	}

	values := make([]any, len(groups))
	for i, val := range ordered {
		values[i] = any(val)
	}
	if len(na) > 0 {
		values[len(values)-1] = nil
	}

	return groups, values
}

func groupsForDataWithHash[T any, S comparable](srcData []T, srcNA []bool, fnHash func(T) S) ([][]int, []any) {
	groupMap := map[S][]int{}
	ordered := []T{}
	na := []int{}

	for i, val := range srcData {
		idx := i + 1
		h := fnHash(val)

		if srcNA[i] {
			na = append(na, idx)
			continue
		}

		if _, ok := groupMap[h]; !ok {
			groupMap[h] = []int{}
			ordered = append(ordered, val)
		}

		groupMap[h] = append(groupMap[h], idx)
	}

	groups := make([][]int, len(ordered))
	for i, val := range ordered {
		groups[i] = groupMap[fnHash(val)]
	}

	if len(na) > 0 {
		groups = append(groups, na)
	}

	values := make([]any, len(groups))
	for i, val := range ordered {
		values[i] = any(val)
	}
	if len(na) > 0 {
		values[len(values)-1] = nil
	}

	return groups, values
}

func supportsSummarizer[T any](summarizer any) bool {
	if _, ok := summarizer.(func(int, T, T, bool) (T, bool)); ok {
		return true
	}

	return false
}

func summarize[T any](inData []T, inNA []bool, summarizer any, valInit T, naDef T) (T, bool) {
	fn, ok := summarizer.(func(int, T, T, bool) (T, bool))
	if ok {
		return summarizeByFunc(inData, inNA, fn, valInit, naDef)
	}

	return naDef, true
}

func summarizeByFunc[T any](inData []T, inNA []bool, fn func(int, T, T, bool) (T, bool), valInit T, naDef T) (T, bool) {
	val := valInit
	na := false

	for i := 0; i < len(inData); i++ {
		val, na = fn(i+1, val, inData[i], inNA[i])
		if na {
			return naDef, true
		}
	}

	return val, false
}

func eq[T comparable](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum == v
		}
	}

	return cmp
}

func eqFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool), eqFn func(T, T) bool) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = eqFn(datum, v)
		}
	}

	return cmp
}

func neq[T comparable](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
	cmp := make([]bool, len(inData))
	v, ok := convertor(val)

	if !ok {
		for i := range inData {
			cmp[i] = true
		}

		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = true
		} else {
			cmp[i] = datum != v
		}
	}

	return cmp
}

func neqFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool), eqFn func(T, T) bool) []bool {
	cmp := make([]bool, len(inData))
	v, ok := convertor(val)

	if !ok {
		for i := range inData {
			cmp[i] = true
		}

		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = true
		} else {
			cmp[i] = !eqFn(datum, v)
		}
	}

	return cmp
}

func gt[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum > v
		}
	}

	return cmp
}

func gtFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
	ltFn func(T, T) bool) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = ltFn(v, datum)
		}
	}

	return cmp
}

func lt[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum < v
		}
	}

	return cmp
}

func ltFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
	ltFn func(T, T) bool) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = ltFn(datum, v)
		}
	}

	return cmp
}

func gte[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum >= v
		}
	}

	return cmp
}

func gteFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
	eqFn func(T, T) bool, ltFn func(T, T) bool) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = ltFn(v, datum) || eqFn(datum, v)
		}
	}

	return cmp
}

func lte[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum <= v
		}
	}

	return cmp
}

func lteFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
	eqFn func(T, T) bool, ltFn func(T, T) bool) []bool {
	cmp := make([]bool, len(inData))

	v, ok := convertor(val)
	if !ok {
		return cmp
	}

	for i, datum := range inData {
		if inNA[i] {
			cmp[i] = false
		} else {
			cmp[i] = ltFn(datum, v) || eqFn(v, datum)
		}
	}

	return cmp
}

func find[T comparable](needle any, inData []T, inNA []bool, convertor func(any) (T, bool)) int {
	val, ok := convertor(needle)
	if !ok {
		return 0
	}

	for i, datum := range inData {
		if !inNA[i] && val == datum {
			return i + 1
		}
	}

	return 0
}

func findFn[T any](needle any, inData []T, inNA []bool, convertor func(any) (T, bool), eqFn func(T, T) bool) int {
	val, ok := convertor(needle)
	if !ok {
		return 0
	}

	for i, datum := range inData {
		if !inNA[i] && eqFn(val, datum) {
			return i + 1
		}
	}

	return 0
}

func findAll[T comparable](needle any, inData []T, inNA []bool, convertor func(any) (T, bool)) []int {
	val, ok := convertor(needle)
	if !ok {
		return []int{}
	}

	found := []int{}
	for i, datum := range inData {
		if !inNA[i] && val == datum {
			found = append(found, i+1)
		}
	}

	return found
}

func findAllFn[T any](needle any, inData []T, inNA []bool, convertor func(any) (T, bool),
	eqFn func(T, T) bool) []int {
	val, ok := convertor(needle)
	if !ok {
		return []int{}
	}

	found := []int{}
	for i, datum := range inData {
		if !inNA[i] && eqFn(val, datum) {
			found = append(found, i+1)
		}
	}

	return found
}
