package delay

import (
	"time"
	"sync/atomic"
	"bytes"
	"sort"
	"fmt"
)

const (
	maxDuration time.Duration = 1<<63 - 1
)

type TimeDurationSlice []time.Duration
func (p TimeDurationSlice) Len() int           { return len(p) }
func (p TimeDurationSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p TimeDurationSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type DelayStat struct {
	sections TimeDurationSlice
	m map[time.Duration]*int64
}

func NewDelayStat() *DelayStat {
	this := &DelayStat{
		sections : TimeDurationSlice{},
		m : make(map[time.Duration]*int64),
    }
	this.SetSections()
	return this
}

func (this *DelayStat) SetSections(sec ... time.Duration) {
	this.sections = TimeDurationSlice(sec)
	this.sections = append(this.sections, maxDuration)
	// sort
	sort.Sort(this.sections)
	// insert m
	this.m = make(map[time.Duration]*int64)
	for _, s := range this.sections {
		var v int64
		this.m[s] = &v
    }
}

func (this *DelayStat) Add(d time.Duration) {
	for _, s := range this.sections {
		if d <= s {
			atomic.AddInt64(this.m[s], 1)
			return
        }
	}
}

type Value struct {
	T time.Duration
	V int64
}

type Values []Value
func (p Values) Len() int           { return len(p) }
func (p Values) Less(i, j int) bool { return p[i].T < p[j].T }
func (p Values) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (this *DelayStat) Rotate() (values Values) {
	values = make(Values, 0, len(this.m))
	for k, v := range this.m {
		val := atomic.SwapInt64(v, 0)
		values = append(values, Value{ T : k, V : val })
	}
	sort.Sort(values)
	return
}

func (this *DelayStat) ToString(values Values) string {
	buf := bytes.NewBufferString("")
	var all int64
	for _, v := range values {
		all += v.V
	}
	buf.WriteString(fmt.Sprintf("TPS:%d", all))

	for _, v := range values {
		buf.WriteString(", ")
		if v.T != maxDuration {
			buf.WriteString(fmt.Sprintf("<%s:%d(%d%%)", v.T.String(), v.V, v.V * 100 / (all + 1)))
		} else {
			buf.WriteString(fmt.Sprintf("Others:%d(%d%%)", v.V, v.V * 100 / (all + 1)))
        }
	}
	return buf.String()
}

func (this *DelayStat) RotateToString() string {
	return this.ToString(this.Rotate())
}
