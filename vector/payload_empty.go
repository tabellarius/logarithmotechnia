package vector

type emptyPayload struct {
}

func (e *emptyPayload) ByIndices(ints []int) Payload {
	return e
}

func (e *emptyPayload) SupportsFilter(selector interface{}) bool {
	return false
}

func (e *emptyPayload) Filter(selector interface{}) []int {
	return []int{}
}

func (e *emptyPayload) Length() int {
	return 0
}

func (e *emptyPayload) NA() []bool {
	return []bool{}
}

func NewEmptyPayload() Payload {
	return &emptyPayload{}
}
