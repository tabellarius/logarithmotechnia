package vector

import (
	"time"
)

func (p *timePayload) Max() Payload {
	if p.length == 0 {
		return TimePayload([]time.Time{{}}, []bool{true}, p.Options()...)
	}

	max := p.data[0]
	for i := 1; i < p.length; i++ {
		if p.NA[i] {
			return TimePayload([]time.Time{{}}, []bool{true}, p.Options()...)
		}

		if p.data[i].After(max) {
			max = p.data[i]
		}
	}

	return TimePayload([]time.Time{max}, []bool{false}, p.Options()...)
}

func (p *timePayload) Min() Payload {
	if p.length == 0 {
		return TimePayload([]time.Time{{}}, []bool{true}, p.Options()...)
	}

	min := p.data[0]
	for i := 1; i < p.length; i++ {
		if p.NA[i] {
			return TimePayload([]time.Time{{}}, []bool{true}, p.Options()...)
		}

		if p.data[i].Before(min) {
			min = p.data[i]
		}
	}

	return TimePayload([]time.Time{min}, []bool{false}, p.Options()...)
}
