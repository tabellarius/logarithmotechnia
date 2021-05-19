package vector

import (
	"github.com/shopspring/decimal"
	"time"
)

type str struct {
	length int
	data   []string
	DefNAble
}

func (p *str) Len() int {
	return p.length
}

func (p *str) Integers() ([]int, []bool) {
	panic("implement me")
}

func (p *str) Floats() ([]float64, []bool) {
	panic("implement me")
}

func (p *str) Booleans() ([]bool, []bool) {
	panic("implement me")
}

func (p *str) Strings() ([]string, []bool) {
	panic("implement me")
}

func (p *str) Complexes() ([]complex128, []bool) {
	panic("implement me")
}

func (p *str) Decimals() ([]decimal.Decimal, []bool) {
	panic("implement me")
}

func (p *str) Times() ([]time.Time, []bool) {
	panic("implement me")
}

func (p *str) ByIndices(indices []int) Payload {
	panic("implement me")
}

func (p *str) SupportsSelector(filter interface{}) bool {
	panic("implement me")
}

func (p *str) Select(selector interface{}) []bool {
	panic("implement me")
}

func (p *str) StrForElem(idx int) string {
	panic("implement me")
}
