package vector

type vectorPayload struct {
	length int
	data   []Vector
}

func (p *vectorPayload) Type() string {
	return "vector"
}

func (p *vectorPayload) Len() int {
	return p.length
}

func (p *vectorPayload) ByIndices(indices []int) Payload {
	data := byIndicesWithoutNA(indices, p.data, nil)

	return VectorPayload(data, p.Options()...)
}

func (p *vectorPayload) StrForElem(idx int) string {
	if p.data[idx-1] == nil {
		return "NA"
	}

	return p.data[idx-1].String()
}

func (p *vectorPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []Vector

	if vectorable, ok := payload.(Vectorable); ok {
		vals = vectorable.Vectors()
	} else {
		vals = make([]Vector, payload.Len())
	}

	newVals := make([]Vector, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)

	return VectorPayload(newVals, p.Options()...)
}

func (p *vectorPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *vectorPayload) adjustToLesserSize(size int) Payload {
	data := adjustToLesserSizeWithoutNA(p.data, size)

	return VectorPayload(data, p.Options()...)
}

func (p *vectorPayload) adjustToBiggerSize(size int) Payload {
	data := adjustToBiggerSizeWithoutNA(p.data, p.length, size)

	return VectorPayload(data, p.Options()...)
}

func (p *vectorPayload) Options() []Option {
	return []Option{}
}

func (p *vectorPayload) SetOption(string, any) bool {
	return false
}

func (p *vectorPayload) Pick(idx int) any {
	return p.data[idx-1]
}

func (p *vectorPayload) Data() []any {
	data := make([]any, p.length)

	for i, v := range p.data {
		data[i] = v
	}

	return data
}

func (p *vectorPayload) Vectors() []Vector {
	vectors := make([]Vector, p.length)
	copy(vectors, p.data)

	return vectors
}

func (p *vectorPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[Vector](whicher)
}

func (p *vectorPayload) Which(whicher any) []bool {
	return whichWithoutNA[Vector](p.data, whicher)
}

func (p *vectorPayload) Apply(applier any) Payload {
	return nil
}

func (p *vectorPayload) ApplyTo(indices []int, applier any) Payload {
	return nil
}

func VectorPayload(data []Vector, options ...Option) Payload {
	return nil
}
