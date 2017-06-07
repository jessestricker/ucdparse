package ucdparse

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	n, k, dk := 6, 200, 100
	a := make(RangeList, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		r1 := uint32(rand.Intn(k))
		r2 := r1 + uint32(rand.Intn(dk))
		a[i] = NewRange(r1, r2)
	}

	sort.Sort(a)
	t.Log(a)
	a.SortAndMerge()
	t.Log(a)

	// THIS NEED HUMAN VERIFICATION
}
