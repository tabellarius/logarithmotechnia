package vector

func Empty() Vector {
	return &empty{}
}

type empty struct {
}

func (v *empty) IsEmpty() bool {
	return true
}

func (v *empty) ByIndex([]int) Vector {
	return Empty()
}

func (v *empty) ByFromTo(int, int) Vector {
	return Empty()
}

func (v *empty) Clone() Vector {
	return Empty()
}

func (v *empty) Length() int {
	return 0
}
