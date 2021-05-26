package vector

import (
	"github.com/shopspring/decimal"
)

type decimalPayload struct {
	length int
	data   []decimal.Decimal
	DefNAble
}

func (p *decimalPayload) Len() int {
	return p.length
}

func (p *decimalPayload) Integers() ([]int, []bool) {
	panic("implement me")
}

func (p *decimalPayload) Floats() ([]float64, []bool) {
	panic("implement me")
}

func (p *decimalPayload) Booleans() ([]bool, []bool) {
	panic("implement me")
}

func (p *decimalPayload) Strings() ([]string, []bool) {
	panic("implement me")
}

func (p *decimalPayload) Complexes() ([]complex128, []bool) {
	panic("implement me")
}

func (p *decimalPayload) Decimals() ([]decimal.Decimal, []bool) {
	panic("implement me")
}

func (p *decimalPayload) ByIndices(indices []int) Payload {
	panic("implement me")
}

func (p *decimalPayload) SupportsSelector(filter interface{}) bool {
	panic("implement me")
}

func (p *decimalPayload) Select(selector interface{}) []bool {
	panic("implement me")
}

func (p *decimalPayload) StrForElem(idx int) string {
	panic("implement me")
}
