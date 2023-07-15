package vector

import (
	"golang.org/x/exp/constraints"
	"logarithmotechnia/option"
	"math"
	"math/cmplx"
	"time"
)

func PickValueWithNA[T any](idx int, data []T, na []bool, maxLen int) any {
	if idx < 1 || idx > maxLen {
		return nil
	}

	if na[idx-1] {
		return nil
	}

	return any(data[idx-1])
}

func DataWithNAToInterfaceArray[T any](data []T, na []bool) []any {
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

func ByIndicesWithNA[T any](indices []int, srcData []T, srcNA []bool, naDef T) ([]T, []bool) {
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

func ByIndicesWithoutNA[T any](indices []int, srcData []T, naDef T) []T {
	data := make([]T, len(indices))

	for i, idx := range indices {
		if idx == 0 {
			data[i] = naDef
		} else {
			data[i] = srcData[idx-1]
		}
	}

	return data
}

func AdjustToLesserSizeWithNA[T any](srcData []T, srcNA []bool, size int) ([]T, []bool) {
	data := make([]T, size)
	na := make([]bool, size)

	copy(data, srcData)
	copy(na, srcNA)

	return data, na
}

func AdjustToLesserSizeWithoutNA[T any](srcData []T, size int) []T {
	data := make([]T, size)

	copy(data, srcData)

	return data
}

func AdjustToBiggerSizeWithNA[T any](srcData []T, srcNA []bool, length int, size int) ([]T, []bool) {
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

func AdjustToBiggerSizeWithoutNA[T any](srcData []T, length int, size int) []T {
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

func SupportsWhicherWithNA[T any](whicher any) bool {
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

func supportsWhicherWithoutNA[T any](whicher any) bool {
	if _, ok := whicher.(func(int, T) bool); ok {
		return true
	}

	if _, ok := whicher.(func(T) bool); ok {
		return true
	}

	return false
}

func WhichWithNA[T any](inData []T, inNA []bool, whicher any) []bool {
	if byFunc, ok := whicher.(func(int, T, bool) bool); ok {
		return SelectByFuncWithNA(inData, inNA, byFunc)
	}

	if byFunc, ok := whicher.(func(T, bool) bool); ok {
		return SelectByCompactFuncWithNA(inData, inNA, byFunc)
	}

	if byFunc, ok := whicher.(func(T) bool); ok {
		return SelectByBriefFuncWithNA(inData, inNA, byFunc)
	}

	return make([]bool, len(inData))
}

func WhichWithoutNA[T any](inData []T, whicher any) []bool {
	if byFunc, ok := whicher.(func(int, T) bool); ok {
		return SelectByFuncWithoutNA(inData, byFunc)
	}

	if byFunc, ok := whicher.(func(T) bool); ok {
		return SelectByCompactFuncWithoutNA(inData, byFunc)
	}

	return make([]bool, len(inData))
}

func SelectByFuncWithNA[T any](inData []T, inNA []bool, byFunc func(int, T, bool) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(idx+1, val, inNA[idx])
	}

	return booleans
}

func SelectByFuncWithoutNA[T any](inData []T, byFunc func(int, T) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(idx+1, val)
	}

	return booleans
}

func SelectByCompactFuncWithNA[T any](inData []T, inNA []bool, byFunc func(T, bool) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(val, inNA[idx])
	}

	return booleans
}

func SelectByCompactFuncWithoutNA[T any](inData []T, byFunc func(T) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		booleans[idx] = byFunc(val)
	}

	return booleans
}

func SelectByBriefFuncWithNA[T any](inData []T, inNA []bool, byFunc func(T) bool) []bool {
	booleans := make([]bool, len(inData))

	for idx, val := range inData {
		if !inNA[idx] && byFunc(val) {
			booleans[idx] = true
		}
	}

	return booleans
}

func ApplyWithNA[T any](inData []T, inNA []bool, applier any, options []option.Option) Payload {
	if data, na, ok := ApplyTypeWithNA[T, bool](inData, inNA, applier, false); ok {
		return BooleanPayload(data, na, options...)
	}

	if data, na, ok := ApplyTypeWithNA[T, int](inData, inNA, applier, 0); ok {
		return IntegerPayload(data, na, options...)
	}

	if data, na, ok := ApplyTypeWithNA[T, float64](inData, inNA, applier, math.NaN()); ok {
		return FloatPayload(data, na, options...)
	}

	if data, na, ok := ApplyTypeWithNA[T, complex128](inData, inNA, applier, cmplx.NaN()); ok {
		return ComplexPayload(data, na, options...)
	}

	if data, na, ok := ApplyTypeWithNA[T, string](inData, inNA, applier, ""); ok {
		return StringPayload(data, na, options...)
	}

	if data, na, ok := ApplyTypeWithNA[T, time.Time](inData, inNA, applier, time.Time{}); ok {
		return TimePayload(data, na, options...)
	}

	if data, na, ok := ApplyTypeWithNA[T, Vector](inData, inNA, applier, nil); ok {
		for i, isNA := range na {
			if isNA {
				data[i] = nil
			}
		}
		return VectorPayload(data, options...)
	}

	if data, na, ok := ApplyTypeWithNA[T, any](inData, inNA, applier, nil); ok {
		return AnyPayload(data, na, options...)
	}

	return NAPayload(len(inData))
}

func ApplyTypeWithNA[T, S any](inData []T, inNA []bool, applier any, naDef S) ([]S, []bool, bool) {
	if applyFunc, ok := applier.(func(int, T, bool) (S, bool)); ok {
		data, na := ApplyByFunc[T, S](inData, inNA, applyFunc, naDef)
		return data, na, true
	}

	if applyFunc, ok := applier.(func(T, bool) (S, bool)); ok {
		data, na := ApplyByCompactFunc[T, S](inData, inNA, applyFunc, naDef)
		return data, na, true
	}

	if applyFunc, ok := applier.(func(T) S); ok {
		data, na := ApplyByBriefFunc[T, S](inData, inNA, applyFunc, naDef)
		return data, na, true
	}

	return nil, nil, false
}

func ApplyWithoutNA[T any](inData []T, applier any, options []option.Option) Payload {
	if data, ok := ApplyTypeWithoutNA[T, Vector](inData, applier); ok {
		return VectorPayload(data, options...)
	}

	return NAPayload(len(inData))
}

func ApplyTypeWithoutNA[T, S any](inData []T, applier any) ([]S, bool) {
	if applyFunc, ok := applier.(func(int, T) S); ok {
		data := ApplyByNoNAFunc[T, S](inData, applyFunc)
		return data, true
	}

	if applyFunc, ok := applier.(func(T) S); ok {
		data := ApplyByNoNACompactFunc[T, S](inData, applyFunc)
		return data, true
	}

	return nil, false
}

func ApplyByNoNAFunc[T, S any](inData []T, applyFunc func(int, T) S) []S {
	length := len(inData)

	data := make([]S, length)

	for i := 0; i < length; i++ {
		data[i] = applyFunc(i+1, inData[i])
	}

	return data
}

func ApplyByNoNACompactFunc[T, S any](inData []T, applyFunc func(T) S) []S {
	length := len(inData)

	data := make([]S, length)

	for i := 0; i < length; i++ {
		data[i] = applyFunc(inData[i])
	}

	return data
}

func ApplyByFunc[T, S any](inData []T, inNA []bool,
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

func ApplyByCompactFunc[T, S any](inData []T, inNA []bool,
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

func ApplyByBriefFunc[T, S any](inData []T, inNA []bool,
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

func ApplyToWithNA[T any](indices []int, inData []T, inNA []bool, applier any, naDef T) ([]T, []bool) {
	var data []T = nil
	var na []bool = nil

	switch fn := applier.(type) {
	case func(int, T, bool) (T, bool):
		data, na = ApplyToWithNAByFunc(indices, inData, inNA, fn, naDef)
	case func(T, bool) (T, bool):
		data, na = ApplyToByWithNACompactFunc(indices, inData, inNA, fn, naDef)
	case func(T) T:
		data, na = ApplyToByBriefFunc(indices, inData, inNA, fn)
	case T:
		data, na = ApplyToWithNAByValue(indices, inData, inNA, fn)
	}

	return data, na
}

func ApplyToWithNAByFunc[T any](indices []int, inData []T, inNA []bool,
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

func ApplyToByWithNACompactFunc[T any](indices []int, inData []T, inNA []bool,
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

func ApplyToWithNAByValue[T any](indices []int, inData []T, inNA []bool, val T) ([]T, []bool) {
	data := make([]T, len(inData))
	na := make([]bool, len(inData))

	copy(data, inData)
	copy(na, inNA)

	for _, idx := range indices {
		idx = idx - 1
		data[idx] = val
		na[idx] = false
	}

	return data, na
}

func ApplyToWithoutNA[T any](indices []int, inData []T, applier any) []T {
	var data []T = nil

	switch fn := applier.(type) {
	case func(int, T) T:
		data = ApplyToWithoutNAByFunc(indices, inData, fn)
	case func(T) T:
		data = ApplyToByWithoutNACompactFunc(indices, inData, fn)
	case T:
		data = ApplyToWithoutNAByValue(indices, inData, fn)
	}

	return data
}

func ApplyToWithoutNAByFunc[T any](indices []int, inData []T, applyFunc func(int, T) T) []T {
	length := len(inData)
	data := make([]T, length)
	copy(data, inData)

	for _, idx := range indices {
		idx = idx - 1
		dataVal := applyFunc(idx+1, inData[idx])
		data[idx] = dataVal
	}

	return data
}

func ApplyToByWithoutNACompactFunc[T any](indices []int, inData []T, applyFunc func(T) T) []T {
	length := len(inData)
	data := make([]T, length)
	copy(data, inData)

	for _, idx := range indices {
		idx = idx - 1
		dataVal := applyFunc(inData[idx])
		data[idx] = dataVal
	}

	return data
}

func ApplyToWithoutNAByValue[T any](indices []int, inData []T, val T) []T {
	data := make([]T, len(inData))
	na := make([]bool, len(inData))

	copy(data, inData)

	for _, idx := range indices {
		idx = idx - 1
		data[idx] = val
		na[idx] = false
	}

	return data
}

func TraverseWithNA[T any](inData []T, inNA []bool, traverser any) {
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

func TraverseWithoutNA[T any](inData []T, traverser any) {
	length := len(inData)

	switch fn := traverser.(type) {
	case func(int, T):
		for i := 0; i < length; i++ {
			fn(i, inData[i])
		}
	case func(T):
		for i := 0; i < length; i++ {
			fn(inData[i])
		}
	}
}

func ApplyToByBriefFunc[T any](indices []int, inData []T, inNA []bool, applyFunc func(T) T) ([]T, []bool) {
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

func GroupsForData[T comparable](srcData []T, srcNA []bool) ([][]int, []any) {
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

func GroupsForDataWithHash[T any, S comparable](srcData []T, srcNA []bool, fnHash func(T) S) ([][]int, []any) {
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

func SupportsSummarizer[T any](summarizer any) bool {
	if _, ok := summarizer.(func(int, T, T, bool) (T, bool)); ok {
		return true
	}

	return false
}

func Summarize[T any](inData []T, inNA []bool, summarizer any, valInit T, naDef T) (T, bool) {
	fn, ok := summarizer.(func(int, T, T, bool) (T, bool))
	if ok {
		return SummarizeByFunc(inData, inNA, fn, valInit, naDef)
	}

	return naDef, true
}

func SummarizeByFunc[T any](inData []T, inNA []bool, fn func(int, T, T, bool) (T, bool), valInit T, naDef T) (T, bool) {
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

func Eq[T comparable](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
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

func EqFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool), eqFn func(T, T) bool) []bool {
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

func Neq[T comparable](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
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

func NeqFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool), eqFn func(T, T) bool) []bool {
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

func Gt[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
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

func GtFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
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

func Lt[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
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

func LtFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
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

func Gte[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
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

func GteFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
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

func Lte[T constraints.Ordered](val any, inData []T, inNA []bool, convertor func(any) (T, bool)) []bool {
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

func LteFn[T any](val any, inData []T, inNA []bool, convertor func(any) (T, bool),
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

func Find[T comparable](needle any, inData []T, inNA []bool, convertor func(any) (T, bool)) int {
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

func FindFn[T any](needle any, inData []T, inNA []bool, convertor func(any) (T, bool), eqFn func(T, T) bool) int {
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

func FindAll[T comparable](needle any, inData []T, inNA []bool, convertor func(any) (T, bool)) []int {
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

func FindAllFn[T any](needle any, inData []T, inNA []bool, convertor func(any) (T, bool),
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
