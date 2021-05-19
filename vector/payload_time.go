package vector

import (
	"github.com/shopspring/decimal"
	"time"
)

type timePayload struct {
	length int
	data   []time.Time
	DefNAble
}

func (p *timePayload) Len() int {
	return p.length
}

func (p *timePayload) Integers() ([]int, []bool) {
	panic("implement me")
}

func (p *timePayload) Floats() ([]float64, []bool) {
	panic("implement me")
}

func (p *timePayload) Booleans() ([]bool, []bool) {
	panic("implement me")
}

func (p *timePayload) Strings() ([]string, []bool) {
	panic("implement me")
}

func (p *timePayload) Complexes() ([]complex128, []bool) {
	panic("implement me")
}

func (p *timePayload) Decimals() ([]decimal.Decimal, []bool) {
	panic("implement me")
}

func (p *timePayload) Times() ([]time.Time, []bool) {
	panic("implement me")
}

func (p *timePayload) ByIndices(indices []int) Payload {
	panic("implement me")
}

func (p *timePayload) SupportsSelector(filter interface{}) bool {
	panic("implement me")
}

func (p *timePayload) Select(selector interface{}) []bool {
	panic("implement me")
}

func (p *timePayload) StrForElem(idx int) string {
	panic("implement me")
}
