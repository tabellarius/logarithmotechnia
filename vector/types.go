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

type Int struct {
	Val int
	NA  bool
}

type Float struct {
	Val float64
	NA  bool
}

type Bool struct {
	Val bool
	NA  bool
}

type String struct {
	Val string
	NA  bool
}
