package vector

import (
	"logarithmotechnia/internal/util"
	"reflect"
	"testing"
)

func TestVectorVector(t *testing.T) {
	testData := []struct {
		name    string
		in      []Vector
		vecType []string
		vecLen  []int
	}{
		{
			name: "regular",
			in: []Vector{
				Integer([]int{1, 2, 3}),
				String([]string{"a", "b", "c", "d", "e"}),
				NA(10),
			},
			vecType: []string{"integer", "string", "na"},
			vecLen:  []int{3, 5, 10},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			for i := 0; i < len(data.in); i++ {
				if data.in[i].Type() != data.vecType[i] {
					t.Errorf("Element type (%v) is not equal to expected %v", data.in[i].Type(), data.vecType[i])
				}

				if data.in[i].Len() != data.vecLen[i] {
					t.Errorf("Element length (%v) is not equal to expected %v", data.in[i].Len(), data.vecLen[i])
				}
			}
		})
	}
}

func TestVectorPayload_Type(t *testing.T) {
	vec := VectorPayload([]Vector{})

	if vec.Type() != "vector" {
		t.Error("Vector payload's type is not \"vector\"")
	}
}

func TestVectorPayload_Len(t *testing.T) {
	testData := []struct {
		name   string
		vec    Payload
		length int
	}{
		{
			name: "regular",
			vec: VectorPayload([]Vector{
				Integer([]int{1, 2, 3}),
				String([]string{"a", "b", "c", "d", "e"}),
				NA(10),
			}),
			length: 3,
		},
		{
			name:   "empty",
			vec:    VectorPayload([]Vector{}),
			length: 0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.vec.Len() != data.length {
				t.Errorf("Vector payload's length (%v) is not equal to expected %v", data.vec.Len(), data.length)
			}
		})
	}
}

func TestVectorPayload_ByIndices(t *testing.T) {
	srcPayload := VectorPayload([]Vector{
		Integer([]int{1, 2, 3, 4, 5}),
		String([]string{"a", "b", "c"}),
		Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		Integer([]int{6, 7, 8, 9}),
		String([]string{"d"}),
	})

	testData := []struct {
		name    string
		in      Payload
		vecType []string
		vecLen  []int
	}{
		{
			name:    "regular",
			in:      srcPayload.ByIndices([]int{0, 2, 3, 4, 1}),
			vecType: []string{"", "string", "float", "integer", "integer"},
			vecLen:  []int{0, 3, 10, 4, 5},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vectors := data.in.Data()
			for i, vec := range vectors {
				if data.vecType[i] == "" && vec != nil || vec == nil && data.vecType[i] != "" {
					t.Errorf("%v: Element is not nil", i)
				}

				if vec == nil {
					continue
				}

				if vec.(Vector).Type() != data.vecType[i] {
					t.Errorf("%v: Element type (%v) is not equal to expected %v", i, vec.(Vector).Type(), data.vecType[i])
				}

				if vec.(Vector).Len() != data.vecLen[i] {
					t.Errorf("%v: Element length (%v) is not equal to expected %v", i, vec.(Vector).Len(), data.vecLen[i])
				}
			}

		})
	}
}

func TestVectorPayload_Append(t *testing.T) {
	testData := []struct {
		name   string
		a      Payload
		b      Payload
		result Payload
	}{
		{
			name: "regular",
			a: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
			b: VectorPayload([]Vector{
				Integer([]int{6, 7, 8, 9}),
				String([]string{"d"}),
			}),
			result: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				Integer([]int{6, 7, 8, 9}),
				String([]string{"d"}),
			}),
		},
		{
			name: "b empty",
			a: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
			b: VectorPayload([]Vector{}),
			result: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
		},
		{
			name: "a empty",
			a:    VectorPayload([]Vector{}),
			b: VectorPayload([]Vector{
				Integer([]int{6, 7, 8, 9}),
				String([]string{"d"}),
			}),
			result: VectorPayload([]Vector{
				Integer([]int{6, 7, 8, 9}),
				String([]string{"d"}),
			}),
		},
		{
			name:   "both empty",
			a:      VectorPayload([]Vector{}),
			b:      VectorPayload([]Vector{}),
			result: VectorPayload([]Vector{}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			resVectors := data.a.Append(data.b).Data()
			dataVectors := data.result.Data()
			for i := 0; i < len(dataVectors); i++ {
				if !CompareVectorsForTest(resVectors[i].(Vector), dataVectors[i].(Vector)) {
					t.Errorf("Result (%v) is not equal to expected %v", resVectors, dataVectors)
					break
				}
			}
		})
	}
}

func TestVectorPayload_Adjust(t *testing.T) {
	inPayload := VectorPayload([]Vector{
		Integer([]int{1, 2, 3, 4, 5}),
		String([]string{"a", "b", "c"}),
		Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		Integer([]int{6, 7, 8, 9}),
		String([]string{"d"}),
	})

	testData := []struct {
		name string
		in   Payload
		size int
		out  Payload
	}{
		{
			name: "size less",
			in:   inPayload,
			size: 3,
			out: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
		},
		{
			name: "size more",
			in:   inPayload,
			size: 11,
			out: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				Integer([]int{6, 7, 8, 9}),
				String([]string{"d"}),
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				Integer([]int{6, 7, 8, 9}),
				String([]string{"d"}),
				Integer([]int{1, 2, 3, 4, 5}),
			}),
		},
		{
			name: "size zero",
			in:   inPayload,
			size: 0,
			out:  VectorPayload([]Vector{}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			CompareVectorArrs(inPayload.(*vectorPayload).data, data.out.(*vectorPayload).data)
		})
	}
}

func TestVectorPayload_Options(t *testing.T) {
	testData := []struct {
		name string
		in   Payload
		out  []Option
	}{
		{
			name: "no options",
			in: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
			out: []Option{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.in.Options(), data.out) {
				t.Errorf("Options (%v) are not equal to expected %v", data.in.Options(), data.out)
			}
		})
	}
}

func TestVectorPayload_SetOption(t *testing.T) {
	if result := VectorPayload([]Vector{}).SetOption("", nil); result != false {
		t.Errorf("SetOption should return false")
	}
}

func TestVectorPayload_Pick(t *testing.T) {
	inPayload := VectorPayload([]Vector{
		Integer([]int{1, 2, 3, 4, 5}),
		String([]string{"a", "b", "c"}),
		Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
	})

	testData := []struct {
		name  string
		in    Payload
		vec   Vector
		index int
	}{
		{
			name:  "index 1",
			in:    inPayload,
			vec:   Integer([]int{1, 2, 3, 4, 5}),
			index: 1,
		},
		{
			name:  "index 3",
			in:    inPayload,
			vec:   Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			index: 3,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := data.in.Pick(data.index).(Vector)
			if !CompareVectorsForTest(vec, data.vec) {
				t.Errorf("Pick (%v) is not equal to expected %v", vec, data.vec)
			}
		})
	}
}

func TestVectorPayload_Data(t *testing.T) {
	testData := []struct {
		name string
		in   Payload
		vecs []Vector
	}{
		{
			name: "regular",
			in: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
			vecs: []Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			},
		},
		{
			name: "empty",
			in:   VectorPayload([]Vector{}),
			vecs: []Vector{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vecs := util.ArrayFromAnyTo[Vector](data.in.Data())
			if !CompareVectorArrs(vecs, data.vecs) {
				t.Errorf("Data (%v) is not equal to expected %v", vecs, data.vecs)
			}
		})
	}
}

func TestVectorPayload_Vectors(t *testing.T) {
	testData := []struct {
		name string
		in   Payload
		vecs []Vector
	}{
		{
			name: "regular",
			in: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
			vecs: []Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			},
		},
		{
			name: "empty",
			in:   VectorPayload([]Vector{}),
			vecs: []Vector{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vecs := data.in.(*vectorPayload).Vectors()
			if !CompareVectorArrs(vecs, data.vecs) {
				t.Errorf("Vectors (%v) is not equal to expected %v", vecs, data.vecs)
			}
		})
	}
}

func TestVectorPayload_Anies(t *testing.T) {
	testData := []struct {
		name string
		in   Payload
		vecs []Vector
	}{
		{
			name: "regular",
			in: VectorPayload([]Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			}),
			vecs: []Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			},
		},
		{
			name: "empty",
			in:   VectorPayload([]Vector{}),
			vecs: []Vector{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			anies, _ := data.in.(*vectorPayload).Anies()
			vecs := util.ArrayFromAnyTo[Vector](anies)
			if !CompareVectorArrs(vecs, data.vecs) {
				t.Errorf("Vectors (%v) is not equal to expected %v", vecs, data.vecs)
			}
		})
	}
}

func TestVectorPayload_SupportsWhicher(t *testing.T) {
	inPayload := VectorPayload([]Vector{
		Integer([]int{1, 2, 3, 4, 5}),
	}).(*vectorPayload)

	testData := []struct {
		name     string
		payload  *vectorPayload
		fn       any
		supports bool
	}{
		{
			name:    "full",
			payload: inPayload,
			fn: func(int, Vector) bool {
				return true
			},
			supports: true,
		},
		{
			name:    "compact",
			payload: inPayload,
			fn: func(Vector) bool {
				return true
			},
			supports: true,
		},
		{
			name:    "not supported",
			payload: inPayload,
			fn: func(any) bool {
				return true
			},
			supports: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.payload.SupportsWhicher(data.fn) != data.supports {
				t.Errorf("SupportsWhicher(%v) is not equal to expected %v", data.fn, data.supports)
			}
		})
	}
}

func TestVectorPayload_Which(t *testing.T) {
	inPayload := VectorPayload([]Vector{
		Integer([]int{1, 2, 3, 4, 5}),
		String([]string{"a", "b", "c"}),
		nil,
		Integer([]int{6, 7, 8, 9}),
		String([]string{"d"}),
	}).(*vectorPayload)

	testData := []struct {
		name    string
		payload *vectorPayload
		fn      any
		which   []bool
	}{
		{
			name:    "full",
			payload: inPayload,
			fn: func(i int, v Vector) bool {
				if i == 1 || v == nil {
					return true
				}
				return false
			},
			which: []bool{true, false, true, false, false},
		},
		{
			name:    "compact",
			payload: inPayload,
			fn: func(v Vector) bool {
				if v == nil {
					return true
				}
				return false
			},
			which: []bool{false, false, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			which := data.payload.Which(data.fn)
			if !reflect.DeepEqual(which, data.which) {
				t.Errorf("Which(%v) is not equal to expected %v", which, data.which)
			}
		})
	}

}

func TestVectorPayload_Apply(t *testing.T) {
	inPayload := VectorPayload([]Vector{
		Integer([]int{1, 2, 3, 4, 5}),
		String([]string{"a", "b", "c"}),
		nil,
		Integer([]int{6, 7, 8, 9}),
		String([]string{"d"}),
	}).(*vectorPayload)

	outVectors := []Vector{
		String([]string{"integer"}),
		String([]string{"string"}),
		String([]string{""}),
		String([]string{"integer"}),
		String([]string{"string"}),
	}

	testData := []struct {
		name    string
		in      *vectorPayload
		applier any
		out     []Vector
		na      bool
	}{
		{
			name: "full",
			in:   inPayload,
			applier: func(i int, v Vector) Vector {
				if v == nil {
					return String([]string{""})
				}

				return String([]string{v.Type()})
			},
			out: outVectors,
			na:  false,
		},
		{
			name: "compact",
			in:   inPayload,
			applier: func(v Vector) Vector {
				if v == nil {
					return String([]string{""})
				}

				return String([]string{v.Type()})
			},
			out: outVectors,
			na:  false,
		},
		{
			name: "not supported",
			in:   inPayload,
			applier: func(v int) Vector {
				return nil
			},
			na: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.na {
				out := data.in.Apply(data.applier)
				if out.Type() != PayloadTypeNA {
					t.Error("Payload is not NA")
				}

				if out.Len() != data.in.Len() {
					t.Errorf("Payload length (%v) is not equal to expected %v", out.Len(), data.in.Len())
				}
			} else {
				out := data.in.Apply(data.applier).(*vectorPayload).data
				if !CompareVectorArrs(out, data.out) {
					t.Errorf("Apply(%v) is not equal to expected %v", out, data.out)
				}
			}
		})
	}
}

func TestVectorPayload_ApplyTo(t *testing.T) {
	inPayload := VectorPayload([]Vector{
		Integer([]int{1, 2, 3, 4, 5}),
		String([]string{"a", "b", "c"}),
		Float([]float64{1.1, 2.2, 3.3, 4.4, 5.5}),
		Integer([]int{6, 7, 8, 9}),
		String([]string{"d"}),
	}).(*vectorPayload)

	testData := []struct {
		name    string
		in      *vectorPayload
		indices []int
		applier any
		out     []Vector
	}{
		{
			name:    "full",
			in:      inPayload,
			indices: []int{1, 3, 5},
			applier: func(i int, v Vector) Vector {
				if i == 1 || i == 5 {
					return nil
				}

				return v
			},
			out: []Vector{
				nil,
				String([]string{"a", "b", "c"}),
				Float([]float64{1.1, 2.2, 3.3, 4.4, 5.5}),
				Integer([]int{6, 7, 8, 9}),
				nil,
			},
		},
		{
			name:    "compact",
			in:      inPayload,
			indices: []int{2, 3, 4},
			applier: func(v Vector) Vector {
				if v.Type() == PayloadTypeInteger {
					return nil
				}
				return v
			},
			out: []Vector{
				Integer([]int{1, 2, 3, 4, 5}),
				String([]string{"a", "b", "c"}),
				Float([]float64{1.1, 2.2, 3.3, 4.4, 5.5}),
				Integer([]int{6, 7, 8, 9}),
				nil,
			},
		},
		{
			name:    "not supported",
			in:      inPayload,
			indices: []int{1, 2, 3},
			applier: func(v int) Vector {
				return nil
			},
			out: []Vector{
				nil,
				nil,
				nil,
				Integer([]int{6, 7, 8, 9}),
				String([]string{"d"}),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := data.in.ApplyTo(data.indices, data.applier).(*vectorPayload).data
			if !CompareVectorArrs(out, data.out) {
				t.Errorf("Apply(%v) is not equal to expected %v", out, data.out)
			}
		})
	}
}

func TestVectorPayload_Traverse(t *testing.T) {

}

func TestVectorPayload_String(t *testing.T) {

}

func TestVectorPayload_Coalesce(t *testing.T) {

}
