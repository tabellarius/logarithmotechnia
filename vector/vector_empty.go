package vector

func Empty() Vector {
	return &empty{}
}

type empty struct {
}

func (v *empty) Report() Report {
	return Report{}
}

func (v *empty) Clone() Vector {
	return Empty()
}

func (v *empty) Length() int {
	return 0
}

func (v *empty) Names() []string {
	return nil
}

func (v *empty) NamesMap() map[int]string {
	return nil
}

func (v *empty) SetName(int, string) Vector {
	return v
}

func (v *empty) SetNames([]string) Vector {
	return v
}

func (v *empty) SetNamesMap(map[int]string) Vector {
	return v
}

func (v *empty) IfNameFor(int) bool {
	return false
}

func (v *empty) NA() []bool {
	return nil
}

func (v *empty) NAMap() map[int]bool {
	return nil
}

func (v *empty) SetNA(na []bool) Vector {
	return v
}

func (v *empty) SetNAMap(na map[int]bool) Vector {
	return v
}
