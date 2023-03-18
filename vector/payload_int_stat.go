package vector

func (p *integerPayload) Sum() Payload {
	sum, na := genSum(p.data, p.na)

	return IntegerPayload([]int{sum}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Max() Payload {
	max, na := genMax(p.data, p.na)

	return IntegerPayload([]int{max}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Min() Payload {
	min, na := genMin(p.data, p.na)

	return IntegerPayload([]int{min}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Mean() Payload {
	mean, na := genMean(p.data, p.na)

	return FloatPayload([]float64{mean}, []bool{na})
}

func (p *integerPayload) Median() Payload {
	median, na := genMedian(p.data, p.DefNAble, p.sortedIndices)

	return IntegerPayload([]int{median}, []bool{na})
}
