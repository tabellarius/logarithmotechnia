package vector

import (
	"time"
)

type TimePrinter struct {
	Format string
}

type timePayload struct {
	length  int
	data    []time.Time
	printer TimePrinter
	DefNAble
	DefArrangeable
}

func (p *timePayload) Type() string {
	return "time"
}

func (p *timePayload) Len() int {
	return p.length
}

func (p *timePayload) Pick(idx int) any {
	return pickValueWithNA(idx, p.data, p.na, p.length)
}

func (p *timePayload) Data() []any {
	return dataWithNAToInterfaceArray(p.data, p.na)
}

func (p *timePayload) ByIndices(indices []int) Payload {
	data, na := byIndices(indices, p.data, p.na, time.Time{})

	return TimePayload(data, na, p.Options()...)
}

func (p *timePayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[time.Time](whicher)
}

func (p *timePayload) Which(whicher any) []bool {
	return which(p.data, p.na, whicher)
}

func (p *timePayload) SupportsApplier(applier any) bool {
	return supportsApplier[time.Time](applier)
}

func (p *timePayload) Apply(applier any) Payload {
	data, na := apply(p.data, p.na, applier, time.Time{})

	if data == nil {
		return NAPayload(p.length)
	}

	return TimePayload(data, na)
}

func (p *timePayload) ApplyTo(indices []int, applier any) Payload {
	data, na := applyTo(indices, p.data, p.na, applier, time.Time{})

	if data == nil {
		return NAPayload(p.length)
	}

	return TimePayload(data, na, p.Options()...)
}

func (p *timePayload) SupportsSummarizer(summarizer any) bool {
	return supportsSummarizer[time.Time](summarizer)
}

func (p *timePayload) Summarize(summarizer any) Payload {
	val, na := summarize(p.data, p.na, summarizer, time.Time{}, time.Time{})

	return TimePayload([]time.Time{val}, []bool{na})
}

func (p *timePayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = ""
		} else {
			data[i] = p.StrForElem(i + 1)
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *timePayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *timePayload) Anies() ([]any, []bool) {
	if p.length == 0 {
		return []any{}, []bool{}
	}

	data := make([]any, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = nil
		} else {
			data[i] = p.data[i]
		}
	}

	na := make([]bool, p.length)
	copy(na, p.na)

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
	copy(newNA, p.na)
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
	data, na := adjustToLesserSizeWithNA(p.data, p.na, size)

	return TimePayload(data, na)
}

func (p *timePayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.na, p.length, size)

	return TimePayload(data, na)
}

func (p *timePayload) StrForElem(idx int) string {
	return p.data[idx-1].Format(p.printer.Format)
}

/* Finder interface */

func (p *timePayload) Find(needle any) int {
	return findFn(needle, p.data, p.na, p.convertComparator, p.eqFn)
}

func (p *timePayload) FindAll(needle any) []int {
	return findAllFn(needle, p.data, p.na, p.convertComparator, p.eqFn)
}

/* Ordered interface */

func (p *timePayload) Eq(val any) []bool {
	return eqFn(val, p.data, p.na, p.convertComparator, p.eqFn)
}

func (p *timePayload) Neq(val any) []bool {
	return neqFn(val, p.data, p.na, p.convertComparator, p.eqFn)
}

func (p *timePayload) convertComparator(val any) (time.Time, bool) {
	v, ok := val.(time.Time)

	return v, ok
}

func (p *timePayload) eqFn(f, s time.Time) bool {
	return f.Equal(s)
}

func (p *timePayload) Gt(val any) []bool {
	return gtFn(val, p.data, p.na, p.convertComparator, p.gtFn)
}

func (p *timePayload) gtFn(f, s time.Time) bool {
	return f.After(s)
}

func (p *timePayload) Lt(val any) []bool {
	return ltFn(val, p.data, p.na, p.convertComparator, p.gtFn)
}

func (p *timePayload) Gte(val any) []bool {
	return gteFn(val, p.data, p.na, p.convertComparator, p.eqFn, p.gtFn)
}

func (p *timePayload) Lte(val any) []bool {
	return lteFn(val, p.data, p.na, p.convertComparator, p.eqFn, p.gtFn)
}

func (p *timePayload) Groups() ([][]int, []any) {
	groups, values := groupsForData(p.data, p.na)

	return groups, values
}

func (p *timePayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[string]bool{}
	wasNA := false
	for i := 0; i < p.length; i++ {
		is := false

		if p.na[i] {
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

func (p *timePayload) Options() []Option {
	return []Option{
		OptionTimeFormat(p.printer.Format),
	}
}

func TimePayload(data []time.Time, na []bool, options ...Option) Payload {
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
	if conf.HasOption(KeyOptionTimeFormat) {
		printer.Format = conf.Value(KeyOptionTimeFormat).(string)
	}

	payload := &timePayload{
		length:  length,
		data:    vecData,
		printer: printer,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}

	payload.DefArrangeable = DefArrangeable{
		Length:   payload.length,
		DefNAble: payload.DefNAble,
		FnLess: func(i, j int) bool {
			return payload.data[i].Before(payload.data[j])
		},
		FnEqual: func(i, j int) bool {
			return payload.data[i].Equal(payload.data[j])
		},
	}

	return payload
}

func (p *timePayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []time.Time
	var srcNA []bool

	if same, ok := payload.(*timePayload); ok {
		srcData = same.data
		srcNA = same.na
	} else if timeable, ok := payload.(Timeable); ok {
		srcData, srcNA = timeable.Times()
	} else {
		return p
	}

	dstData := make([]time.Time, p.length)
	dstNA := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] && !srcNA[i] {
			dstData[i] = srcData[i]
			dstNA[i] = false
		} else {
			dstData[i] = p.data[i]
			dstNA[i] = p.na[i]
		}
	}

	return TimePayload(dstData, dstNA, p.Options()...)
}

func TimeWithNA(data []time.Time, na []bool, options ...Option) Vector {
	return New(TimePayload(data, na, options...), options...)
}

func Time(data []time.Time, options ...Option) Vector {
	return TimeWithNA(data, nil, options...)
}
