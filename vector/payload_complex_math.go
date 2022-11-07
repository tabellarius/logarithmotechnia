package vector

func (p *complexPayload) Sum() Payload {
	sum, na := genSum(p.data, p.na)

	return ComplexPayload([]complex128{sum}, []bool{na}, p.Options()...)
}
