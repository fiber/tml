// Package selekt simplifies the use of the Select vehicle provided by the reflect package
package selekt

import (
	"reflect"
)

type S struct {
	cases []reflect.SelectCase
	value reflect.Value
}

func New() *S {
	return &S{}
}

func (s *S) Send(channel interface{}, value interface{}) int {
	s.cases = append(s.cases, Send(channel, value))
	return len(s.cases) - 1
}

func (s *S) Recv(channel interface{}) int {
	s.cases = append(s.cases, Recv(channel))
	return len(s.cases) - 1
}

func (s *S) Default() int {
	s.cases = append(s.cases, Default())
	return len(s.cases) - 1
}

func (s *S) Select() (chosen int, ok bool) {
	chosen, value, ok := reflect.Select(s.cases)
	s.value = value
	return chosen, ok
}

func (s *S) Value(v interface{}) bool {
	reflect.ValueOf(v).Set(s.value)
	return true
}

func Send(channel interface{}, value interface{}) reflect.SelectCase {
	return reflect.SelectCase{
		Dir:  reflect.SelectSend,
		Chan: reflect.ValueOf(channel),
		Send: reflect.ValueOf(value),
	}
}

func Recv(channel interface{}) reflect.SelectCase {
	return reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(channel),
	}
}

func Default() reflect.SelectCase {
	return reflect.SelectCase{
		Dir: reflect.SelectDefault,
	}
}
