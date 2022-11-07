package vector

func (p *integerPayload) Sum() Payload {
	sum, na := genSum(p.data, p.na)

	return IntegerPayload([]int{sum}, []bool{na}, p.Options()...)
}
