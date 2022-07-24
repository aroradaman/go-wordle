package set

import (
	"fmt"
	"reflect"
)

type Set struct {
	data map[rune]struct{}
}

func NewSet(members ...rune) *Set {
	data := make(map[rune]struct{})
	for _, m := range members {
		data[m] = struct{}{}

	}
	return &Set{data: data}

}

func (s *Set) String() string {
	runes := make([]rune, 0)

	for r, _ := range s.data {
		runes = append(runes, r)
	}
	return fmt.Sprintf("Set(%s)", string(runes))

}
func (s *Set) Equals(p *Set) bool {
	return reflect.DeepEqual(s, p)
}

func (s Set) Contains(r rune) bool {
	_, ok := s.data[r]
	return ok
}

func (s *Set) Add(runes ...rune) {
	for _, r := range runes {
		s.data[r] = struct{}{}
	}
}

func (s *Set) Pop(runes ...rune) {
	for _, r := range runes {
		_, ok := s.data[r]
		if ok {
			delete(s.data, r)
		}
	}
}
