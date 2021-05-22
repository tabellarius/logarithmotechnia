package vector

import (
	"github.com/shopspring/decimal"
	"time"
)

type complexPayload struct {
	length int
	data   []complex128
	DefNAble
}

func (p *complexPayload) Len() int {
	return p.length
}

func (p *complexPayload) Integers() ([]int, []bool) {
	panic("implement me")
}

func (p *complexPayload) Floats() ([]float64, []bool) {
	panic("implement me")
}

func (p *complexPayload) Booleans() ([]bool, []bool) {
	panic("implement me")
}

func (p *complexPayload) Strings() ([]string, []bool) {
	panic("implement me")
}

func (p *complexPayload) Complexes() ([]complex128, []bool) {
	panic("implement me")
}

func (p *complexPayload) Decimals() ([]decimal.Decimal, []bool) {
	panic("implement me")
}

func (p *complexPayload) Times() ([]time.Time, []bool) {
	panic("implement me")
}

func (p *complexPayload) ByIndices(indices []int) Payload {
	panic("implement me")
}

func (p *complexPayload) SupportsSelector(filter interface{}) bool {
	panic("implement me")
}

func (p *complexPayload) Select(selector interface{}) []bool {
	panic("implement me")
}

func (p *complexPayload) StrForElem(idx int) string {
	panic("implement me")
}
