package vector

import "math/cmplx"

func (p *complexPayload) Sum() Payload {
	sum, na := genSum(p.data, p.NA)

	return ComplexPayload([]complex128{sum}, []bool{na}, p.Options()...)
}

func (p *complexPayload) Prod() Payload {
	product, na := genProd(p.data, p.NA)

	return ComplexPayload([]complex128{product}, []bool{na}, p.Options()...)
}

func (p *complexPayload) Mean() Payload {
	var sum, mean complex128
	var na bool

	if p.length != 0 {
		for i := 0; i < p.length; i++ {
			if p.NA[i] {
				mean, na = 0, true
				goto outOfLoop
			}

			sum += p.data[i]
		}

		mean = sum / complex(float64(p.length), 0)
	}
outOfLoop:

	return ComplexPayload([]complex128{mean}, []bool{na}, p.Options()...)
}

func (p *complexPayload) CumSum() Payload {
	data, na := genCumSum(p.data, p.NA, cmplx.NaN())

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) CumProd() Payload {
	data, na := genCumProd(p.data, p.NA, cmplx.NaN())

	return ComplexPayload(data, na, p.Options()...)
}
