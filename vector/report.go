package vector

func NewReport() Report {
	return Report{
		Errors:   make([]string, 0),
		Warnings: make([]string, 0),
	}
}

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
