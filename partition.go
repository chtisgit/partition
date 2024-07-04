package partition

import (
	"reflect"
)

type PredicateFunc = func(i int) bool

func Slice(slice any, predicate PredicateFunc) int {
	s := reflect.ValueOf(slice)

	n := s.Len()

	i, j := 0, 0
	for i != n {
		if predicate(i) {
			if i != j {
				tmp := s.Index(i)
				s.Index(i).Set(s.Index(j))
				s.Index(j).Set(tmp)
			}
			j++
		}
		i++
	}

	return j
}
