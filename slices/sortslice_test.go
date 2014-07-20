package slices

import (
	"math/rand"
	"sort"
	"testing"
)

type iv struct {
	dummy0 int64
	dummy1 int64
	val    int
	dummy2 int64
	dummy3 int64
	dummy4 int64
}

func TestSortSlice(t *testing.T) {
	slice := []int{3, 2, 1}
	Sort(slice, func(a, b int) bool { return slice[a] < slice[b] })
	if len(slice) != 3 {
		t.Fatalf("slice length changed")
	}
	if slice[0] != 1 {
		t.Fatalf("slice 0 is %v, not 1, %#v", slice[0], slice)
	}
	if slice[1] != 2 {
		t.Fatalf("slice 1 is %v, not 2", slice[1])
	}
	if slice[2] != 3 {
		t.Fatalf("slice 2 is %v, not 3", slice[2])
	}
}

func BenchmarkPkgSort(b *testing.B) {
	islice := rand.Perm(30000)
	for i := 0; i < b.N; i++ {
		sl := make([]int, len(islice), len(islice))
		copy(sl, islice)
		sort.Ints(sl)
	}
}

func BenchmarkSortSlices(b *testing.B) {
	islice := rand.Perm(30000)
	for i := 0; i < b.N; i++ {
		sl := make([]int, len(islice), len(islice))
		copy(sl, islice)
		Sort(sl, func(a, b int) bool { return sl[a] < sl[b] })
	}
}

func BenchmarkSortSwapper(b *testing.B) {
	islice := rand.Perm(30000)
	for i := 0; i < b.N; i++ {
		sl := make([]int, len(islice), len(islice))
		copy(sl, islice)
		Sort(sort.IntSlice(sl), func(a, b int) bool { return sl[a] < sl[b] })
	}
}

func BenchmarkStructSort(b *testing.B) {
	isl := rand.Perm(30000)
	islice := make([]iv, len(isl), len(isl))
	for i, v := range isl {
		islice[i] = iv{val: v}
	}
	for i := 0; i < b.N; i++ {
		sl := make([]iv, len(islice), len(islice))
		copy(sl, islice)
		Sort(sl, func(a, b int) bool { return sl[a].val < sl[b].val })
	}
}

type iswrapper []iv

func (w iswrapper) Len() int           { return len(w) }
func (w iswrapper) Swap(a, b int)      { w[a], w[b] = w[b], w[a] }
func (w iswrapper) Less(a, b int) bool { return w[a].val < w[b].val }

func BenchmarkStructSwapper(b *testing.B) {
	isl := rand.Perm(30000)
	islice := make([]iv, len(isl), len(isl))
	for i, v := range isl {
		islice[i] = iv{val: v}
	}
	for i := 0; i < b.N; i++ {
		sl := make([]iv, len(islice), len(islice))
		copy(sl, islice)
		Sort(iswrapper(sl), func(a, b int) bool { return sl[a].val < sl[b].val })
	}
}

func BenchmarkStructPkgSort(b *testing.B) {
	isl := rand.Perm(30000)
	islice := make([]iv, len(isl), len(isl))
	for i, v := range isl {
		islice[i] = iv{val: v}
	}
	for i := 0; i < b.N; i++ {
		sl := make([]iv, len(islice), len(islice))
		copy(sl, islice)
		sort.Sort(iswrapper(sl))
	}
}
