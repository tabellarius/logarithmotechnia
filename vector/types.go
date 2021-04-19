package vector

type ArrBool []bool

func (b ArrBool) Any() bool {
	for _, val := range b {
		if val {
			return true
		}
	}

	return false
}

func (b ArrBool) All() bool {
	for _, val := range b {
		if val {
			return false
		}
	}

	return true
}

type IntNA struct {
	Val int
	NA  bool
}

type FloatNA struct {
	Val float64
	NA  bool
}

type BoolNA struct {
	Val bool
	NA  bool
}

type StringNA struct {
	Val string
	NA  bool
}
