package dataframe

import (
	"testing"
	"time"
)

func TestFromStructs(t *testing.T) {
	type Finance struct {
		Money   int
		Account string
	}

	type A struct {
		Title    string
		Status   int
		Kpi      float64
		Cpx      complex128
		IsActive bool
		Date     time.Time
		Misc     Finance
	}

	stArr := []A{
		{
			Title:    "Baron",
			Status:   1,
			Kpi:      1.2,
			Cpx:      1 + 1i,
			IsActive: true,
			Date:     time.Now(),
			Misc:     Finance{1000, "br"},
		},
		{
			Title:    "Earl",
			Status:   3,
			Kpi:      2.2,
			Cpx:      1 + 3i,
			IsActive: false,
			Date:     time.Now().Add(7 * 24 * 60 * time.Minute),
			Misc:     Finance{15000, "br"},
		},
		{
			Title:    "King",
			Status:   5,
			Kpi:      4.45,
			Cpx:      4 + 2i,
			IsActive: true,
			Date:     time.Now().Add(360 * 24 * 60 * time.Minute),
			Misc:     Finance{275000, "br"},
		},
	}

	FromStructs(stArr)
}
