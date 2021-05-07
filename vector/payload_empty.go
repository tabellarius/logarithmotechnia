package vector

type emptyPayload struct {
}

func (e *emptyPayload) ByIndices([]int) Payload {
	return &emptyPayload{}
}

func (e *emptyPayload) SupportsFilter(interface{}) bool {
	return false
}

func (e *emptyPayload) Filter(interface{}) []bool {
	return []bool{}
}

func (e *emptyPayload) Len() int {
	return 0
}

func (e *emptyPayload) NAP() []bool {
	return []bool{}
}

func EmptyPayload() Payload {
	return &emptyPayload{}
}

func (e *emptyPayload) StrForElem(int) string {
	return ""
}
