package vector

import (
	"logarithmotechnia/embed"
	"logarithmotechnia/option"
	"time"
)

type TimePrinter struct {
	Format string
}

type timePayload struct {
	length  int
	data    []time.Time
	printer TimePrinter
	embed.NAble
	embed.Arrangeable
}

func (p *timePayload) Type() string {
	return "time"
}

func (p *timePayload) Len() int {
	return p.length
}

func (p *timePayload) Pick(idx int) any {
	return PickValueWithNA(idx, p.data, p.NA, p.length)
}

func (p *timePayload) Data() []any {
	return DataWithNAToInterfaceArray(p.data, p.NA)
}

func (p *timePayload) ByIndices(indices []int) Payload {
	data, na := ByIndicesWithNA(indices, p.data, p.NA, time.Time{})

	return TimePayload(data, na, p.Options()...)
}

func (p *timePayload) SupportsWhicher(whicher any) bool {
	return SupportsWhicherWithNA[time.Time](whicher)
}

func (p *timePayload) Which(whicher any) []bool {
	return WhichWithNA(p.data, p.NA, whicher)
}

func (p *timePayload) Apply(applier any) Payload {
	return ApplyWithNA(p.data, p.NA, applier, p.Options())
}

func (p *timePayload) ApplyTo(indices []int, applier any) Payload {
	data, na := ApplyToWithNA(indices, p.data, p.NA, applier, time.Time{})

	if data == nil {
		return NAPayload(p.length)
	}

	return TimePayload(data, na, p.Options()...)
}

func (p *timePayload) Traverse(traverser any) {
	TraverseWithNA(p.data, p.NA, traverser)
}

func (p *timePayload) SupportsSummarizer(summarizer any) bool {
	return SupportsSummarizer[time.Time](summarizer)
}

func (p *timePayload) Summarize(summarizer any) Payload {
	val, na := Summarize(p.data, p.NA, summarizer, time.Time{}, time.Time{})

	return TimePayload([]time.Time{val}, []bool{na})
}

func (p *timePayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = ""
		} else {
			data[i] = p.StrForElem(i + 1)
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *timePayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *timePayload) Anies() ([]any, []bool) {
	if p.length == 0 {
		return []any{}, []bool{}
	}

	data := make([]any, p.length)
	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = nil
		} else {
			data[i] = p.data[i]
		}
	}

	na := make([]bool, p.length)
	copy(na, p.NA)

	return data, na
}

func (p *timePayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []time.Time
	var na []bool

	if timeable, ok := payload.(Timeable); ok {
		vals, na = timeable.Times()
	} else {
		vals, na = NAPayload(payload.Len()).(Timeable).Times()
	}

	newVals := make([]time.Time, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.NA)
	copy(newNA[p.length:], na)

	return TimePayload(newVals, newNA)
}

func (p *timePayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *timePayload) adjustToLesserSize(size int) Payload {
	data, na := AdjustToLesserSizeWithNA(p.data, p.NA, size)

	return TimePayload(data, na)
}

func (p *timePayload) adjustToBiggerSize(size int) Payload {
	data, na := AdjustToBiggerSizeWithNA(p.data, p.NA, p.length, size)

	return TimePayload(data, na)
}

func (p *timePayload) StrForElem(idx int) string {
	return p.data[idx-1].Format(p.printer.Format)
}

/* Finder interface */

func (p *timePayload) Find(needle any) int {
	return FindFn(needle, p.data, p.NA, p.convertComparator, p.eqFn)
}

func (p *timePayload) FindAll(needle any) []int {
	return FindAllFn(needle, p.data, p.NA, p.convertComparator, p.eqFn)
}

/* Ordered interface */

func (p *timePayload) Eq(val any) []bool {
	return EqFn(val, p.data, p.NA, p.convertComparator, p.eqFn)
}

func (p *timePayload) Neq(val any) []bool {
	return NeqFn(val, p.data, p.NA, p.convertComparator, p.eqFn)
}

func (p *timePayload) convertComparator(val any) (time.Time, bool) {
	v, ok := val.(time.Time)

	return v, ok
}

func (p *timePayload) eqFn(f, s time.Time) bool {
	return f.Equal(s)
}

func (p *timePayload) Gt(val any) []bool {
	return GtFn(val, p.data, p.NA, p.convertComparator, p.ltFn)
}

func (p *timePayload) ltFn(f, s time.Time) bool {
	return f.Before(s)
}

func (p *timePayload) Lt(val any) []bool {
	return LtFn(val, p.data, p.NA, p.convertComparator, p.ltFn)
}

func (p *timePayload) Gte(val any) []bool {
	return GteFn(val, p.data, p.NA, p.convertComparator, p.eqFn, p.ltFn)
}

func (p *timePayload) Lte(val any) []bool {
	return LteFn(val, p.data, p.NA, p.convertComparator, p.eqFn, p.ltFn)
}

func (p *timePayload) Groups() ([][]int, []any) {
	groups, values := GroupsForData(p.data, p.NA)

	return groups, values
}

func (p *timePayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[string]bool{}
	wasNA := false
	for i := 0; i < p.length; i++ {
		is := false

		if p.NA[i] {
			if !wasNA {
				is = true
				wasNA = true
			}
		} else {
			strTime := p.data[i].Format(p.printer.Format)

			if _, ok := valuesMap[strTime]; !ok {
				is = true
				valuesMap[strTime] = true
			}
		}

		booleans[i] = is
	}

	return booleans
}

func (p *timePayload) Options() []option.Option {
	return []option.Option{
		ConfOption{keyOptionTimeFormat, p.printer.Format},
	}
}

func (p *timePayload) SetOption(name string, val any) bool {
	switch name {
	case keyOptionTimeFormat:
		p.printer.Format = val.(string)
	default:
		return false
	}

	return true
}

func (p *timePayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []time.Time
	var srcNA []bool

	if same, ok := payload.(*timePayload); ok {
		srcData = same.data
		srcNA = same.NA
	} else if timeable, ok := payload.(Timeable); ok {
		srcData, srcNA = timeable.Times()
	} else {
		return p
	}

	dstData := make([]time.Time, p.length)
	dstNA := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.NA[i] && !srcNA[i] {
			dstData[i] = srcData[i]
			dstNA[i] = false
		} else {
			dstData[i] = p.data[i]
			dstNA[i] = p.NA[i]
		}
	}

	return TimePayload(dstData, dstNA, p.Options()...)
}

// TimePayload creates a payload with string data.
//
// Available options are:
//   - OptionTimeFormat(format string) - sets a time format for conversion to string.
func TimePayload(data []time.Time, na []bool, options ...option.Option) Payload {
	length := len(data)
	conf := MergeOptions(options)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := NAPayload(0)
			return emp
		}
	}

	vecData := make([]time.Time, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = time.Time{}
		} else {
			vecData[i] = data[i]
		}
	}

	printer := TimePrinter{Format: time.RFC3339}

	payload := &timePayload{
		length:  length,
		data:    vecData,
		printer: printer,
		NAble: embed.NAble{
			NA: vecNA,
		},
	}

	conf.SetOptions(payload)

	payload.Arrangeable = embed.Arrangeable{
		Length: payload.length,
		NAble:  payload.NAble,
		FnLess: func(i, j int) bool {
			return payload.data[i].Before(payload.data[j])
		},
		FnEqual: func(i, j int) bool {
			return payload.data[i].Equal(payload.data[j])
		},
	}

	return payload
}

// TimeWithNA creates a vector with TimePayload and allows to set NA-values.
func TimeWithNA(data []time.Time, na []bool, options ...option.Option) Vector {
	return New(TimePayload(data, na, options...), options...)
}

// Time creates a vector with TimePayload.
func Time(data []time.Time, options ...option.Option) Vector {
	return TimeWithNA(data, nil, options...)
}
