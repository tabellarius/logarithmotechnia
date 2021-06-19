package vector

type emptyPayload struct {
}

func (e *emptyPayload) Type() string {
	return "empty"
}

func (e *emptyPayload) Len() int {
	return 0
}

func (e *emptyPayload) ByIndices([]int) Payload {
	return EmptyPayload()
}

func (e *emptyPayload) NAPayload() Payload {
	return EmptyPayload()
}

func EmptyPayload() Payload {
	return &emptyPayload{}
}

func (e *emptyPayload) StrForElem(int) string {
	return ""
}
