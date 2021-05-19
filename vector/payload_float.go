package vector

import (
	"github.com/shopspring/decimal"
	"time"
)

type float struct {
	length int
	data   []float64
	DefNAble
}

func (p *float) Len() int {
	return p.length
}

func (p *float) Integers() ([]int, []bool) {
	panic("implement me")
}

func (p *float) Floats() ([]float64, []bool) {
	panic("implement me")
}

func (p *float) Booleans() ([]bool, []bool) {
	panic("implement me")
}

func (p *float) Strings() ([]string, []bool) {
	panic("implement me")
}

func (p *float) Complexes() ([]complex128, []bool) {
	panic("implement me")
}

func (p *float) Decimals() ([]decimal.Decimal, []bool) {
	panic("implement me")
}

func (p *float) Times() ([]time.Time, []bool) {
	panic("implement me")
}

func (p *float) ByIndices(indices []int) Payload {
	panic("implement me")
}

func (p *float) SupportsSelector(filter interface{}) bool {
	panic("implement me")
}

func (p *float) Select(selector interface{}) []bool {
	panic("implement me")
}

func (p *float) StrForElem(idx int) string {
	panic("implement me")
}
