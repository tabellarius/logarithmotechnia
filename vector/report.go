package vector

type Report struct {
	Errors   []string
	Warnings []string
}

func (r *Report) IsClear() bool {
	return len(r.Errors) == 0 && len(r.Warnings) == 0
}
