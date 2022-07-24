package solver

import (
	. "github.com/daman1807/go-wordle/pkg/container"
	. "github.com/daman1807/go-wordle/pkg/set"
	"testing"
)

func TestState_Equals(t *testing.T) {
	s1 := &State{
		includes: NewContainer('A', 'B', 'C'),
		positionalStates: []*PositionalState{
			{
				fixed:    true,
				val:      'A',
				excludes: NewSet('D', 'E', 'F'),
			},
			{
				fixed:    false,
				val:      'B',
				excludes: NewSet('M', 'N', 'O'),
			},
		},
	}

	s2 := &State{
		includes: NewContainer('A', 'B', 'C'),
		positionalStates: []*PositionalState{
			{
				fixed:    true,
				val:      'A',
				excludes: NewSet('D', 'E', 'F'),
			},
			{
				fixed:    false,
				val:      'B',
				excludes: NewSet('M', 'N', 'O'),
			},
		},
	}

	if !s1.Equals(s2) {
		t.Errorf("States s1 and s2 are same")
	}
}

func TestState_Update(t *testing.T) {
	var s1, s2 *State
	var attempt string
	var feedback Feedback

	s1 = newState()

	// answer: CLOTH; attempt: SCRAM
	attempt = "SCRAM"

	feedback = NewFeedback(attempt)

	feedback[0].annotation = Excluded
	feedback[1].annotation = Included
	feedback[2].annotation = Excluded
	feedback[3].annotation = Excluded
	feedback[4].annotation = Excluded

	s1.Update(attempt, feedback)

	s2 = &State{
		includes: NewContainer('C'),
		positionalStates: []*PositionalState{
			{
				fixed:    false,
				excludes: NewSet('S', 'R', 'A', 'M'),
			},
			{
				fixed:    false,
				excludes: NewSet('S', 'C', 'R', 'A', 'M'),
			},
			{
				fixed:    false,
				excludes: NewSet('S', 'R', 'A', 'M'),
			},
			{
				fixed:    false,
				excludes: NewSet('S', 'R', 'A', 'M'),
			},
			{
				fixed:    false,
				excludes: NewSet('S', 'R', 'A', 'M'),
			},
		},
	}

	if !s1.Equals(s2) {
		t.Errorf("States s1 and s2 are same")
	}

}
