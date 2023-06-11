package vector

import "math"

func (p *floatPayload) Sum() Payload {
	sum, na := genSum(p.data, p.NA)

	return FloatPayload([]float64{sum}, []bool{na}, p.Options()...)
}

func (p *floatPayload) Prod() Payload {
	product, na := genProd(p.data, p.NA)

	return FloatPayload([]float64{product}, []bool{na}, p.Options()...)
}

func (p *floatPayload) Max() Payload {
	max, na := genMax(p.data, p.NA)

	return FloatPayload([]float64{max}, []bool{na}, p.Options()...)
}

func (p *floatPayload) Min() Payload {
	min, na := genMin(p.data, p.NA)

	return FloatPayload([]float64{min}, []bool{na}, p.Options()...)
}

func (p *floatPayload) Mean() Payload {
	mean, na := genMean(p.data, p.NA)

	return FloatPayload([]float64{mean}, []bool{na})
}

func (p *floatPayload) Median() Payload {
	median, na := genMedian(p.data, p.DefNAble, p.sortedIndices)

	return FloatPayload([]float64{median}, []bool{na})
}

func (p *floatPayload) CumSum() Payload {
	data, na := genCumSum(p.data, p.NA, math.NaN())

	return FloatPayload(data, na, p.Options()...)
}

func (p *floatPayload) CumProd() Payload {
	data, na := genCumProd(p.data, p.NA, math.NaN())

	return FloatPayload(data, na, p.Options()...)
}

func (p *floatPayload) CumMax() Payload {
	data, na := genCumMax(p.data, p.NA, math.NaN())

	return FloatPayload(data, na, p.Options()...)
}

func (p *floatPayload) CumMin() Payload {
	data, na := genCumMin(p.data, p.NA, math.NaN())

	return FloatPayload(data, na, p.Options()...)
}
