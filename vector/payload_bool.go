package vector

type boolean struct {
	length int
	data   []bool
	DefNAble
}

func (p *boolean) Len() int {
	return p.length
}

func (p *boolean) ByIndices(indices []int) Payload {
	panic("implement me")
}

func (p *boolean) SupportsSelector(filter interface{}) bool {
	panic("implement me")
}

func (p *boolean) Select(selector interface{}) []bool {
	panic("implement me")
}

func (p *boolean) Integers() ([]int, []bool) {
	panic("implement me")
}

func (p *boolean) Floats() ([]float64, []bool) {
	panic("implement me")
}

func (p *boolean) Booleans() ([]bool, []bool) {
	panic("implement me")
}

func (p *boolean) Strings() ([]string, []bool) {
	panic("implement me")
}

func (p *boolean) Complexes() ([]complex128, []bool) {
	panic("implement me")
}

func (p *boolean) StrForElem(idx int) string {
	panic("implement me")
}
