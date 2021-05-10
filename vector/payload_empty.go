package vector

type emptyPayload struct {
}

func (e *emptyPayload) SupportsSelector(interface{}) bool {
	return false
}

func (e *emptyPayload) ByIndices([]int) Payload {
	return EmptyPayload()
}

func (e *emptyPayload) Select(interface{}) []bool {
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
