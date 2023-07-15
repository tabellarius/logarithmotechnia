package option

type Option interface {
	Key() string
	Value() any
}
