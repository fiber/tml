package slices

import (
	"reflect"
	"sort"
	// "sync"
)

type (
	sliceSort struct {
		less  func(a, b int) bool
		slice reflect.Value
		temp  reflect.Value
		swap  swapper
		//sync.RWMutex
	}
	sliceSwap sliceSort
	swapper   interface {
		Swap(a, b int)
	}
)

// Sort sorts a slice using a provided "Less" func. It provides Swap() and
// Len() using the reflect package.  If slice has a Swap() method, that will be
// used instead of the generic Swap(). Note that this is probably substantially
// slower (approx. 3 times) than building your own generic swap function.
func Sort(slice interface{}, lessfunc func(a, b int) bool) {
	s := sliceSort{slice: reflect.ValueOf(slice), less: lessfunc}
	s.temp = reflect.New(reflect.TypeOf(slice).Elem()).Elem()
	if sw, ok := slice.(swapper); ok {
		s.swap = sw
	} else {
		sw := sliceSwap(s)
		s.swap = &sw
	}
	sort.Sort(&s)
}

func (s *sliceSort) Len() int {
	return s.slice.Len()
}

func (s *sliceSort) Swap(a, b int) {
	s.swap.Swap(a, b)
}

func (s sliceSwap) Swap(a, b int) {
	s.temp.Set(s.slice.Index(a))
	s.slice.Index(a).Set(s.slice.Index(b))
	s.slice.Index(b).Set(s.temp)
}

func (s *sliceSort) Less(a, b int) bool {
	r := s.less(a, b)
	return r
}
