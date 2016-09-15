package ucdparse

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	n, k, dk := 5, 50, 5
	a := make(RangeList, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		r1 := uint32(rand.Intn(k + 1))
		r2 := r1 + uint32(rand.Intn(dk+1))
		a[i] = NewRange(r1, r2)
	}

	sort.Sort(a)
	t.Log(a)
	a = a.MergeRanges()
	t.Log(a)

	// THIS NEED HUMAN VERIFICATION
}
