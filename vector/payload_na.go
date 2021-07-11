package vector

import (
	"math"
	"math/cmplx"
	"time"
)

type naPayload struct {
	length int
}

/* Payload interface */

func (p *naPayload) Type() string {
	return "na"
}

func (p *naPayload) Len() int {
	return p.length
}

func (p *naPayload) ByIndices(indices []int) Payload {
	length := 0

	for _, idx := range indices {
		if idx >= 1 && idx <= p.length {
			length++
		}
	}

	return &naPayload{length: length}
}

func (p *naPayload) StrForElem(int) string {
	return "NA"
}

/* NAble interface */

func (p *naPayload) IsNA() []bool {
	isna := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		isna[i] = true
	}

	return isna
}

func (p *naPayload) NotNA() []bool {
	isna := make([]bool, p.length)

	return isna
}

func (p *naPayload) HasNA() bool {
	return true
}

func (p *naPayload) WithNA() []int {
	naIndices := make([]int, p.length)

	for i := 0; i < p.length; i++ {
		naIndices[i] = i + 1
	}

	return naIndices
}

func (p *naPayload) WithoutNA() []int {
	return []int{}
}

/* Convertors */

func (p *naPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)

	return data, p.naArray()
}

func (p *naPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)

	for i := 0; i < p.length; i++ {
		data[i] = math.NaN()
	}

	return data, p.naArray()
}

func (p *naPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)

	for i := 0; i < p.length; i++ {
		data[i] = cmplx.NaN()
	}

	return data, p.naArray()
}

func (p *naPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)

	return data, p.naArray()
}

func (p *naPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)

	return data, p.naArray()
}

func (p *naPayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)

	return data, p.naArray()
}

func (p *naPayload) naArray() []bool {
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		na[i] = true
	}

	return na
}

func NAPayload(length int) Payload {
	if length < 0 {
		length = 0
	}

	payload := &naPayload{
		length: length,
	}

	return payload
}

func NA(length int) Vector {
	return New(NAPayload(length), Config{})
}
