package solver

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewSolver(t *testing.T) {
	words := []string{"AWARD", "AWAKE", "AWAIT"}

	tree := newNode('.')

	// Adding AWA to tree
	tree.children['A'] = newNode('A')
	tree.children['A'].children['W'] = newNode('W')
	tree.children['A'].children['W'].children['A'] = newNode('A')

	// Adding AWARD to tree
	tree.children['A'].children['W'].children['A'].children['R'] = newNode('R')
	tree.children['A'].children['W'].children['A'].children['R'].children['D'] = newNode('D')

	// Adding AWAKE to tree
	tree.children['A'].children['W'].children['A'].children['K'] = newNode('K')
	tree.children['A'].children['W'].children['A'].children['K'].children['E'] = newNode('E')

	// Adding AWAIT to tree
	tree.children['A'].children['W'].children['A'].children['I'] = newNode('I')
	tree.children['A'].children['W'].children['A'].children['I'].children['T'] = newNode('T')

	solver := NewSolver(words, words[0])

	if !solver.tree.Equals(tree) {
		t.Errorf("Incorrect Tree")
	}
}

func TestSolver_Validate1(t *testing.T) {
	solver := NewSolver([]string{"CLOTH", "COLOR"}, "CLOTH")

	attempt := "COLOR"

	expected := NewFeedback(attempt)
	expected[0].annotation = Fixed    // C
	expected[1].annotation = Included // O
	expected[2].annotation = Included // L
	expected[3].annotation = Excluded // O
	expected[4].annotation = Excluded // R

	feedback, _ := solver.Validate(attempt)

	if !feedback.Equals(expected) {
		t.Errorf("Expecting %s, got %s ", expected.String(), feedback.String())
	}

}

func TestSolver_Validate2(t *testing.T) {
	solver := NewSolver([]string{"SLEPT", "SWELL"}, "SLEPT")

	attempt := "SWELL"

	expected := NewFeedback(attempt)
	expected[0].annotation = Fixed    // S
	expected[1].annotation = Excluded // W
	expected[2].annotation = Fixed    // E
	expected[3].annotation = Included // L
	expected[4].annotation = Excluded // L

	feedback, _ := solver.Validate(attempt)

	if !feedback.Equals(expected) {
		t.Errorf("Expecting %s, got %s ", expected.String(), feedback.String())
	}

}

func TestSolver_Validate3(t *testing.T) {
	solver := NewSolver([]string{"HATCH", "WATCH"}, "HATCH")

	attempt := "WATCH"

	expected := NewFeedback(attempt)
	expected[0].annotation = Excluded
	expected[1].annotation = Fixed
	expected[2].annotation = Fixed
	expected[3].annotation = Fixed
	expected[4].annotation = Fixed

	feedback, _ := solver.Validate(attempt)

	if !feedback.Equals(expected) {
		t.Errorf("Expecting %s, got %s ", expected.String(), feedback.String())
	}

}

func checkEquality(a []string, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}

func TestSolver_GetHints1(t *testing.T) {
	var expected, hints []string
	solver := NewSolver([]string{
		"ARROW", "ARSON", "AWARD", "AWAKE", "AWAIT", "AWARE", "DREAM", "BLESS", "BLIND", "BRAWL", "BRAWN",
	}, "ARROW")

	solver.Validate("BLESS")
	hints = solver.GetHints()

	expected = []string{"ARROW", "AWARD", "AWAIT"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}

	solver.Validate("AWARD")
	hints = solver.GetHints()

	expected = []string{"ARROW"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}
}

func TestSolver_GetHints2(t *testing.T) {
	var expected, hints []string
	solver := NewSolver([]string{"CLOTH", "SCRAM", "CRANE", "BLESS", "BLIND", "MANGO"}, "CLOTH")

	solver.Validate("SCRAM")
	hints = solver.GetHints()

	expected = []string{"CLOTH"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}
}

func TestSolver_GetHints3(t *testing.T) {
	var hints, expected []string

	solver := NewSolver([]string{
		"VAULT", "LAUGH", "VALUE", "URBAN", "FAULT", "AUDIO",
	}, "FAULT")

	solver.Validate("CRANE")
	hints = solver.GetHints()
	expected = []string{"LAUGH", "FAULT", "AUDIO", "VAULT"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}

	solver.Validate("AUDIO")
	hints = solver.GetHints()
	expected = []string{"LAUGH", "FAULT", "VAULT"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}

}

func TestSolver_GetHints4(t *testing.T) {
	var hints, expected []string

	solver := NewSolver([]string{
		"SWELL", "SLEPT",
	}, "SLEPT")

	solver.Validate("SWELL")

	hints = solver.GetHints()
	expected = []string{"SLEPT"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}

}

func TestSolver_GetHints5(t *testing.T) {
	var hints, expected []string

	solver := NewSolver([]string{
		"HATCH", "YACHT", "WATCH", "TACKY", "LATCH", "URBAN", "FAULT", "AUDIO", "DREAM", "DRONE",
	}, "HATCH")

	solver.Validate("AUDIO")

	hints = solver.GetHints()
	expected = []string{"HATCH", "YACHT", "WATCH", "TACKY", "LATCH"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}

	solver.Validate("CANDY")
	hints = solver.GetHints()
	expected = []string{"HATCH", "WATCH", "LATCH"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}

	solver.Validate("WATCH")

	hints = solver.GetHints()
	expected = []string{"HATCH", "LATCH"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}

	solver.Validate("LATCH")
	hints = solver.GetHints()
	expected = []string{"HATCH"}

	if !checkEquality(hints, expected) {
		t.Errorf("Expected %s, got %s", expected, hints)
	}
}
