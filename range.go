package ucdparse

import (
	"fmt"
	"sort"
	"strconv"
)

type Range struct {
	first, last uint32
}

func NewRange(a, b uint32) Range {
	if a <= b {
		return Range{a, b}
	}
	return Range{b, a}
}
func NewRangeSingle(a uint32) Range {
	return Range{a, a}
}

func (r *Range) First() uint32  { return r.first }
func (r *Range) Last() uint32   { return r.last }
func (r *Range) IsSingle() bool { return r.first == r.last }
func (r *Range) Count() uint32  { return (r.last - r.first) + 1 }

func (r Range) String() string {
	if r.IsSingle() {
		return fmt.Sprintf("%04X", r.first)
	}
	return fmt.Sprintf("%04X..%04X", r.first, r.last)
}

type RangeList []Range

func (a RangeList) String() string {
	out := "[" + strconv.Itoa(len(a)) + "] {"
	for i, v := range a {
		if i != 0 {
			out += ", "
		}
		out += v.String()
	}
	return out + "}"
}

func (a RangeList) Len() int           { return len(a) }
func (a RangeList) Less(i, j int) bool { return a[i].first < a[j].first } // decreasing order
func (a RangeList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (rlp *RangeList) SortAndMerge() {
	rl := *rlp
	if len(rl) <= 1 {
		return
	}
	sort.Sort(rl)
	w := 0
	for r := 1; r < len(rl); r++ {
		last, current := rl[w], rl[r]
		if last.last+1 >= current.first {
			// overlapping ranges, merge
			if current.last > last.last {
				last.last = current.last
				rl[w] = last
			}
		} else {
			// no overlap, add new
			w++
			rl[w] = current
		}
	}
	*rlp = rl[:w+1]
}
