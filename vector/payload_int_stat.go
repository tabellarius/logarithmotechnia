package vector

func (p *integerPayload) Sum() Payload {
	sum, na := genSum(p.data, p.NA)

	return IntegerPayload([]int{sum}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Prod() Payload {
	product, na := genProd(p.data, p.NA)

	return IntegerPayload([]int{product}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Max() Payload {
	max, na := genMax(p.data, p.NA)

	return IntegerPayload([]int{max}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Min() Payload {
	min, na := genMin(p.data, p.NA)

	return IntegerPayload([]int{min}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Mean() Payload {
	mean, na := genMean(p.data, p.NA)

	return FloatPayload([]float64{mean}, []bool{na})
}

func (p *integerPayload) Median() Payload {
	median, na := genMedian(p.data, p.NAble, p.SortedIndicesZeroBased)

	return IntegerPayload([]int{median}, []bool{na})
}

func (p *integerPayload) CumSum() Payload {
	data, na := genCumSum(p.data, p.NA, 0)

	return IntegerPayload(data, na, p.Options()...)
}

func (p *integerPayload) CumProd() Payload {
	data, na := genCumProd(p.data, p.NA, 1)

	return IntegerPayload(data, na, p.Options()...)
}

func (p *integerPayload) CumMax() Payload {
	data, na := genCumMax(p.data, p.NA, 0)

	return IntegerPayload(data, na, p.Options()...)
}

func (p *integerPayload) CumMin() Payload {
	data, na := genCumMin(p.data, p.NA, 0)

	return IntegerPayload(data, na, p.Options()...)
}
