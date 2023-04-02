package vector

import (
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
					t.Errorf("Element type (%v) is not equal to %v", data.in[i].Type(), data.vecType[i])
				}

				if data.in[i].Len() != data.vecLen[i] {
					t.Errorf("Element length (%v) is not equal to %v", data.in[i].Len(), data.vecLen[i])
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
				t.Errorf("Vector payload's length (%v) is not equal to %v", data.vec.Len(), data.length)
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
					t.Errorf("%v: Element type (%v) is not equal to %v", i, vec.(Vector).Type(), data.vecType[i])
				}

				if vec.(Vector).Len() != data.vecLen[i] {
					t.Errorf("%v: Element length (%v) is not equal to %v", i, vec.(Vector).Len(), data.vecLen[i])
				}
			}

		})
	}
}

func TestVectorPayload_Append(t *testing.T) {

}

func TestVectorPayload_Adjust(t *testing.T) {

}

func TestVectorPayload_Options(t *testing.T) {

}

func TestVectorPayload_SetOption(t *testing.T) {

}

func TestVectorPayload_Pick(t *testing.T) {

}

func TestVectorPayload_Data(t *testing.T) {

}

func TestVectorPayload_Vectors(t *testing.T) {

}

func TestVectorPayload_Anies(t *testing.T) {

}

func TestVectorPayload_SupportsWhicher(t *testing.T) {

}

func TestVectorPayload_Which(t *testing.T) {

}

func TestVectorPayload_Apply(t *testing.T) {

}

func TestVectorPayload_ApplyTo(t *testing.T) {

}

func TestVectorPayload_Traverse(t *testing.T) {

}

func TestVectorPayload_String(t *testing.T) {

}

func TestVectorPayload_Coalesce(t *testing.T) {

}
