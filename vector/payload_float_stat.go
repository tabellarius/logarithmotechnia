package vector

func (p *floatPayload) Sum() Payload {
	sum, na := genSum(p.data, p.na)

	return FloatPayload([]float64{sum}, []bool{na}, p.Options()...)
}

func (p *floatPayload) Max() Payload {
	max, na := genMax(p.data, p.na)

	return FloatPayload([]float64{max}, []bool{na}, p.Options()...)
}

func (p *floatPayload) Min() Payload {
	min, na := genMin(p.data, p.na)

	return FloatPayload([]float64{min}, []bool{na}, p.Options()...)
}

func (p *floatPayload) Mean() Payload {
	mean, na := genMean(p.data, p.na)

	return FloatPayload([]float64{mean}, []bool{na})
}
