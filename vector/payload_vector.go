package vector

type vectorPayload struct {
	length int
	data   []Vector
	DefNAble
}

func (p *vectorPayload) Type() string {
	return "vector"
}

func (p *vectorPayload) Len() int {
	return p.length
}

func (p *vectorPayload) ByIndices(indices []int) Payload {
	data, _ := byIndicesWithNA(indices, p.data, p.na, nil)

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

func (p *vectorPayload) Anies() ([]any, []bool) {
	if p.length == 0 {
		return []any{}, []bool{}
	}

	data := make([]any, p.length)
	for i := 0; i < p.length; i++ {
		data[i] = p.data[i]
	}

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *vectorPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicherWithoutNA[Vector](whicher)
}

func (p *vectorPayload) Which(whicher any) []bool {
	return whichWithoutNA[Vector](p.data, whicher)
}

func (p *vectorPayload) Apply(applier any) Payload {
	return apply(p.data, p.na, applier, p.Options())
}

func (p *vectorPayload) ApplyTo(indices []int, applier any) Payload {
	data, _ := applyTo(indices, p.data, p.na, applier, nil)

	if data == nil {
		return NAPayload(p.length)
	}

	return VectorPayload(data, p.Options()...)
}

func (p *vectorPayload) Traverse(traverser any) {
	traverse(p.data, p.na, traverser)
}

func (p *vectorPayload) String() string {
	if p.length == 0 {
		return "<>"
	}

	str := "<" + p.data[0].String()
	for _, v := range p.data[1:] {
		str += ", " + v.String()
	}
	str += ">"

	return str
}

func (p *vectorPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []Vector

	if same, ok := payload.(*vectorPayload); ok {
		srcData = same.data
	} else if intable, ok := payload.(Vectorable); ok {
		srcData = intable.Vectors()
	} else {
		return p
	}

	dstData := make([]Vector, p.length)

	for i := 0; i < p.length; i++ {
		if p.data[i] == nil {
			dstData[i] = srcData[i]
		} else {
			dstData[i] = p.data[i]
		}
	}

	return VectorPayload(dstData, p.Options()...)
}

func VectorPayload(data []Vector, options ...Option) Payload {
	length := len(data)

	vecNA := make([]bool, length)
	for i, v := range data {
		if v == nil {
			vecNA[i] = true
		}
	}

	vecData := make([]Vector, length)
	copy(vecData, data)

	return &vectorPayload{
		length: length,
		data:   vecData,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}
}

func VectorVector(data []Vector, options ...Option) Vector {
	return New(VectorPayload(data, options...), options...)
}
