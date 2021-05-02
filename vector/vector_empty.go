package vector

func Empty() EmptyVector {
	return &empty{
		report: Report{},
	}
}

type EmptyVector interface {
	Vector
	Reporter
}

type empty struct {
	report Report
}

func (v *empty) ByBool([]bool) Vector {
	return Empty()
}

func (v *empty) IsEmpty() bool {
	return true
}

func (v *empty) ByIndices([]int) Vector {
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

func (v *empty) Report() Report {
	return v.report
}
