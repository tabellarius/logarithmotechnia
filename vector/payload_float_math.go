package vector

func (p *floatPayload) Sum() Payload {
	sum, na := genSum(p.data, p.na)

	return FloatPayload([]float64{sum}, []bool{na}, p.Options()...)
}
