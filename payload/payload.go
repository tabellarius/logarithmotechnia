package payload

import "logarithmotechnia/option"

type Payload interface {
	// Type returns a type of the payload which should be a unique string.
	Type() string
	// Len returns length of the payload.
	Len() int
	// ByIndices returns a new payload which contains elements from the old one with provided indices.
	ByIndices(indices []int) Payload
	// StrForElem return a string representation of the payload's element.
	StrForElem(idx int) string
	// Append appends another payload to the current one based on the type of the current one. If it is impossible
	// to convert an element of another payload, it will be converted to NA-value.
	Append(payload Payload) Payload
	// Adjust adjusts the payload to the provided size either by dropping excessive values or by extending
	// the payload with recycling.
	Adjust(size int) Payload
	// Options returns options of the payload.
	Options() []option.Option
	// SetOption sets a payload's option.
	SetOption(string, any) bool
	// Pick returns a value of a payload element (using interface{} type).
	Pick(idx int) any
	// Data returns all payload elements as an array of values having the interface{} type.
	Data() []any
}
