package vector

type Report struct {
	Errors   []string
	Warnings []string
}

func (r Report) IsClear() bool {
	return len(r.Errors) == 0 && len(r.Warnings) == 0
}

func (r Report) AddError(error string) {
	r.Errors = append(r.Errors, error)
}

func (r Report) AddWarning(warning string) {
	r.Warnings = append(r.Warnings, warning)
}

func NewReport() Report {
	return Report{
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}
}

func (r Report) Copy() Report {
	newRep := Report{
		Errors:   make([]string, len(r.Errors)),
		Warnings: make([]string, len(r.Warnings)),
	}

	copy(newRep.Errors, r.Errors)
	copy(newRep.Warnings, r.Warnings)

	return newRep
}
