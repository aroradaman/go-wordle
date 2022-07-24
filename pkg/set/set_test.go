package set

import (
	"testing"
)

func TestSet_Contains(t *testing.T) {
	runes := []rune{'a', 'b', 'd', 'e'}
	s := NewSet(runes...)

	if s.Contains('c') {
		t.Errorf("c not in Container(a, b, d, e)")
	}

	for _, r := range runes {
		if !s.Contains(r) {
			t.Errorf("%s is in Container(a, b, d, e)", string(r))
		}
	}
}

func TestSet_Equals(t *testing.T) {
	s1 := NewSet('a', 'b', 'c')
	s2 := NewSet('c', 'a', 'b')

	if !s1.Equals(s2) {
		t.Errorf("%s & %s are same", s1.String(), s2.String())
	}
}

func TestSet_Add(t *testing.T) {

	s1 := NewSet('a', 'b', 'c')
	s1.Add('d')

	s2 := NewSet('a', 'b', 'c', 'd')

	if !s1.Equals(s2) {
		t.Errorf("%s & %s are same", s1.String(), s2.String())
	}

	s3 := NewSet('a', 'b', 'c')
	s4 := NewSet('a')
	s4.Add('b', 'c')

	if !s3.Equals(s4) {
		t.Errorf("%s & %s are same", s3.String(), s4.String())
	}
}

func TestSet_Pop(t *testing.T) {
	s1 := NewSet('a', 'b', 'c', 'c', 'd', 'd', 'd', 'e')
	s2 := NewSet('a', 'b', 'c', 'c', 'e')

	s1.Pop('d')
	if !s1.Equals(s2) {
		t.Errorf("Container(a,b,c,c,e) & Container(a,b,c,c,e) are same")
	}

	s1.Pop('c', 'e')
	s3 := NewSet('a', 'b')

	if !s1.Equals(s3) {
		t.Errorf("Container(a,b) & Container(a,b) are same")
	}
}
