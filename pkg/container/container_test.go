package container

import (
	"testing"
)

func TestContainer_Contains(t *testing.T) {
	runes := []rune{'a', 'b', 'd', 'e'}
	s := NewContainer(runes...)

	if s.Contains('c') {
		t.Errorf("c not in Container(a, b, d, e)")
	}

	for _, r := range runes {
		if !s.Contains(r) {
			t.Errorf("%s is in Container(a, b, d, e)", string(r))
		}
	}
}

func TestContainer_Equals(t *testing.T) {
	s1 := NewContainer('a', 'b', 'c')
	s2 := NewContainer('c', 'a', 'b')

	if !s1.Equals(s2) {
		t.Errorf("Container(a,b,c) & Container(c,a,b) are same")
	}
}

func TestContainer_Add(t *testing.T) {

	s1 := NewContainer('a', 'b', 'c')
	s1.Add('d')

	s2 := NewContainer('a', 'b', 'c', 'd')

	if !s1.Equals(s2) {
		t.Errorf("Container(a,b,c,d) & Container(a,b,c,d) are same")
	}

	s3 := NewContainer('a', 'b', 'c')
	s4 := NewContainer('a')
	s4.Add('b', 'c')

	if !s3.Equals(s4) {
		t.Errorf("Container(a,b,c) & Container(a,b,c) are same")
	}
}

func TestContainer_Pop(t *testing.T) {

	s1 := NewContainer('a', 'b', 'c', 'c')
	s1.Pop('c')

	s2 := NewContainer('a', 'b', 'c')

	if !s1.Equals(s2) {
		t.Errorf("Container(a,b,c) & Container(a,b,c) are same")
	}

	s3 := NewContainer('a', 'b')
	s1.Pop('c')

	if !s1.Equals(s3) {
		t.Errorf("Container(a,b) & Container(a,b) are same")
	}

}

func TestContainer_Remove(t *testing.T) {
	s1 := NewContainer('a', 'b', 'c', 'c', 'd', 'd', 'd', 'e')
	s2 := NewContainer('a', 'b', 'c', 'c', 'e')

	s1.Remove('d')
	if !s1.Equals(s2) {
		t.Errorf("Container(a,b,c,c,e) & Container(a,b,c,c,e) are same")
	}

	s1.Remove('c', 'e')
	s3 := NewContainer('a', 'b')

	if !s1.Equals(s3) {
		t.Errorf("Container(a,b) & Container(a,b) are same")
	}
}

func TestContainer_GetCount(t *testing.T) {
	s := NewContainer('a', 'b', 'c', 'c', 'd', 'd', 'd', 'e')

	if s.GetCount('c') != 2 {
		t.Errorf("Expected 2, got %d", s.GetCount('c'))
	}

	if s.GetCount('d') != 3 {
		t.Errorf("Expected 3, got %d", s.GetCount('c'))
	}

}
